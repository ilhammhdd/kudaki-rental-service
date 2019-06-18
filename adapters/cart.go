package adapters

import (
	"github.com/golang/protobuf/proto"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/kudaki-entities/events"
)

type AddCartItem struct{}

func (aci *AddCartItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.AddCartItemRequested

	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (aci *AddCartItem) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.CartItemAdded)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}

type DeleteCartItem struct{}

func (dci *DeleteCartItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.DeleteCartItemRequested

	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (dci *DeleteCartItem) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.CartItemDeleted)

	outByte, err := proto.Marshal(outEvent)
	errorkit.ErrorHandled(err)
	return outEvent.Uid, outByte
}

type UpdateCartItem struct{}

func (uci *UpdateCartItem) ParseIn(msg []byte) (proto.Message, bool) {
	var inEvent events.UpdateCartItemRequested
	if proto.Unmarshal(msg, &inEvent) == nil {
		return &inEvent, true
	}
	return nil, false
}

func (uci *UpdateCartItem) ParseOut(out proto.Message) (key string, message []byte) {
	outEvent := out.(*events.CartItemUpdated)

	outByte, err := proto.Marshal(out)
	errorkit.ErrorHandled(err)

	return outEvent.Uid, outByte
}
