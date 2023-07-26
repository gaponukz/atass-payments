package outbox

import (
	"fmt"
	"payments/src/entities"
	"time"
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

func (s sendEventsService) Run() {
	for {
		payment, err := s.storage.Pop()
		if err != nil {
			time.Sleep(time.Second * 3)
			continue
		}

		err = s.notifier.Notify(payment)
		if err != nil {
			err = s.storage.Rollback(payment)
			if err != nil {
				fmt.Printf("Warning: Failed to rollback to storage: %v\n", err)
			}
		}

	}
}
