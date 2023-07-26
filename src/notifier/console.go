package notifier

import (
	"fmt"
	"math/rand"
	"payments/src/entities"
	"time"
)

type testNotifier struct{}

func NewTestNotifier() testNotifier {
	return testNotifier{}
}

func (s testNotifier) Notify(payment entities.Payment) error {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(3)

	if randomNumber == 1 {
		return fmt.Errorf("blabla")
	}

	fmt.Println(payment)
	return nil
}
