package logger

import "fmt"

type sendEventsService interface {
	SendNewEvent() error
}

type logSendEventsService struct {
	s sendEventsService
	l logger
}

func NewLogSendEventsServiceDecorator(s sendEventsService, l logger) logSendEventsService {
	return logSendEventsService{s: s, l: l}
}

func (s logSendEventsService) SendNewEvent() error {
	err := s.s.SendNewEvent()
	if err != nil {
		s.l.Error(fmt.Sprintf("Can not send new event: %v", err))
	}

	return err
}
