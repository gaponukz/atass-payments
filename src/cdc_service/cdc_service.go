package cdc_service

import (
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
		if err != nil {
			time.Sleep(time.Second * 10)
		}
	}
}
