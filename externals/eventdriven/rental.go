package eventdriven

import (
	"os"
	"os/signal"

	"github.com/ilhammhdd/kudaki-rental-service/adapters"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	"github.com/ilhammhdd/kudaki-rental-service/externals/kafka"
	"gopkg.in/Shopify/sarama.v1"
)

const TOTAL_CONSUMER_MEMBER = 5

type EventDrivenConsumer struct {
	topics             []string
	eventName          string
	eventDrivenAdapter adapters.EventDrivenAdapter
}

func (edc *EventDrivenConsumer) consume() {
	groupID := uuid.New().String()

	for i := 0; i < TOTAL_CONSUMER_MEMBER; i++ {
		consMember := kafka.NewConsumptionMember(groupID, edc.topics, sarama.OffsetNewest, edc.eventName, i)
		signals := make(chan os.Signal)
		signal.Notify(signals)

		safekit.Do(func() {
			defer close(consMember.Close)
		ConsLoop:
			for {
				select {
				case msg := <-consMember.Messages:
					if edc.eventDrivenAdapter.Adapt(msg.Value) {
						edc.eventDrivenAdapter.Log(msg.Partition, msg.Offset, string(msg.Key))
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

func SubmitRental() {
	edc := EventDrivenConsumer{
		eventDrivenAdapter: new(adapters.SubmitRental),
		eventName:          "RentalSubmissionRequested",
		topics:             []string{""},
	}
	edc.consume()
}
