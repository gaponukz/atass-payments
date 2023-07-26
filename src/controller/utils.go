package controller

import (
	"encoding/json"
	"net/http"
	"payments/src/dto"
)

func getCratePaymentDTO(request *http.Request) (dto.CratePaymentDTO, error) {
	var creds dto.CratePaymentDTO

	err := decodeRequestBody(request, &creds)
	return creds, err
}

func decodeRequestBody(request *http.Request, data interface{}) error {
	err := json.NewDecoder(request.Body).Decode(data)

	if err != nil {
		return err
	}

	return nil
}
