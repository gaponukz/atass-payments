package outbox

import (
	"payments/src/entities"
	"payments/src/errors"
	"payments/src/outbox"
	"payments/tests/mocks"
	"testing"
)

func TestOnSuccessfulEvent(t *testing.T) {
	db := mocks.NewQueueMockStorage()
	service := outbox.NewSendEventsService(db, mocks.NewSuccessfulEventNotifier())

	err := service.SendNewEvent()
	if err != nil {
		if err != errors.ErrStorageEmpty {
			t.Errorf("We got error: %v", err)
		}
	}

	if err == nil {
		t.Error("Can send event with empty storage")
	}

	err = db.PushBack(entities.OutboxData{PaymentID: "1"})
	if err != nil {
		t.Errorf("Error adding payment to mock: %v", err)
	}

	err = service.SendNewEvent()
	if err != nil {
		t.Error(err.Error())
	}

	if _, err = db.PopPayment(); err != errors.ErrStorageEmpty {
		t.Error("Payment was not removed after sending")
	}
}

func TestOnUnSuccessfulEvent(t *testing.T) {
	db := mocks.NewQueueMockStorage()
	service := outbox.NewSendEventsService(db, mocks.NewUnsuccessfulEventNotifier())

	err := db.PushBack(entities.OutboxData{PaymentID: "1"})
	if err != nil {
		t.Errorf("Error adding payment to mock: %v", err)
	}

	err = service.SendNewEvent()
	if err != nil {
		t.Error(err.Error())
	}

	payment, err := db.PopPayment()
	if err != nil {
		t.Error(err.Error())
	}

	if payment.PaymentID != "1" {
		t.Errorf("Rollback was unsuccessful, payment id is %s", payment.PaymentID)
	}
}
