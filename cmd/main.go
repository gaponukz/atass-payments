package main

import (
	"fmt"
	"net/http"
	"payments/src/controller"
	"payments/src/errors"
	"payments/src/notifier"
	"payments/src/outbox"
	"payments/src/storage"
	"payments/src/usecase"
)

func main() {
	paymentsDB := storage.NewJsonPaymentsStorage("payments.json")
	outboxDB := storage.NewJsonPaymentsStorage("outbox.json")
	paymentService := usecase.NewPaymentService(paymentsDB)
	serviceWithOutbox := outbox.NewSaveToOutboxDecorator(paymentService, outboxDB)
	controller := controller.NewController(serviceWithOutbox)
	sendEventsService := outbox.NewSendEventsService(outboxDB, notifier.NewTestNotifier())

	go func() {
		for {
			err := sendEventsService.SendNewEvent()
			if err != nil {
				if err == errors.ErrStorageEmpty {
					continue
				}

				fmt.Printf("Warning: %v", err)
			}
		}
	}()

	handler := http.NewServeMux()
	handler.HandleFunc("/processPayment", controller.ProcessPayment)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
