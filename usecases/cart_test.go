package usecases_test

import (
	"reflect"
	"testing"

	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-externals/mysql"
	"github.com/ilhammhdd/kudaki-rental-service/usecases"
)

type deleteCartItemTestCase struct {
	inEvent          *events.DeleteCartItemRequested
	outEvent         *events.CartItemDeleted
	expectedOutEvent *events.CartItemDeleted
}

var deleteCartItemTestCases = []deleteCartItemTestCase{
	{},
}

func TestDeleteCartItem(t *testing.T) {
	handler := usecases.DeleteCartItem{
		DBO: mysql.NewDBOperation()}

	for _, testCase := range deleteCartItemTestCases {
		t.Log(testCase)
		existedCartItem, _ := handler.CartItemExists(testCase.inEvent.CartItemUuid)
		if !reflect.DeepEqual(existedCartItem, testCase.expectedOutEvent.CartItem) {
			t.Error("retrieved CartItem from DB doesn't match with expected CartItem inside OutEvent")
		}
	}
}
