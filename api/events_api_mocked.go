package api

import (
	"encoding/json"
	// "fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO exclude from release compile

// GetEventListMocked test mocked function
func GetEventListMocked(t *testing.T, eventsIn *[]types.Event) *[]types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "Event service not instanced")

	// to json
	dIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "Event test data corrupted")

	// call service
	cs.On("Get", "/v1/audit/events").Return(dIn, 200, nil)
	eventsOut, err := ds.GetEventList()
	assert.Nil(err, "Error getting event list")
	assert.Equal(*eventsIn, eventsOut, "GetEventList returned different events")

	return &eventsOut
}

// GetSysEventListMocked test mocked function
func GetSysEventListMocked(t *testing.T, eventsIn *[]types.Event) *[]types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "Event service not instanced")

	// to json
	dIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "Event test data corrupted")

	// call service
	cs.On("Get", "/v1/audit/system_events").Return(dIn, 200, nil)
	eventsOut, err := ds.GetSysEventList()
	assert.Nil(err, "Error getting event list")
	assert.Equal(*eventsIn, eventsOut, "GetEventList returned different events")

	return &eventsOut
}
