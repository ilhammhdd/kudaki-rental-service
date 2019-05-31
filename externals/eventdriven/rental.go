package eventdriven

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/ilhammhdd/kudaki-rental-service/externals/kudakiredisearch"

	"github.com/ilhammhdd/kudaki-rental-service/externals/mysql"

	"github.com/ilhammhdd/kudaki-rental-service/usecases"

	"github.com/ilhammhdd/kudaki-entities/events"

	"github.com/ilhammhdd/kudaki-rental-service/adapters"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	"github.com/ilhammhdd/kudaki-rental-service/externals/kafka"
	"gopkg.in/Shopify/sarama.v1"
)

const TOTAL_CONSUMER_MEMBER = 5

type EventDrivenExternal struct {
	inTopics           []string
	eventName          string
	eventDrivenAdapter adapters.EventDrivenAdapter
	eventDrivenUsecase usecases.EventDrivenUsecase
	outTopic           string
}

func (edc *EventDrivenExternal) produce(key string, msg []byte) {
	prod := kafka.NewProduction()
	prod.Set(edc.outTopic)
	start := time.Now()
	partition, offset, err := prod.SyncProduce(key, msg)
	duration := time.Since(start)
	errorkit.ErrorHandled(err)

	log.Printf("produced %s : partition = %d, offset = %d, key = %s, duration = %f seconds", edc.outTopic, partition, offset, key, duration.Seconds())
}

func (edc *EventDrivenExternal) handle() {
	groupID := uuid.New().String()
	cl := adapters.ConsumerLog{EventName: edc.eventName}

	for i := 0; i < TOTAL_CONSUMER_MEMBER; i++ {
		consMember := kafka.NewConsumptionMember(groupID, edc.inTopics, sarama.OffsetNewest, edc.eventName, i)
		signals := make(chan os.Signal)
		signal.Notify(signals)

		safekit.Do(func() {
			defer close(consMember.Close)
		ConsLoop:
			for {
				select {
				case msg := <-consMember.Messages:
					if in, ok := edc.eventDrivenAdapter.ParseIn(msg.Value); ok {
						cl.Log(msg.Partition, msg.Offset, string(msg.Key))
						out := edc.eventDrivenUsecase.Process(in)
						outKey, outMsg := edc.eventDrivenAdapter.ParseOut(out)
						edc.produce(outKey, outMsg)
					}
				case errs := <-consMember.Errs:
					errorkit.ErrorHandled(errs)
				case <-signals:
					break ConsLoop
				}
			}
		})
	}
}

func Checkout() {
	usecase := &usecases.Checkout{
		DBO:            mysql.NewDBOperation(),
		CheckoutSchema: kudakiredisearch.CheckoutsSchema.Schema(),
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
		CartItemsSchema: kudakiredisearch.CartItemsSchema.Schema(),
		DBO:             mysql.NewDBOperation(),
	}

	edc := EventDrivenExternal{
		eventDrivenAdapter: new(adapters.AddCartItem),
		eventDrivenUsecase: usecase,
		eventName:          events.RentalTopic_ADD_CART_ITEM_REQUESTED.String(),
		inTopics:           []string{events.RentalTopic_ADD_CART_ITEM_REQUESTED.String()},
		outTopic:           events.RentalTopic_CART_ITEM_ADDED.String(),
	}
	edc.handle()
}
