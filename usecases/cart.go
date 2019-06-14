package usecases

import (
	"database/sql"
	"fmt"
	"log"
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
	kudakiredisearch "github.com/ilhammhdd/kudaki-externals/redisearch"
)

type AddCartItem struct {
	DBO DBOperator
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
		log.Println("cart exists")
		cart.TotalItems += uint32(inEvent.ItemAmount)
		cart.TotalPrice += uint32(inEvent.ItemAmount * existedItem.Price)
	} else {
		cart = aci.createNewCart(inEvent, existedItem)
	}

	cartItem, ok := aci.cartItemExists(cart, existedItem)
	if ok {
		log.Println("cart item exists")
		cartItem.TotalAmount += uint32(inEvent.ItemAmount)
		cartItem.TotalPrice += uint32(inEvent.ItemAmount * existedItem.Price)
	} else {
		cartItem = aci.createNewCartItem(inEvent, cart, existedItem)
	}

	cartItem.Cart = cart
	cartItem.Item = existedItem

	log.Printf("existed item : %v", existedItem)
	log.Printf("cart item after processing : total amount = %d", cartItem.TotalAmount)
	log.Printf("cart item after processing : total price = %d", cartItem.TotalPrice)
	log.Printf("cart after processing : total item = %d", cart.TotalItems)
	log.Printf("cart after processing : total price = %d", cart.TotalPrice)

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
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Storefront.Name())
	client.CreateIndex(kudakiredisearch.Storefront.Schema())

	rawQuery := fmt.Sprintf(`@storefront_uuid:"%s"`, kudakiredisearch.RedisearchText(storefrontUUID).Sanitize())
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
		storefront.Uuid = kudakiredisearch.RedisearchText(storefrontDocs[0].Properties["storefront_uuid"].(string)).UnSanitize()
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
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Item.Name())
	client.CreateIndex(kudakiredisearch.Item.Schema())

	rawQuery := fmt.Sprintf(`@item_uuid:"%s" @storefront_uuid:"%s"`, kudakiredisearch.RedisearchText(itemUUID).Sanitize(), kudakiredisearch.RedisearchText(storefront.Uuid).Sanitize())
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
