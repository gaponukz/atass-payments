package mocks

import (
	"fmt"
	"payments/src/entities"
)

type successfulEventNotifier struct{}

type unsuccessfulEventNotifier struct{}

func (s successfulEventNotifier) Notify(p entities.OutboxData) error {
	return nil
}

func (s unsuccessfulEventNotifier) Notify(p entities.OutboxData) error {
	return fmt.Errorf("balabla")
}

func NewUnsuccessfulEventNotifier() unsuccessfulEventNotifier {
	return unsuccessfulEventNotifier{}
}

func NewSuccessfulEventNotifier() successfulEventNotifier {
	return successfulEventNotifier{}
}
