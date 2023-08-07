package errors

import "errors"

var (
	ErrStorageEmpty    = errors.New("storage is empty")
	ErrPaymentNotValid = errors.New("payment not valid")
)
