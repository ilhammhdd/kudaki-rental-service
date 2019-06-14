package eventdriven

import (
	"net/http"
	"os"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/rental"
	"github.com/ilhammhdd/kudaki-externals/mysql"
	kudakiredisearch "github.com/ilhammhdd/kudaki-externals/redisearch"
	"github.com/ilhammhdd/kudaki-rental-service/adapters"
	"github.com/ilhammhdd/kudaki-rental-service/usecases"
)

type AddCartItem struct {
}

func (aci *AddCartItem) Work() interface{} {
	usecase := usecases.AddCartItem{DBO: mysql.NewDBOperation()}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: aci,
		eventDrivenAdapter:  new(adapters.AddCartItem),
		eventDrivenUsecase:  &usecase,
		eventName:           events.RentalTopic_ADD_CART_ITEM_REQUESTED.String(),
		inTopics:            []string{events.RentalTopic_ADD_CART_ITEM_REQUESTED.String()},
		outTopic:            events.RentalTopic_CART_ITEM_ADDED.String()}
	ede.handle()

	return nil
}

func (aci *AddCartItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.CartItemAdded)
	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	aci.upsertCart(out.CartItem.Cart)
	aci.reIndexCart(out.CartItem.Cart)
	aci.upsertCartItem(out.CartItem)
	aci.reIndexCartItem(out.CartItem)
}

func (aci *AddCartItem) upsertCart(cart *rental.Cart) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("INSERT INTO carts(uuid,user_uuid,total_price,total_items,open) VALUES(?,?,?,?,?) ON DUPLICATE KEY UPDATE total_price=?,total_items=?;",
		cart.Uuid, cart.User.Uuid, cart.TotalPrice, cart.TotalItems, 1, cart.TotalPrice, cart.TotalItems)
	errorkit.ErrorHandled(err)
}

func (aci *AddCartItem) reIndexCart(cart *rental.Cart) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Cart.Name())
	client.CreateIndex(kudakiredisearch.Cart.Schema())

	doc := redisearch.NewDocument(kudakiredisearch.RedisearchText(cart.Uuid).Sanitize(), 1.0)
	doc.Set("cart_uuid", cart.Uuid)
	doc.Set("cart_total_price", cart.TotalPrice)
	doc.Set("cart_total_items", cart.TotalItems)
	doc.Set("cart_open", 1)

	err := client.IndexOptions(redisearch.IndexingOptions{Replace: true}, doc)
	errorkit.ErrorHandled(err)
}

func (aci *AddCartItem) upsertCartItem(cartItem *rental.CartItem) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("INSERT INTO cart_items(uuid,cart_uuid,item_uuid,total_item,total_price) VALUES(?,?,?,?,?) ON DUPLICATE KEY UPDATE total_item=?,total_price=?;",
		cartItem.Uuid, cartItem.Cart.Uuid, cartItem.Item.Uuid, cartItem.TotalAmount, cartItem.TotalPrice, cartItem.TotalAmount, cartItem.TotalPrice)
	errorkit.ErrorHandled(err)
}

func (aci *AddCartItem) reIndexCartItem(cartItem *rental.CartItem) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.CartItem.Name())
	client.CreateIndex(kudakiredisearch.CartItem.Schema())

	sanitizedCartItemUUID := kudakiredisearch.RedisearchText(cartItem.Uuid).Sanitize()
	doc := redisearch.NewDocument(sanitizedCartItemUUID, 1.0)
	doc.Set("cart_item_uuid", sanitizedCartItemUUID)
	doc.Set("cart_item_total_amount", cartItem.TotalAmount)
	doc.Set("cart_item_total_price", cartItem.TotalAmount)

	doc.Set("cart_uuid", kudakiredisearch.RedisearchText(cartItem.Cart.Uuid).Sanitize())
	doc.Set("cart_total_price", cartItem.Cart.TotalPrice)
	doc.Set("cart_total_items", cartItem.Cart.TotalItems)
	doc.Set("cart_open", 1)

	doc.Set("item_uuid", kudakiredisearch.RedisearchText(cartItem.Item.Uuid).Sanitize())
	doc.Set("storefront_uuid", kudakiredisearch.RedisearchText(cartItem.Item.Storefront.Uuid).Sanitize())
	doc.Set("item_name", cartItem.Item.Name)
	doc.Set("item_amount", cartItem.Item.Amount)
	doc.Set("item_unit", cartItem.Item.Unit)
	doc.Set("item_price", cartItem.Item.Price)
	doc.Set("item_description", cartItem.Item.Description)
	doc.Set("item_photo", cartItem.Item.Photo)
	doc.Set("item_rating", cartItem.Item.Rating)

	err := client.IndexOptions(redisearch.IndexingOptions{Replace: true}, doc)
	errorkit.ErrorHandled(err)
}
