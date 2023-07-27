package usecase

import (
	"payments/src/dto"
	"payments/src/entities"
	"payments/src/usecase"
	"payments/tests/mocks"
	"testing"
)

func TestOnSuccessfulPayment(t *testing.T) {
	const expectedAmount = 100
	const expectedRouteID = "123"
	expectedPassanger := entities.Passanger{ID: "321"}
	db := mocks.NewQueueMockStorage()
	service := usecase.NewPaymentService(db)

	payment, err := service.OnSuccessfulPayment(dto.CratePaymentDTO{
		Amount:    expectedAmount,
		RouteID:   expectedRouteID,
		Passanger: expectedPassanger,
	})
	if err != nil {
		t.Error(err.Error())
	}

	if payment.ID == "" {
		t.Error("Payment ID should not be empty")
	}

	if payment.RouteID != expectedRouteID {
		t.Errorf("route id expected %s, got %s", expectedRouteID, payment.RouteID)
	}

	if payment.Passanger.ID != expectedPassanger.ID {
		t.Errorf("passanger id expected %s, got %s", expectedPassanger.ID, payment.Passanger.ID)
	}

	paymentFromStorage, err := db.Pop()
	if err != nil {
		t.Error(err.Error())
	}

	if paymentFromStorage.ID != payment.ID {
		t.Errorf("Payment ID expected %s, got %s", payment.ID, paymentFromStorage.ID)
	}
}