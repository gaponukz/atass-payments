package controller

import (
	"net/http"
	"payments/internal/application/dto"
	"payments/internal/domain/entities"
)

type paymentService interface {
	OnSuccessfulPayment(dto.CratePaymentDTO) (entities.Payment, error)
}

type controller struct {
	paymentService paymentService
}

func NewController(paymentService paymentService) controller {
	return controller{paymentService: paymentService}
}

func (c controller) ProcessPayment(responseWriter http.ResponseWriter, request *http.Request) {
	newPayment, err := getCratePaymentDTO(request)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	payment, err := c.paymentService.OnSuccessfulPayment(newPayment)
	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.Write([]byte(payment.ID))
}
