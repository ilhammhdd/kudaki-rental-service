package adapters

import (
	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-entities/events"
)

type SubmitRental struct{}

func (sr *SubmitRental) ParseIn(msg []byte) (proto.Message, bool) {
	var in events.CheckoutRequested
	if err := proto.Unmarshal(msg, &in); err == nil {
		return &in, true
	}

	return nil, false
}

func (sr *SubmitRental) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.Checkedout)
	outEventByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outEventByte
}
