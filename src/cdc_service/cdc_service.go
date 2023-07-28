package cdc_service

import (
	"fmt"
	"payments/src/errors"
)

type sendEventsService interface {
	SendNewEvent() error
}

type cdcService struct {
	s sendEventsService
}

func NewCDCservice(s sendEventsService) cdcService {
	return cdcService{s: s}
}

func (c cdcService) Serve() {
	for {
		err := c.s.SendNewEvent()
		if err != nil {
			if err == errors.ErrStorageEmpty {
				continue
			}

			fmt.Println(err.Error())
		}
	}
}
