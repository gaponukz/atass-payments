package outbox

import "payments/internal/domain/entities"

type popAbleStorage interface {
	PopPayment() (entities.OutboxData, error)
	PushBack(entities.OutboxData) error
}

type eventSender interface {
	Notify(entities.OutboxData) error
}

type sendEventsService struct {
	storage  popAbleStorage
	notifier eventSender
}

func NewSendEventsService(storage popAbleStorage, notifier eventSender) sendEventsService {
	return sendEventsService{storage: storage, notifier: notifier}
}

func (s sendEventsService) SendNewEvent() error {
	payment, err := s.storage.PopPayment()
	if err != nil {
		return err
	}

	err = s.notifier.Notify(payment)
	if err != nil {
		err = s.storage.PushBack(payment)
		if err != nil {
			return err
		}
	}

	return nil
}
