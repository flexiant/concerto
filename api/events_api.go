package api

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/utils"
	"time"
)

// Event stores an Concerto event item
type Event struct {
	Id          string    `json:"id" header:"ID"`
	Timestamp   time.Time `json:"timestamp" header:"TIMESTAMP"`
	Level       string    `json:"level" header:"LEVEL"`
	Header      string    `json:"header" header:"HEADER"`
	Description string    `json:"description" header:"DESCRIPTION"`
}

// EventService manages event operations
type EventService struct {
	concertoService utils.ConcertoService
}

// NewEventService returns a Concerto event service
func NewEventService(concertoService utils.ConcertoService) (*EventService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &EventService{
		concertoService: concertoService,
	}, nil
}

// GetEventList returns the list of events as an array of Event
func (cl *EventService) GetEventList() (events []Event, err error) {
	log.Debug("GetEventList")

	data, status, err := cl.concertoService.Get("/v1/audit/events")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {

	}

	if err = json.Unmarshal(data, &events); err != nil {
		return nil, err
	}

	return events, nil
}

// GetSysEventList returns the list of events as an array of Event
func (cl *EventService) GetSysEventList() (events []Event, err error) {
	log.Debug("GetEventList")

	data, status, err := cl.concertoService.Get("/v1/audit/system_events")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {

	}

	if err = json.Unmarshal(data, &events); err != nil {
		return nil, err
	}

	return events, nil
}
