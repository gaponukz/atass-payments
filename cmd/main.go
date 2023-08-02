package main

import (
	"fmt"
	"net/http"
	"payments/src/cdc_service"
	"payments/src/controller"
	"payments/src/logger"
	"payments/src/notifier"
	"payments/src/outbox"
	"payments/src/settings"
	"payments/src/storage"
	"payments/src/usecase"
)

func main() {
	config := settings.NewDotEnvSettings().Load()
	rabbitMQNotifier, err := notifier.NewRabbitMQNotifier(config.RabbitmqUrl)
	if err != nil {
		panic(err.Error())
	}

	defer rabbitMQNotifier.Close()

	paymentsDB, err := storage.NewPostgresUserStorage(storage.PostgresCredentials{
		Host:     config.PostgresHost,
		User:     config.PostgresUser,
		Password: config.PostgresPassword,
		Dbname:   config.PostgresDbname,
		Port:     config.PostgresPort,
		Sslmode:  config.PostgresSslmode,
	})
	if err != nil {
		panic(err.Error())
	}

	logging := logger.NewConsoleLogger()
	loggeredRabbitMQNotifier := logger.NewLoggingNotifierDecorator(rabbitMQNotifier, logging)
	loggeredPaymentsDB := logger.NewStorageLoggerDecorator(paymentsDB, logging)
	paymentService := logger.NewlogPaymentServiceDecorator(usecase.NewPaymentService(loggeredPaymentsDB), logging)
	controller := controller.NewController(paymentService)
	sendEventsService := logger.NewLogSendEventsServiceDecorator(outbox.NewSendEventsService(loggeredPaymentsDB, loggeredRabbitMQNotifier), logging)

	handler := http.NewServeMux()
	handler.HandleFunc("/processPayment", controller.ProcessPayment)

	server := http.Server{
		Addr:    ":9090",
		Handler: handler,
	}

	go cdc_service.NewCDCservice(sendEventsService).Serve()

	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
