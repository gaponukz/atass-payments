package mocks

import (
	"fmt"
	"payments/internal/domain/entities"
)

type successfulEventNotifier struct {
	LastMessage entities.OutboxData
}

type unsuccessfulEventNotifier struct{}

func (s *successfulEventNotifier) Notify(p entities.OutboxData) error {
	s.LastMessage = p
	return nil
}

func (s unsuccessfulEventNotifier) Notify(p entities.OutboxData) error {
	return fmt.Errorf("balabla")
}

func NewUnsuccessfulEventNotifier() unsuccessfulEventNotifier {
	return unsuccessfulEventNotifier{}
}

func NewSuccessfulEventNotifier() *successfulEventNotifier {
	return &successfulEventNotifier{}
}
