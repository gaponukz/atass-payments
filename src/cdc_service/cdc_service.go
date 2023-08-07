package cdc_service

import (
	"payments/src/errors"
	"time"
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
		if err == nil {
			continue
		}

		if err == errors.ErrStorageEmpty {
			time.Sleep(time.Second * 5)
			continue
		}

		time.Sleep(time.Second * 3)
	}
}
