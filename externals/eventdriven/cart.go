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
	usecase := usecases.AddCartItem{
		DBO:              mysql.NewDBOperation(),
		ItemClient:       kudakiredisearch.Item,
		Sanitizer:        new(kudakiredisearch.RedisearchText),
		StorefrontClient: kudakiredisearch.Storefront}

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

	sanitizer := new(kudakiredisearch.RedisearchText)

	sanitizer.Set(cart.Uuid)
	sanitizedCartUUID := sanitizer.Sanitize()
	doc := redisearch.NewDocument(sanitizedCartUUID, 1.0)
	doc.Set("cart_uuid", sanitizedCartUUID)
	sanitizer.Set(cart.User.Uuid)
	doc.Set("user_uuid", sanitizer.Sanitize())
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

	sanitizer := new(kudakiredisearch.RedisearchText)

	sanitizer.Set(cartItem.Uuid)
	sanitizedCartItemUUID := sanitizer.Sanitize()
	doc := redisearch.NewDocument(sanitizedCartItemUUID, 1.0)
	doc.Set("cart_item_uuid", sanitizedCartItemUUID)
	doc.Set("cart_item_total_amount", cartItem.TotalAmount)
	doc.Set("cart_item_total_price", cartItem.TotalPrice)

	sanitizer.Set(cartItem.Cart.Uuid)
	doc.Set("cart_uuid", sanitizer.Sanitize())
	sanitizer.Set(cartItem.Cart.User.Uuid)
	doc.Set("user_uuid", sanitizer.Sanitize())
	doc.Set("cart_total_price", cartItem.Cart.TotalPrice)
	doc.Set("cart_total_items", cartItem.Cart.TotalItems)
	doc.Set("cart_open", 1)

	sanitizer.Set(cartItem.Item.Uuid)
	doc.Set("item_uuid", sanitizer.Sanitize())
	sanitizer.Set(cartItem.Item.Storefront.Uuid)
	doc.Set("storefront_uuid", sanitizer.Sanitize())
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

type DeleteCartItem struct{}

func (dci *DeleteCartItem) Work() interface{} {
	usecase := &usecases.DeleteCartItem{
		DBO:        mysql.NewDBOperation(),
		ItemClient: kudakiredisearch.Item,
		Sanitizer:  new(kudakiredisearch.RedisearchText),
	}

	ede := EventDrivenExternal{
		PostUsecaseExecutor: dci,
		eventDrivenAdapter:  new(adapters.DeleteCartItem),
		eventDrivenUsecase:  usecase,
		eventName:           events.RentalTopic_DELETE_CART_ITEM_REQUESTED.String(),
		inTopics:            []string{events.RentalTopic_DELETE_CART_ITEM_REQUESTED.String()},
		outTopic:            events.RentalTopic_CART_ITEM_DELETED.String()}
	ede.handle()
	return nil
}

func (dci *DeleteCartItem) ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message) {
	out := outEvent.(*events.CartItemDeleted)

	if out.EventStatus.HttpCode != http.StatusOK {
		return
	}

	dci.deleteCartItemFromDB(out)
	dci.deleteCartItemIndex(out)
	dci.updateCart(out.CartItem.Cart)
	dci.reIndexCart(out.CartItem.Cart)
}

func (dci *DeleteCartItem) deleteCartItemFromDB(outEvent *events.CartItemDeleted) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("DELETE FROM cart_items WHERE uuid=?;", outEvent.CartItem.Uuid)
	errorkit.ErrorHandled(err)
}

func (dci *DeleteCartItem) deleteCartItemIndex(outEvent *events.CartItemDeleted) {
	host := redisearch.NewSingleHostPool(os.Getenv("REDISEARCH_SERVER"))
	defer host.Close()
	conn := host.Get()
	defer conn.Close()

	sanitizer := new(kudakiredisearch.RedisearchText)

	sanitizer.Set(outEvent.CartItem.Uuid)
	_, err := conn.Do(
		"FT.DEL",
		kudakiredisearch.CartItem.Name(),
		sanitizer.Sanitize(),
		"DD")
	errorkit.ErrorHandled(err)
}

func (dci *DeleteCartItem) updateCart(cart *rental.Cart) {
	dbo := mysql.NewDBOperation()
	_, err := dbo.Command("UPDATE carts SET total_items=?,total_price=? WHERE uuid=?;", cart.TotalItems, cart.TotalPrice, cart.Uuid)
	errorkit.ErrorHandled(err)
}

func (dci *DeleteCartItem) reIndexCart(cart *rental.Cart) {
	client := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), kudakiredisearch.Cart.Name())
	client.CreateIndex(kudakiredisearch.Cart.Schema())

	sanitizer := new(kudakiredisearch.RedisearchText)

	sanitizer.Set(cart.Uuid)
	updateCartDoc := redisearch.NewDocument(sanitizer.Sanitize(), 1.0)
	updateCartDoc.Set("cart_total_price", cart.TotalPrice)
	updateCartDoc.Set("cart_total_items", cart.TotalItems)

	err := client.IndexOptions(redisearch.IndexingOptions{Partial: true, Replace: true}, updateCartDoc)
	errorkit.ErrorHandled(err)
}
