package eventdriven

import (
	"github.com/ilhammhdd/kudaki-rental-service/usecases"

	"github.com/ilhammhdd/kudaki-rental-service/externals/mysql"

	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-entities/kudakiredisearch"

	"github.com/ilhammhdd/kudaki-rental-service/adapters"
)

func Checkout() {
	usecase := &usecases.Checkout{
		DBO:                mysql.NewDBOperation(),
		CartCheckoutClient: kudakiredisearch.CartCheckout,
	}
	edc := EventDrivenExternal{
		eventDrivenAdapter: new(adapters.Checkout),
		eventDrivenUsecase: usecase,
		eventName:          events.RentalTopic_CHECKOUT_REQUESTED.String(),
		inTopics:           []string{events.RentalTopic_name[int32(events.RentalTopic_CHECKOUT_REQUESTED)]},
		outTopic:           events.RentalTopic_name[int32(events.RentalTopic_CHECKEDOUT)],
	}
	edc.handle()
}

func AddCartItem() {
	usecase := &usecases.AddCartItem{
		CartClient:     kudakiredisearch.Cart,
		CartItemClient: kudakiredisearch.CartItem,
		ItemClient:     kudakiredisearch.Item,
		DBO:            mysql.NewDBOperation()}

	edc := EventDrivenExternal{
		eventDrivenAdapter: new(adapters.AddCartItem),
		eventDrivenUsecase: usecase,
		eventName:          events.RentalTopic_ADD_CART_ITEM_REQUESTED.String(),
		inTopics:           []string{events.RentalTopic_ADD_CART_ITEM_REQUESTED.String()},
		outTopic:           events.RentalTopic_CART_ITEM_ADDED.String()}
	edc.handle()
}
