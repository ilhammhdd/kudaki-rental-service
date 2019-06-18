package usecases

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/ilhammhdd/kudaki-entities/store"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/rental"
	"github.com/ilhammhdd/kudaki-entities/user"
)

type AddCartItem struct {
	DBO              DBOperator
	Sanitizer        RedisearchTextSanitizer
	StorefrontClient RedisClient
	ItemClient       RedisClient
}

func (aci *AddCartItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := aci.initInOutEvent(in)

	existedStorefront, ok := aci.storefrontExists(inEvent.StorefrontUuid)
	if !ok {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"storefront with the given UUID not found"}
		return outEvent
	}

	existedItem, ok := aci.itemExists(inEvent.ItemUuid, existedStorefront)
	if !ok {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"item with the given UUID not found"}
		return outEvent
	}

	usr := GetUserFromKudakiToken(inEvent.KudakiToken)

	cart, ok := aci.cartExists(usr)
	if ok {
		cart.TotalItems += uint32(inEvent.ItemAmount)
		cart.TotalPrice += uint32(inEvent.ItemAmount * existedItem.Price)
	} else {
		cart = aci.createNewCart(inEvent, existedItem)
	}

	cartItem, ok := aci.cartItemExists(cart, existedItem)
	if ok {
		cartItem.TotalAmount += uint32(inEvent.ItemAmount)
		cartItem.TotalPrice += uint32(inEvent.ItemAmount * existedItem.Price)
	} else {
		cartItem = aci.createNewCartItem(inEvent, cart, existedItem)
	}

	cartItem.Cart = cart
	cartItem.Item = existedItem

	outEvent.CartItem = cartItem
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (aci *AddCartItem) initInOutEvent(in proto.Message) (inEvent *events.AddCartItemRequested, outEvent *events.CartItemAdded) {
	inEvent = in.(*events.AddCartItemRequested)

	outEvent = new(events.CartItemAdded)
	outEvent.AddCartItemRequested = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid
	outEvent.CartItem = new(rental.CartItem)

	return
}

func (aci *AddCartItem) storefrontExists(storefrontUUID string) (*store.Storefront, bool) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), aci.StorefrontClient.Name())
	client.CreateIndex(aci.StorefrontClient.Schema())

	aci.Sanitizer.Set(storefrontUUID)
	rawQuery := fmt.Sprintf(`@storefront_uuid:"%s"`, aci.Sanitizer.Sanitize())
	storefrontDocs, total, err := client.Search(redisearch.NewQuery(rawQuery))
	errorkit.ErrorHandled(err)

	if total == 1 {
		rating, err := strconv.ParseFloat(storefrontDocs[0].Properties["storefront_rating"].(string), 10)
		errorkit.ErrorHandled(err)
		totalItem, err := strconv.ParseInt(storefrontDocs[0].Properties["storefront_total_item"].(string), 10, 32)
		errorkit.ErrorHandled(err)

		storefront := new(store.Storefront)
		storefront.Rating = float32(rating)
		storefront.TotalItem = int32(totalItem)
		aci.Sanitizer.Set(storefrontDocs[0].Properties["storefront_uuid"].(string))
		storefront.Uuid = aci.Sanitizer.UnSanitize()
		return storefront, true
	}

	return nil, false
}

func (aci *AddCartItem) cartExists(usr *user.User) (*rental.Cart, bool) {
	row, err := aci.DBO.QueryRow("SELECT uuid,total_price,total_items FROM carts WHERE user_uuid=? AND open=1;", usr.Uuid)
	errorkit.ErrorHandled(err)

	var existedCart rental.Cart
	if row.Scan(&existedCart.Uuid, &existedCart.TotalPrice, &existedCart.TotalItems) != sql.ErrNoRows {
		existedCart.Open = true
		existedCart.User = usr
		return &existedCart, true
	}
	return nil, false
}

func (aci *AddCartItem) createNewCart(inEvent *events.AddCartItemRequested, item *store.Item) *rental.Cart {
	newCart := new(rental.Cart)
	newCart.Open = true
	newCart.TotalItems = uint32(inEvent.ItemAmount)
	newCart.TotalPrice = uint32(item.Price * inEvent.ItemAmount)
	newCart.User = GetUserFromKudakiToken(inEvent.KudakiToken)
	newCart.Uuid = uuid.New().String()
	return newCart
}

func (aci *AddCartItem) itemExists(itemUUID string, storefront *store.Storefront) (*store.Item, bool) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), aci.ItemClient.Name())
	client.CreateIndex(aci.ItemClient.Schema())

	aci.Sanitizer.Set(itemUUID)
	sanitizedItemUUID := aci.Sanitizer.Sanitize()
	aci.Sanitizer.Set(storefront.Uuid)
	sanitizedstorefrontUUID := aci.Sanitizer.Sanitize()
	rawQuery := fmt.Sprintf(`@item_uuid:"%s" @storefront_uuid:"%s"`, sanitizedItemUUID, sanitizedstorefrontUUID)
	itemDocs, total, err := client.Search(redisearch.NewQuery(rawQuery))
	errorkit.ErrorHandled(err)

	if total == 1 {
		amount, _ := strconv.ParseInt(itemDocs[0].Properties["item_amount"].(string), 10, 32)
		price, _ := strconv.ParseInt(itemDocs[0].Properties["item_price"].(string), 10, 32)
		rating, _ := strconv.ParseFloat(itemDocs[0].Properties["item_rating"].(string), 10)

		var item store.Item
		item.Amount = int32(amount)
		item.Description = itemDocs[0].Properties["item_description"].(string)
		item.Name = itemDocs[0].Properties["item_name"].(string)
		item.Photo = itemDocs[0].Properties["item_photo"].(string)
		item.Price = int32(price)
		item.Rating = float32(rating)
		item.Storefront = storefront
		item.Unit = itemDocs[0].Properties["item_unit"].(string)
		item.Uuid = itemUUID
		return &item, true
	}
	return nil, false
}

func (aci *AddCartItem) cartItemExists(cart *rental.Cart, item *store.Item) (*rental.CartItem, bool) {
	row, err := aci.DBO.QueryRow("SELECT uuid,total_item,total_price FROM cart_items WHERE cart_uuid=? AND item_uuid=?;", cart.Uuid, item.Uuid)
	errorkit.ErrorHandled(err)

	var existedCartitem rental.CartItem
	if row.Scan(&existedCartitem.Uuid, &existedCartitem.TotalAmount, &existedCartitem.TotalPrice) != sql.ErrNoRows {
		existedCartitem.Cart = cart
		existedCartitem.Item = item
		return &existedCartitem, true
	}
	return nil, false
}

func (aci *AddCartItem) createNewCartItem(inEvent *events.AddCartItemRequested, cart *rental.Cart, item *store.Item) *rental.CartItem {
	newCartItem := new(rental.CartItem)
	newCartItem.Cart = cart
	newCartItem.Item = item
	newCartItem.TotalAmount = uint32(inEvent.ItemAmount)
	newCartItem.TotalPrice = uint32(inEvent.ItemAmount * item.Price)
	newCartItem.Uuid = uuid.New().String()

	return newCartItem
}

type DeleteCartItem struct {
	DBO        DBOperator
	Sanitizer  RedisearchTextSanitizer
	ItemClient RedisClient
}

func (dci *DeleteCartItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := dci.initInOutEvent(in)
	usr := GetUserFromKudakiToken(inEvent.KudakiToken)

	existedCartItem := dci.CartItemExists(inEvent.CartItemUuid)
	if existedCartItem == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"cart item with the given uuid not found"}
		return outEvent
	}

	existedCart := dci.CartExists(existedCartItem.Cart.Uuid, outEvent.User)
	if existedCart == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"cart corresponding with cart item not found"}
		return outEvent
	} else {
		existedCart.TotalItems -= existedCartItem.TotalAmount
		existedCart.TotalPrice -= existedCartItem.TotalPrice
		existedCartItem.Cart = existedCart
	}

	existedItem := dci.ItemExists(existedCartItem.Item.Uuid, usr.Uuid)
	if existedCart == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"item corresponding with cart item not found"}
		return outEvent
	} else {
		existedCartItem.Item = existedItem
	}

	outEvent.CartItem = existedCartItem
	outEvent.EventStatus.HttpCode = http.StatusOK

	return outEvent
}

func (dci *DeleteCartItem) initInOutEvent(in proto.Message) (inEvent *events.DeleteCartItemRequested, outEvent *events.CartItemDeleted) {
	inEvent = in.(*events.DeleteCartItemRequested)

	outEvent = new(events.CartItemDeleted)
	outEvent.DeleteCartItemRequested = inEvent
	outEvent.EventStatus = new(events.Status)
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	outEvent.Uid = inEvent.Uid
	outEvent.User = GetUserFromKudakiToken(inEvent.KudakiToken)

	return
}

func (dci *DeleteCartItem) CartItemExists(cartItemUUID string) *rental.CartItem {
	row, err := dci.DBO.QueryRow("SELECT cart_uuid,item_uuid,total_item,total_price FROM cart_items WHERE uuid = ?;", cartItemUUID)
	errorkit.ErrorHandled(err)

	var existedCartItem rental.CartItem
	existedCartItem.Cart = new(rental.Cart)
	existedCartItem.Item = new(store.Item)
	if row.Scan(&existedCartItem.Cart.Uuid, &existedCartItem.Item.Uuid, &existedCartItem.TotalAmount, &existedCartItem.TotalPrice) == nil {
		existedCartItem.Uuid = cartItemUUID
		return &existedCartItem
	}
	return nil
}

func (dci *DeleteCartItem) CartExists(cartUUID string, usr *user.User) *rental.Cart {
	row, err := dci.DBO.QueryRow("SELECT open,total_items,total_price FROM carts WHERE uuid=? AND user_uuid=?;", cartUUID, usr.Uuid)
	errorkit.ErrorHandled(err)

	var open int
	var existedCart rental.Cart
	if row.Scan(&open, &existedCart.TotalItems, &existedCart.TotalPrice) == nil {
		existedCart.User = usr
		existedCart.Uuid = cartUUID
		return &existedCart
	}
	return nil
}

func (dci *DeleteCartItem) ItemExists(itemUUID string, userUUID string) *store.Item {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), dci.ItemClient.Name())
	client.CreateIndex(dci.ItemClient.Schema())

	dci.Sanitizer.Set(itemUUID)
	rowQuery := fmt.Sprintf(`@item_uuid:"%s"`, dci.Sanitizer.Sanitize())
	doc, total, err := client.Search(redisearch.NewQuery(rowQuery))
	errorkit.ErrorHandled(err)

	var item store.Item
	item.Storefront = new(store.Storefront)
	if total != 0 {
		amount, err := strconv.ParseInt(doc[0].Properties["item_amount"].(string), 10, 32)
		errorkit.ErrorHandled(err)
		price, err := strconv.ParseInt(doc[0].Properties["item_price"].(string), 10, 32)
		errorkit.ErrorHandled(err)
		rating, err := strconv.ParseFloat(doc[0].Properties["item_rating"].(string), 10)
		errorkit.ErrorHandled(err)

		item.Uuid = dci.Sanitizer.UnSanitize()
		item.Storefront.Uuid = doc[0].Properties["storefront_uuid"].(string)
		item.Name = doc[0].Properties["item_name"].(string)
		item.Amount = int32(amount)
		item.Unit = doc[0].Properties["item_unit"].(string)
		item.Price = int32(price)
		item.Description = doc[0].Properties["item_description"].(string)
		item.Photo = doc[0].Properties["item_photo"].(string)
		item.Rating = float32(rating)

		return &item
	}
	return nil
}

type UpdateCartItem struct {
	DBO        DBOperator
	ItemClient RedisClient
	Sanitizer  RedisearchTextSanitizer
}

func (uci *UpdateCartItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := uci.initInOutEvent(in)

	existedCartItem := uci.cartItemExists(inEvent.CartItemUuid)
	if existedCartItem == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"cart item with the given uuid not found"}
		return outEvent
	}
	*outEvent.InitialCartItem = *existedCartItem

	existedItem := uci.itemExists(existedCartItem.Item.Uuid, outEvent.User.Uuid)
	if existedItem == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"cart item with the given uuid not found"}
		return outEvent
	}

	availableItemAmount := existedItem.Amount + int32(existedCartItem.TotalAmount)

	if availableItemAmount < inEvent.TotalItem {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"available item amount is insufficient"}
		return outEvent
	}

	existedCart := uci.cartExists(existedCartItem.Cart.Uuid, outEvent.User.Uuid)
	if existedCart == nil {
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Errors = []string{"cart with the given uuid not found"}
		return outEvent
	}

	existedCart.TotalItems -= existedCartItem.TotalAmount
	existedCart.TotalPrice -= existedCartItem.TotalPrice
	existedCart.TotalItems += uint32(inEvent.TotalItem)
	existedCart.TotalPrice += uint32(inEvent.TotalItem * existedItem.Price)

	(*existedCartItem).TotalAmount = uint32(inEvent.TotalItem)
	(*existedCartItem).TotalPrice = uint32(inEvent.TotalItem * existedItem.Price)
	existedCartItem.Cart = existedCart

	outEvent.UpdatedCartItem = existedCartItem
	outEvent.EventStatus.HttpCode = http.StatusOK
	return outEvent
}

func (uci *UpdateCartItem) initInOutEvent(in proto.Message) (inEvent *events.UpdateCartItemRequested, outEvent *events.CartItemUpdated) {
	inEvent = in.(*events.UpdateCartItemRequested)

	outEvent = &events.CartItemUpdated{
		EventStatus:             &events.Status{Timestamp: ptypes.TimestampNow()},
		InitialCartItem:         new(rental.CartItem),
		Uid:                     inEvent.Uid,
		UpdateCartItemRequested: inEvent,
		User:                    GetUserFromKudakiToken(inEvent.KudakiToken)}
	return
}

func (uci *UpdateCartItem) cartItemExists(cartItemUUID string) *rental.CartItem {
	row, err := uci.DBO.QueryRow("SELECT cart_uuid,item_uuid,total_item,total_price FROM cart_items WHERE uuid=?;", cartItemUUID)
	errorkit.ErrorHandled(err)

	var cartItem rental.CartItem
	cartItem.Cart = new(rental.Cart)
	cartItem.Item = new(store.Item)
	if row.Scan(&cartItem.Cart.Uuid, &cartItem.Item.Uuid, &cartItem.TotalAmount, &cartItem.TotalPrice) == sql.ErrNoRows {
		return nil
	}
	cartItem.Uuid = cartItemUUID
	return &cartItem
}

func (uci *UpdateCartItem) cartExists(cartUUID, userUUID string) *rental.Cart {
	row, err := uci.DBO.QueryRow("SELECT total_price,total_items,open FROM carts WHERE uuid=? AND user_uuid=?;", cartUUID, userUUID)
	errorkit.ErrorHandled(err)

	var cart rental.Cart
	cart.User = new(user.User)
	var open int
	if row.Scan(&cart.TotalPrice, &cart.TotalItems, &open) == sql.ErrNoRows {
		return nil
	}

	if open == 1 {
		cart.Open = true
	} else if open == 0 {
		cart.Open = false
	}
	cart.Uuid = cartUUID
	return &cart
}

func (uci *UpdateCartItem) itemExists(itemUUID string, userUUID string) *store.Item {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), uci.ItemClient.Name())
	client.CreateIndex(uci.ItemClient.Schema())

	uci.Sanitizer.Set(itemUUID)
	rowQuery := fmt.Sprintf(`@item_uuid:"%s"`, uci.Sanitizer.Sanitize())
	doc, total, err := client.Search(redisearch.NewQuery(rowQuery))
	errorkit.ErrorHandled(err)

	var item store.Item
	item.Storefront = new(store.Storefront)
	if total != 0 {
		amount, err := strconv.ParseInt(doc[0].Properties["item_amount"].(string), 10, 32)
		errorkit.ErrorHandled(err)
		price, err := strconv.ParseInt(doc[0].Properties["item_price"].(string), 10, 32)
		errorkit.ErrorHandled(err)
		rating, err := strconv.ParseFloat(doc[0].Properties["item_rating"].(string), 10)
		errorkit.ErrorHandled(err)

		item.Uuid = uci.Sanitizer.UnSanitize()
		item.Storefront.Uuid = doc[0].Properties["storefront_uuid"].(string)
		item.Name = doc[0].Properties["item_name"].(string)
		item.Amount = int32(amount)
		item.Unit = doc[0].Properties["item_unit"].(string)
		item.Price = int32(price)
		item.Description = doc[0].Properties["item_description"].(string)
		item.Photo = doc[0].Properties["item_photo"].(string)
		item.Rating = float32(rating)

		return &item
	}
	return nil
}
