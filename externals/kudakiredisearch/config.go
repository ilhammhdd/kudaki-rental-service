package kudakiredisearch

import (
	"os"

	"github.com/RediSearch/redisearch-go/redisearch"
)

var RedisearchClient *redisearch.Client

func InitClient() {
	RedisearchClient = redisearch.NewClient(os.Getenv("REDISEARCH_SERVER"), "kudaki_rental_service")
	initSchemas()
}

func initSchemas() {
	cartSchema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewSortableNumericField("id")).
		AddField(redisearch.NewTextField("uuid")).
		AddField(redisearch.NewTextField("user_uuid")).
		AddField(redisearch.NewNumericField("total_item_price")).
		AddField(redisearch.NewNumericField("open"))
	RedisearchClient.CreateIndex(cartSchema)

	checkoutSchema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewSortableNumericField("id")).
		AddField(redisearch.NewTextField("uuid")).
		AddField(redisearch.NewTextField("cart_uuid")).
		AddField(redisearch.NewTextField("payment_method_uuid")).
		AddField(redisearch.NewSortableNumericField("issued_at"))
	RedisearchClient.CreateIndex(checkoutSchema)

	cartItemsSchema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewSortableNumericField("id")).
		AddField(redisearch.NewTextField("uuid")).
		AddField(redisearch.NewTextField("cart_uuid")).
		AddField(redisearch.NewTextField("item_uuid")).
		AddField(redisearch.NewNumericField("total_item")).
		AddField(redisearch.NewNumericField("price_total"))
	RedisearchClient.CreateIndex(cartItemsSchema)

	paymentMethodSchema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewSortableNumericField("id")).
		AddField(redisearch.NewTextField("uuid")).
		AddField(redisearch.NewTextField("name")).
		AddField(redisearch.NewTextField("description"))
	RedisearchClient.CreateIndex(paymentMethodSchema)
}
