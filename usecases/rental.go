package usecases

import (
	"net/http"
	"os"

	"github.com/ilhammhdd/kudaki-entities/kudakiredisearch"

	"github.com/golang/protobuf/ptypes"

	"github.com/RediSearch/redisearch-go/redisearch"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/events"
)

type Checkout struct {
	DBO                DBOperator
	CartCheckoutClient kudakiredisearch.RedisClient
}

func (rs *Checkout) Handle(in proto.Message) (out proto.Message) {
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
	rsClient := redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), rs.CartCheckoutClient.Name())
	err = rsClient.CreateIndex(nil)
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
