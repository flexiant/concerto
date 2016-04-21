package audit

import (
	"encoding/json"
	"fmt"
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

// GetEventListFailErrMocked test mocked function
func GetEventListFailErrMocked(t *testing.T, eventsIn *[]types.Event) *[]types.Event {

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
	cs.On("Get", "/v1/audit/events").Return(dIn, 200, fmt.Errorf("Mocked error"))
	eventsOut, err := ds.GetEventList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(eventsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &eventsOut
}

// GetEventListFailStatusMocked test mocked function
func GetEventListFailStatusMocked(t *testing.T, eventsIn *[]types.Event) *[]types.Event {

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
	cs.On("Get", "/v1/audit/events").Return(dIn, 499, nil)
	eventsOut, err := ds.GetEventList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(eventsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &eventsOut
}

// GetEventListFailJSONMocked test mocked function
func GetEventListFailJSONMocked(t *testing.T, eventsIn *[]types.Event) *[]types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "Event service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/audit/events").Return(dIn, 200, nil)
	eventsOut, err := ds.GetEventList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(eventsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// GetSysEventListFailErrMocked test mocked function
func GetSysEventListFailErrMocked(t *testing.T, eventsIn *[]types.Event) *[]types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "SysEvent service not instanced")

	// to json
	dIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "SysEvent test data corrupted")

	// call service
	cs.On("Get", "/v1/audit/system_events").Return(dIn, 200, fmt.Errorf("Mocked error"))
	eventsOut, err := ds.GetSysEventList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(eventsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &eventsOut
}

// GetSysEventListFailStatusMocked test mocked function
func GetSysEventListFailStatusMocked(t *testing.T, eventsIn *[]types.Event) *[]types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "SysEvent service not instanced")

	// to json
	dIn, err := json.Marshal(eventsIn)
	assert.Nil(err, "SysEvent test data corrupted")

	// call service
	cs.On("Get", "/v1/audit/system_events").Return(dIn, 499, nil)
	eventsOut, err := ds.GetSysEventList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(eventsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &eventsOut
}

// GetSysEventListFailJSONMocked test mocked function
func GetSysEventListFailJSONMocked(t *testing.T, eventsIn *[]types.Event) *[]types.Event {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewEventService(cs)
	assert.Nil(err, "Couldn't load event service")
	assert.NotNil(ds, "SysEvent service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/audit/system_events").Return(dIn, 200, nil)
	eventsOut, err := ds.GetSysEventList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(eventsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &eventsOut
}
