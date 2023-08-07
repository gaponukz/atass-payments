package outbox

import (
	"payments/internal/application/outbox"
	"payments/internal/domain/entities"
	"payments/internal/domain/errors"
	"payments/tests/mocks"
	"testing"
)

func TestOnSuccessfulEvent(t *testing.T) {
	db := mocks.NewQueueMockStorage()
	eventSender := mocks.NewSuccessfulEventNotifier()
	service := outbox.NewSendEventsService(db, eventSender)

	err := service.SendNewEvent()
	if err != nil {
		if err != errors.ErrStorageEmpty {
			t.Errorf("We got error: %v", err)
		}
	}

	if err == nil {
		t.Error("Can send event with empty storage")
	}

	err = db.PushBack(entities.OutboxData{PaymentID: "1", Passenger: entities.Passenger{ID: "11"}})
	if err != nil {
		t.Errorf("Error adding payment to mock: %v", err)
	}

	err = service.SendNewEvent()
	if err != nil {
		t.Error(err.Error())
	}

	if eventSender.LastMessage.PaymentID != "1" {
		t.Errorf("expect 1 got %s", eventSender.LastMessage.PaymentID)
	}

	if eventSender.LastMessage.Passenger.ID != "11" {
		t.Errorf("expect 11 got %s", eventSender.LastMessage.Passenger.ID)
	}

	if _, err = db.PopPayment(); err != errors.ErrStorageEmpty {
		t.Error("Payment was not removed after sending")
	}
}

func TestOnUnSuccessfulEvent(t *testing.T) {
	db := mocks.NewQueueMockStorage()
	service := outbox.NewSendEventsService(db, mocks.NewUnsuccessfulEventNotifier())

	err := db.PushBack(entities.OutboxData{PaymentID: "1", Passenger: entities.Passenger{ID: "11"}})
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

	if payment.Passenger.ID != "11" {
		t.Errorf("Rollback was unsuccessful, passenge id is %s", payment.Passenger.ID)
	}
}
