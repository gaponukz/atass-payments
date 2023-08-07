package dto

import "payments/internal/domain/entities"

type CratePaymentDTO struct {
	Amount    float64            `json:"amount"`
	RouteID   entities.RouteID   `json:"routeId"`
	Passenger entities.Passenger `json:"passenger"`
}