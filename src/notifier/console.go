package notifier

import (
	"fmt"
	"payments/src/entities"
)

type testNotifier struct{}

func NewTestNotifier() testNotifier {
	return testNotifier{}
}

func (s testNotifier) Notify(payment entities.Payment) error {
	fmt.Println(payment)
	return nil
}
