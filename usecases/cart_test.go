package usecases_test

import (
	"log"
	"testing"

	"github.com/ilhammhdd/kudaki-entities/store"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/rental"

	"github.com/ilhammhdd/kudaki-externals/mysql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/user"

	"github.com/ilhammhdd/kudaki-rental-service/usecases"
)

// var deleteCartItemTests = []struct {
// 	cartItemUUID   string
// 	kudakiToken    string
// 	cartItemExists bool
// }{
// 	{"303fdd6e-de1b-4e9f-b6cd-796583e2ffe9",
// 		"eyJhbGciOiJFQ0RTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ2ZXJpZmllZCBLdWRha2kuaWQgdXNlciIsImlzcyI6Ikt1ZGFraS5pZCB1c2VyIHNlcnZpY2UiLCJpYXQiOjE1NjA2NTk2NTg2ODEsImV4cCI6MTU2NTkxNTY1ODY4MSwiY2xhaW1zIjp7InByb2ZpbGUiOnsiZnVsbF9uYW1lIjoiTXVoYW1tYWQgSWxoYW0iLCJwaG90byI6ImltZ3VyLmNvbS9iZXR1bCIsInJlcHV0YXRpb24iOjAsInVzZXJfdXVpZCI6ImJmYTMxMzcwLTViNjMtNDAzYi04NGE0LTZhOTllNDQ2ZDE1YiIsInV1aWQiOiI4NTdmN2MxNS0yYzFhLTQ3NGUtODEzMi1kMjkzMWU0M2I4ODMifSwidXNlciI6eyJhY2NvdW50X3R5cGUiOiJOQVRJVkUiLCJlbWFpbCI6Im1pbGhhbTkzOUBnbWFpbC5jb20iLCJwaG9uZV9udW1iZXIiOiIwODIxNjY1NjIyNzkiLCJyb2xlIjoiVVNFUiIsInV1aWQiOiJiZmEzMTM3MC01YjYzLTQwM2ItODRhNC02YTk5ZTQ0NmQxNWIifX19.eyJoYXNoZWQiOiJINnRMQ3dNQS8waFFYZHlwVytwTHRhVThGV1htdTVyR1REeEFTNFlsWnJzPSIsInIiOjIyOTM2MTY0MTA4MDA2MTM4NTM3NjQ5MjE4NzY4ODgyMjI2MjcyNTY0OTY3ODY5MDc3OTUzOTEzODMxNDEwNTE2Nzk5ODM0NjU4ODAxLCJzIjo4OTAwNDI4MTQ3NzI1MjExMTI2ODg2MjI3Mjk3MzAwNDExMjAwNjAwMjk2MjgzOTIyOTA1NjAyNzQ2MzIwMjc4MTcyMzkwNzY5MzM3OX0=",
// 		true},
// 	{"303fdd6e-de1b-4e9f-b6cd-796583e2ffe9",
// 		"",
// 		true},
// 	{"",
// 		"eyJhbGciOiJFQ0RTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ2ZXJpZmllZCBLdWRha2kuaWQgdXNlciIsImlzcyI6Ikt1ZGFraS5pZCB1c2VyIHNlcnZpY2UiLCJpYXQiOjE1NjA2NTk2NTg2ODEsImV4cCI6MTU2NTkxNTY1ODY4MSwiY2xhaW1zIjp7InByb2ZpbGUiOnsiZnVsbF9uYW1lIjoiTXVoYW1tYWQgSWxoYW0iLCJwaG90byI6ImltZ3VyLmNvbS9iZXR1bCIsInJlcHV0YXRpb24iOjAsInVzZXJfdXVpZCI6ImJmYTMxMzcwLTViNjMtNDAzYi04NGE0LTZhOTllNDQ2ZDE1YiIsInV1aWQiOiI4NTdmN2MxNS0yYzFhLTQ3NGUtODEzMi1kMjkzMWU0M2I4ODMifSwidXNlciI6eyJhY2NvdW50X3R5cGUiOiJOQVRJVkUiLCJlbWFpbCI6Im1pbGhhbTkzOUBnbWFpbC5jb20iLCJwaG9uZV9udW1iZXIiOiIwODIxNjY1NjIyNzkiLCJyb2xlIjoiVVNFUiIsInV1aWQiOiJiZmEzMTM3MC01YjYzLTQwM2ItODRhNC02YTk5ZTQ0NmQxNWIifX19.eyJoYXNoZWQiOiJINnRMQ3dNQS8waFFYZHlwVytwTHRhVThGV1htdTVyR1REeEFTNFlsWnJzPSIsInIiOjIyOTM2MTY0MTA4MDA2MTM4NTM3NjQ5MjE4NzY4ODgyMjI2MjcyNTY0OTY3ODY5MDc3OTUzOTEzODMxNDEwNTE2Nzk5ODM0NjU4ODAxLCJzIjo4OTAwNDI4MTQ3NzI1MjExMTI2ODg2MjI3Mjk3MzAwNDExMjAwNjAwMjk2MjgzOTIyOTA1NjAyNzQ2MzIwMjc4MTcyMzkwNzY5MzM3OX0=",
// 		false},
// 	{"a11e648e-2c6b-4205-92c8-1e456e9537e2",
// 		"eyJhbGciOiJFQ0RTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ2ZXJpZmllZCBLdWRha2kuaWQgdXNlciIsImlzcyI6Ikt1ZGFraS5pZCB1c2VyIHNlcnZpY2UiLCJpYXQiOjE1NjA2NTk2NTg2ODEsImV4cCI6MTU2NTkxNTY1ODY4MSwiY2xhaW1zIjp7InByb2ZpbGUiOnsiZnVsbF9uYW1lIjoiTXVoYW1tYWQgSWxoYW0iLCJwaG90byI6ImltZ3VyLmNvbS9iZXR1bCIsInJlcHV0YXRpb24iOjAsInVzZXJfdXVpZCI6ImJmYTMxMzcwLTViNjMtNDAzYi04NGE0LTZhOTllNDQ2ZDE1YiIsInV1aWQiOiI4NTdmN2MxNS0yYzFhLTQ3NGUtODEzMi1kMjkzMWU0M2I4ODMifSwidXNlciI6eyJhY2NvdW50X3R5cGUiOiJOQVRJVkUiLCJlbWFpbCI6Im1pbGhhbTkzOUBnbWFpbC5jb20iLCJwaG9uZV9udW1iZXIiOiIwODIxNjY1NjIyNzkiLCJyb2xlIjoiVVNFUiIsInV1aWQiOiJiZmEzMTM3MC01YjYzLTQwM2ItODRhNC02YTk5ZTQ0NmQxNWIifX19.eyJoYXNoZWQiOiJINnRMQ3dNQS8waFFYZHlwVytwTHRhVThGV1htdTVyR1REeEFTNFlsWnJzPSIsInIiOjIyOTM2MTY0MTA4MDA2MTM4NTM3NjQ5MjE4NzY4ODgyMjI2MjcyNTY0OTY3ODY5MDc3OTUzOTEzODMxNDEwNTE2Nzk5ODM0NjU4ODAxLCJzIjo4OTAwNDI4MTQ3NzI1MjExMTI2ODg2MjI3Mjk3MzAwNDExMjAwNjAwMjk2MjgzOTIyOTA1NjAyNzQ2MzIwMjc4MTcyMzkwNzY5MzM3OX0=",
// 		false},
// 	{`303fdd6e\-de1b\-4e9f\-b6cd\-796583e2ffe9`,
// 		"eyJhbGciOiJFQ0RTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ2ZXJpZmllZCBLdWRha2kuaWQgdXNlciIsImlzcyI6Ikt1ZGFraS5pZCB1c2VyIHNlcnZpY2UiLCJpYXQiOjE1NjA2NTk2NTg2ODEsImV4cCI6MTU2NTkxNTY1ODY4MSwiY2xhaW1zIjp7InByb2ZpbGUiOnsiZnVsbF9uYW1lIjoiTXVoYW1tYWQgSWxoYW0iLCJwaG90byI6ImltZ3VyLmNvbS9iZXR1bCIsInJlcHV0YXRpb24iOjAsInVzZXJfdXVpZCI6ImJmYTMxMzcwLTViNjMtNDAzYi04NGE0LTZhOTllNDQ2ZDE1YiIsInV1aWQiOiI4NTdmN2MxNS0yYzFhLTQ3NGUtODEzMi1kMjkzMWU0M2I4ODMifSwidXNlciI6eyJhY2NvdW50X3R5cGUiOiJOQVRJVkUiLCJlbWFpbCI6Im1pbGhhbTkzOUBnbWFpbC5jb20iLCJwaG9uZV9udW1iZXIiOiIwODIxNjY1NjIyNzkiLCJyb2xlIjoiVVNFUiIsInV1aWQiOiJiZmEzMTM3MC01YjYzLTQwM2ItODRhNC02YTk5ZTQ0NmQxNWIifX19.eyJoYXNoZWQiOiJINnRMQ3dNQS8waFFYZHlwVytwTHRhVThGV1htdTVyR1REeEFTNFlsWnJzPSIsInIiOjIyOTM2MTY0MTA4MDA2MTM4NTM3NjQ5MjE4NzY4ODgyMjI2MjcyNTY0OTY3ODY5MDc3OTUzOTEzODMxNDEwNTE2Nzk5ODM0NjU4ODAxLCJzIjo4OTAwNDI4MTQ3NzI1MjExMTI2ODg2MjI3Mjk3MzAwNDExMjAwNjAwMjk2MjgzOTIyOTA1NjAyNzQ2MzIwMjc4MTcyMzkwNzY5MzM3OX0=",
// 		false},
// }

var deleteCartItemTests = []struct {
	cartItem               *rental.CartItem
	kudakiToken            string
	cartItemExpectedExists bool
	cartExpectedExists     bool
	itemExpectedExists     bool
}{
	{
		cartItem:               mockCartItems[0],
		kudakiToken:            "eyJhbGciOiJFQ0RTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ2ZXJpZmllZCBLdWRha2kuaWQgdXNlciIsImlzcyI6Ikt1ZGFraS5pZCB1c2VyIHNlcnZpY2UiLCJpYXQiOjE1NjA2NTk2NTg2ODEsImV4cCI6MTU2NTkxNTY1ODY4MSwiY2xhaW1zIjp7InByb2ZpbGUiOnsiZnVsbF9uYW1lIjoiTXVoYW1tYWQgSWxoYW0iLCJwaG90byI6ImltZ3VyLmNvbS9iZXR1bCIsInJlcHV0YXRpb24iOjAsInVzZXJfdXVpZCI6ImJmYTMxMzcwLTViNjMtNDAzYi04NGE0LTZhOTllNDQ2ZDE1YiIsInV1aWQiOiI4NTdmN2MxNS0yYzFhLTQ3NGUtODEzMi1kMjkzMWU0M2I4ODMifSwidXNlciI6eyJhY2NvdW50X3R5cGUiOiJOQVRJVkUiLCJlbWFpbCI6Im1pbGhhbTkzOUBnbWFpbC5jb20iLCJwaG9uZV9udW1iZXIiOiIwODIxNjY1NjIyNzkiLCJyb2xlIjoiVVNFUiIsInV1aWQiOiJiZmEzMTM3MC01YjYzLTQwM2ItODRhNC02YTk5ZTQ0NmQxNWIifX19.eyJoYXNoZWQiOiJINnRMQ3dNQS8waFFYZHlwVytwTHRhVThGV1htdTVyR1REeEFTNFlsWnJzPSIsInIiOjIyOTM2MTY0MTA4MDA2MTM4NTM3NjQ5MjE4NzY4ODgyMjI2MjcyNTY0OTY3ODY5MDc3OTUzOTEzODMxNDEwNTE2Nzk5ODM0NjU4ODAxLCJzIjo4OTAwNDI4MTQ3NzI1MjExMTI2ODg2MjI3Mjk3MzAwNDExMjAwNjAwMjk2MjgzOTIyOTA1NjAyNzQ2MzIwMjc4MTcyMzkwNzY5MzM3OX0=",
		cartItemExpectedExists: true,
		cartExpectedExists:     true,
		itemExpectedExists:     true,
	},
}

func TestDeleteCartItem(t *testing.T) {
	mysql.OpenDB("tcp(178.62.107.160:3306)", "root", "mysqlrocks", "kudaki_rental")

	handler := &usecases.DeleteCartItem{DBO: mysql.NewDBOperation()}
	log.Println(handler)

	// for _, testCase := range  {
	// 	existedCartItem, ok := handler.CartItemExists(testCase.cartItemUUID)
	// 	if ok && testCase.cartItemExists {
	// 		log.Println(existedCartItem)
	// 	} else if !ok && !testCase.cartItemExists {

	// 	} else {
	// 		t.Errorf("cart item existence not matched : given = %v, expected = %v", ok, testCase.cartItemExists)
	// 	}

	// 	existedCart, ok := handler.CartExists(cartUUID)

	// }
}

var mockItemsUUIDs []string = []string{
	"f9a7ac41-d0f5-4704-82d6-ba38033a1160",
	"a3d8e979-bc68-4060-82e3-e74d2bb9ccc6",
	"a39b2474-ef79-4064-b9ce-7864bbd38547"}

var mockCartsUUIDs []string = []string{
	"52c1f9c9-9673-4c51-b6fa-0d2dfc0d12c2",
	"f03df828-618e-4b36-955c-480f39e9f0bc"}

var mockCartItemsUUIDs []string = []string{
	"2bba85c1-e9e4-4bda-8f49-db4bd9fb10b0",
	"8e514eb1-73e8-400b-af3c-a546db35a540",
	"1d6f37b5-53e3-4456-b392-0017a82f6499"}

var mockCarts []*rental.Cart = []*rental.Cart{
	&rental.Cart{
		Open:       true,
		TotalItems: 20,
		TotalPrice: 40000,
		User:       &user.User{Uuid: "845b04c6-259d-4ddb-81cb-b386e4d9f86d"},
		Uuid:       mockCartsUUIDs[0]},
	&rental.Cart{
		Open:       true,
		TotalItems: 10,
		TotalPrice: 20000,
		User:       &user.User{Uuid: "a3d8e979-bc68-4060-82e3-e74d2bb9ccc6"},
		Uuid:       mockCartsUUIDs[1]},
}

var mockItems []*store.Item = []*store.Item{
	&store.Item{
		Amount:      200,
		Description: "nice carrier",
		Name:        "carrier",
		Photo:       "https://google.com",
		Price:       1000,
		Rating:      2.3,
		Storefront:  &store.Storefront{Uuid: "52c1f9c9-9673-4c51-b6fa-0d2dfc0d12c2"},
		Unit:        "unit",
		Uuid:        mockItemsUUIDs[0]},
	&store.Item{
		Amount:      145,
		Description: "nice headlamp",
		Name:        "headlamp",
		Photo:       "https://google.com",
		Price:       2000,
		Rating:      3.4,
		Storefront:  &store.Storefront{Uuid: "2bba85c1-e9e4-4bda-8f49-db4bd9fb10b0"},
		Unit:        "piece",
		Uuid:        mockItemsUUIDs[1]},
	&store.Item{
		Amount:      54,
		Description: "nice stick",
		Name:        "stick",
		Photo:       "https://google.com",
		Price:       500,
		Rating:      4.4,
		Storefront:  &store.Storefront{Uuid: uuid.New().String()},
		Unit:        "piece",
		Uuid:        mockItemsUUIDs[1]},
}

var mockCartItems []*rental.CartItem = []*rental.CartItem{
	&rental.CartItem{
		Cart:        mockCarts[0],
		Item:        mockItems[0],
		TotalAmount: 10,
		TotalPrice:  uint32(mockItems[0].Price * 10),
		Uuid:        uuid.New().String()},
	&rental.CartItem{
		Cart:        mockCarts[0],
		Item:        mockItems[2],
		TotalAmount: 40,
		TotalPrice:  uint32(mockItems[0].Price * 40),
		Uuid:        uuid.New().String()},
	&rental.CartItem{
		Cart:        mockCarts[1],
		Item:        mockItems[1],
		TotalAmount: 10,
		TotalPrice:  uint32(mockItems[0].Price * 10),
		Uuid:        uuid.New().String()},
}

func populateForDeleteCartItem() {
	mysql.OpenDB("tcp(178.62.107.160:3306)", "root", "mysqlrocks", "kudaki_rental")
	dbo := mysql.NewDBOperation()

	for idx, cartItem := range mockCartItems {
		log.Println(idx)
		_, err := dbo.Command("INSERT INTO cart_items(cart_uuid,item_uuid,total_item,total_price,uuid) VALUES(?,?,?,?,?);",
			cartItem.Cart.Uuid, cartItem.Item.Uuid, cartItem.TotalAmount, cartItem.TotalPrice, cartItem.Uuid)
		errorkit.ErrorHandled(err)
		_, err = dbo.Command("INSERT INTO items(amount,description,name,photo,price,rating,storefront_uuid,unit,uuid) VALUES(?,?,?,?,?,?,?,?,?);",
			cartItem.Item.Amount,
			cartItem.Item.Description,
			cartItem.Item.Name,
			cartItem.Item.Photo,
			cartItem.Item.Price,
			cartItem.Item.Rating,
			cartItem.Item.Storefront.Uuid,
			cartItem.Item.Unit,
			cartItem.Item.Uuid)
		errorkit.ErrorHandled(err)

		var open int
		if cartItem.Cart.Open {
			open = 1
		} else {
			open = 0
		}
		_, err = dbo.Command("INSERT INTO carts(open,total_items,total_price,user_uuid,uuid) VALUES(?,?,?,?,?);",
			open,
			cartItem.Cart.TotalItems,
			cartItem.Cart.TotalPrice,
			cartItem.Cart.User.Uuid,
			cartItem.Cart.Uuid)
		errorkit.ErrorHandled(err)
	}
}
