package usecases

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/golang/protobuf/ptypes"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/ilhammhdd/kudaki-rental-service/entities"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/events"
)

type Checkout struct {
	DBO            entities.DBOperator
	CheckoutSchema *redisearch.Schema
}

func (rs *Checkout) Process(in proto.Message) (out proto.Message) {
	inEvent := in.(*events.CheckoutRequested)

	// insert cart detail to checkout
	checkoutUUID := uuid.New().String()
	issuedAt := TimeNowToDateTime()
	result, err := rs.DBO.Command("INSERT INTO checkouts(uuid,cart_uuid,issued_at) VALUES(?,?,?);", checkoutUUID, inEvent.Cart.Uuid, issuedAt)
	errorkit.ErrorHandled(err)
	lastInsertedID, err := result.LastInsertId()
	errorkit.ErrorHandled(err)
	// insert cart detail to checkout

	// indexing doc to redisearch
	rsClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), entities.Checkouts.String())
	err = rsClient.CreateIndex(rs.CheckoutSchema)
	errorkit.ErrorHandled(err)

	checkoutDoc := redisearch.NewDocument(checkoutUUID, 1.0)
	checkoutDoc.Set("id", lastInsertedID)
	checkoutDoc.Set("uuid", checkoutUUID)
	checkoutDoc.Set("cart_uuid", inEvent.Cart.Uuid)
	checkoutDoc.Set("issued_at", issuedAt)

	err = rsClient.IndexOptions(redisearch.DefaultIndexingOptions, checkoutDoc)
	errorkit.ErrorHandled(err)
	// indexing doc to redisearch

	// make out event for rental submitted
	var outEvent events.Checkedout
	outEvent.Uid = inEvent.Uid
	outEvent.EventStatus = &events.Status{
		HttpCode:  http.StatusOK,
		Timestamp: ptypes.TimestampNow()}
	outEvent.Cart = inEvent.Cart
	// make out event for rental submitted

	return &outEvent
}

type AddCartItem struct {
	DBO             entities.DBOperator
	CartItemsSchema *redisearch.Schema
}

func (aci AddCartItem) Process(in proto.Message) (out proto.Message) {
	inEvent := in.(*events.AddCartItemRequested)

	// init event
	var outEvent events.CartItemAdded
	outEvent.Uid = inEvent.Uid
	outEvent.EventStatus = &events.Status{}
	// init event

	// check if cart exists
	row, err := aci.DBO.QueryRow("SELECT id FROM carts WHERE uuid = ?;", inEvent.CartUuid)
	errorkit.ErrorHandled(err)

	var cartID uint64
	err = row.Scan(&cartID)
	if err == sql.ErrNoRows {
		outEvent.EventStatus.Errors = []string{"cart with the given uuid doesn't exists"}
		outEvent.EventStatus.Timestamp = ptypes.TimestampNow()
		outEvent.EventStatus.HttpCode = http.StatusNotFound

		return &outEvent
	}
	// check if cart exists

	return nil
}
