package eventdriven

import (
	"testing"
)

func TestDeleteCartItem(t *testing.T) {
	t.Error("consume in event")
	t.Run("adapter", TestDeleteCartItemSub)
}

func TestDeleteCartItemSub(t *testing.T) {

}
