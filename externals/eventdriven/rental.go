package eventdriven

import (
	"os"
	"os/signal"

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

type EventDrivenConsumer struct {
	inTopics           []string
	eventName          string
	eventDrivenAdapter adapters.EventDrivenAdapter
	eventDrivenUsecase usecases.EventDrivenUsecase
	outTopic           string
}

func (edc *EventDrivenConsumer) produce(key string, msg []byte) {

}

func (edc *EventDrivenConsumer) handle() {
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

func SubmitRental() {
	usecase := &usecases.RentalSubmit{
		DBO:            mysql.NewDBOperation(),
		CheckoutSchema: kudakiredisearch.CheckoutsSchema.Schema(),
	}
	edc := EventDrivenConsumer{
		eventDrivenAdapter: new(adapters.SubmitRental),
		eventDrivenUsecase: usecase,
		eventName:          "RentalSubmissionRequested",
		inTopics:           []string{events.RentalTopic_name[int32(events.RentalTopic_CHECKOUT_REQUESTED)]},
		outTopic:           events.RentalTopic_name[int32(events.RentalTopic_CHECKEDOUT)],
	}
	edc.handle()
}
