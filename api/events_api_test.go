package api

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEventServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewEventService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetEventList(t *testing.T) {
	eventsIn := testdata.GetEventData()
	GetEventListMocked(t, eventsIn)
	GetEventListFailErrMocked(t, eventsIn)
	GetEventListFailStatusMocked(t, eventsIn)
	GetEventListFailJSONMocked(t, eventsIn)
}

func TestGetSysEventList(t *testing.T) {
	eventsIn := testdata.GetEventData()
	GetSysEventListMocked(t, eventsIn)
	GetSysEventListFailErrMocked(t, eventsIn)
	GetSysEventListFailStatusMocked(t, eventsIn)
	GetSysEventListFailJSONMocked(t, eventsIn)
}
