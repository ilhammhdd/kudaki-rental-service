package eventdriven

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	"github.com/ilhammhdd/kudaki-rental-service/adapters"
	"github.com/ilhammhdd/kudaki-rental-service/externals/kafka"
	"github.com/ilhammhdd/kudaki-rental-service/usecases"
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
						out := edc.eventDrivenUsecase.Handle(in)
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
