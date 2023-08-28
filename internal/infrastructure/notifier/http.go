package notifier

import (
	"bytes"
	"encoding/json"
	"net/http"
	"payments/internal/domain/entities"
)

type httpNotifier struct {
	url string
}

func NewHttpNotifier(url string) httpNotifier {
	return httpNotifier{url: url}
}

func (r httpNotifier) Notify(payment entities.OutboxData) error {
	payload, err := json.Marshal(payment)
	if err != nil {
		return err
	}

	resp, err := http.Post(r.url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
