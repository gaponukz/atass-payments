package outbox

import (
	"fmt"
	"payments/src/entities"
	"payments/src/errors"
)

type popAbleStorage interface {
	Pop() (entities.Payment, error)
}

type popFromOutboxService struct {
	storage popAbleStorage
}

func (s popFromOutboxService) WaitForNew() error {
	for {
		payment, err := s.storage.Pop()
		if err != nil {
			if err != errors.ErrStorageEmpty {
				fmt.Printf("Warning: Failed to pop from storage: %v\n", err)
			}
		}

		fmt.Println(payment)
	}
}
