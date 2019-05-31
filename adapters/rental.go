package adapters

import (
	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/kudaki-entities/events"
)

type Checkout struct{}

func (sr *Checkout) ParseIn(msg []byte) (proto.Message, bool) {
	var in events.CheckoutRequested
	if err := proto.Unmarshal(msg, &in); err == nil {
		return &in, true
	}

	return nil, false
}

func (sr *Checkout) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.Checkedout)
	outEventByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outEventByte
}

type AddCartItem struct{}

func (aci AddCartItem) ParseIn(msg []byte) (proto.Message, bool) {
	var in events.AddCartItemRequested
	if err := proto.Unmarshal(msg, &in); err == nil {
		return &in, true
	}

	return nil, false
}

func (aci AddCartItem) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.CartItemAdded)
	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
