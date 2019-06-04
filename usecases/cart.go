package usecases

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ilhammhdd/kudaki-entities/kudakiredisearch"

	"github.com/google/uuid"

	"github.com/ilhammhdd/kudaki-entities/store"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/rental"
)

type AddCartItem struct {
	DBO            DBOperator
	CartClient     kudakiredisearch.RedisClient
	ItemClient     kudakiredisearch.RedisClient
	CartItemClient kudakiredisearch.RedisClient
}

func (aci AddCartItem) initInOutEvent(in proto.Message) (inEvent *events.AddCartItemRequested, outEvent *events.CartItemAdded) {
	inEventTemp := in.(*events.AddCartItemRequested)

	var outEventTemp events.CartItemAdded
	outEventTemp.Uid = inEventTemp.Uid
	outEventTemp.EventStatus = &events.Status{}

	return inEventTemp, &outEventTemp
}

func (aci AddCartItem) itemExists(itemUUID string) (item *store.Item, ok bool) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), aci.ItemClient.Name())
	rawQuery := fmt.Sprintf(`@item_uuid:"%s"`, kudakiredisearch.RedisearchText(itemUUID).Sanitize())
	itemDocs, total, err := client.Search(redisearch.NewQuery(rawQuery))
	errorkit.ErrorHandled(err)

	if total != 1 {
		ok = false
		return
	}

	item = new(store.Item)
	ok = true
	price, err := strconv.ParseInt(itemDocs[0].Properties["item_price"].(string), 10, 32)
	errorkit.ErrorHandled(err)
	rating, err := strconv.ParseFloat(itemDocs[0].Properties["item_rating"].(string), 10)
	errorkit.ErrorHandled(err)

	item.Description = itemDocs[0].Properties["item_description"].(string)
	item.Name = itemDocs[0].Properties["item_name"].(string)
	item.Photo = itemDocs[0].Properties["item_photo"].(string)
	item.Price = int32(price)
	item.Rating = float32(rating)
	item.Unit = itemDocs[0].Properties["item_unit"].(string)
	item.Uuid = itemDocs[0].Properties["item_uuid"].(string)
	return
}

func (aci AddCartItem) cartExists(cartUUID string) (cart *rental.Cart, ok bool) {
	rsClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), aci.CartClient.Name())
	rsClient.CreateIndex(aci.CartClient.Schema())
	rawQuery := fmt.Sprintf(`@cart_uuid:"%s" @cart_open:1`, kudakiredisearch.RedisearchText(cartUUID).Sanitize())
	cartDocs, totalCarts, err := rsClient.Search(redisearch.NewQuery(rawQuery))
	errorkit.ErrorHandled(err)

	if totalCarts == 0 {
		ok = false
		return
	}

	cart = new(rental.Cart)
	ok = true

	cartOpenNum := cartDocs[0].Properties["cart_open"].(int64)
	cartTotalItems, err := strconv.ParseUint(cartDocs[0].Properties["cart_total_items"].(string), 10, 32)
	errorkit.ErrorHandled(err)
	cartTotalPrice, err := strconv.ParseUint(cartDocs[0].Properties["cart_total_price"].(string), 10, 32)
	errorkit.ErrorHandled(err)

	cart.Open = cartOpenNum == 1
	cart.TotalItems = uint32(cartTotalItems)
	cart.TotalPrice = uint32(cartTotalPrice)
	cart.Uuid = cartDocs[0].Properties["cart_uuid"].(string)
	return
}

func (aci AddCartItem) cartItemExists() (*rental.CartItem, bool) {

	return nil, false
}

func (aci AddCartItem) Handle(in proto.Message) (out proto.Message) {
	inEvent, outEvent := aci.initInOutEvent(in)
	outEvent.EventStatus.HttpCode = http.StatusOK
	outEvent.EventStatus.Timestamp = ptypes.TimestampNow()

	cartItem := new(rental.CartItem)
	cartItem.Uuid = uuid.New().String()

	if item, ok := aci.itemExists(inEvent.ItemUuid); ok {
		cartItem.Item = item

		if inEvent.CartUuid == "" {
			newCart := new(rental.Cart)
			newCart.Open = true
			newCart.TotalItems = uint32(inEvent.ItemAmount)
			newCart.TotalPrice = uint32(item.Price * inEvent.ItemAmount)
			newCart.User = inEvent.User
			newCart.Uuid = uuid.New().String()

			cartItem.Cart = newCart
			cartItem.Item = item
		} else {
			if cart, ok := aci.cartExists(inEvent.CartUuid); ok {
				cartItem.Cart = cart
				cartItem.Item = item
				if existedCartItem, ok := aci.cartItemExists(); ok {
					cartItem.TotalAmount = existedCartItem.TotalAmount + uint32(inEvent.ItemAmount)
					cartItem.TotalPrice = existedCartItem.TotalAmount + uint32(item.Price*inEvent.ItemAmount)
					cartItem.Uuid = existedCartItem.Uuid
				} else {
					cartItem.TotalAmount = uint32(inEvent.ItemAmount)
					cartItem.TotalPrice = uint32(item.Price * inEvent.ItemAmount)
					cartItem.Uuid = uuid.New().String()
				}
			} else {
				outEvent.EventStatus.Errors = []string{"cart with the given uuid doesn't exists"}
				outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
				outEvent.EventStatus.HttpCode = http.StatusNotFound
				return outEvent
			}
		}
	} else {
		outEvent.EventStatus.Errors = []string{"item with the given uuid not found"}
		outEvent.EventStatus.HttpCode = http.StatusNotFound
		outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
	}

	return outEvent
}
