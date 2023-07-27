package outbox

import (
	"payments/src/entities"
)

type popAbleStorage interface {
	Pop() (entities.Payment, error)
	Rollback(entities.Payment) error
}

type eventSender interface {
	Notify(entities.Payment) error
}

type sendEventsService struct {
	storage  popAbleStorage
	notifier eventSender
}

func NewSendEventsService(storage popAbleStorage, notifier eventSender) sendEventsService {
	return sendEventsService{storage: storage, notifier: notifier}
}

func (s sendEventsService) SendNewEvent() error {
	payment, err := s.storage.Pop()
	if err != nil {
		return err
	}

	err = s.notifier.Notify(payment)
	if err != nil {
		err = s.storage.Rollback(payment)
		if err != nil {
			return err
		}
	}

	return nil
}
