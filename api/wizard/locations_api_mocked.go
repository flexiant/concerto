package wizard

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// GetLocationListMocked test mocked function
func GetLocationListMocked(t *testing.T, locationsIn *[]types.Location) *[]types.Location {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLocationService(cs)
	assert.Nil(err, "Couldn't load location service")
	assert.NotNil(ds, "Location service not instanced")

	// to json
	dIn, err := json.Marshal(locationsIn)
	assert.Nil(err, "Location test data corrupted")

	// call service
	cs.On("Get", "/v1/wizard/locations").Return(dIn, 200, nil)
	locationsOut, err := ds.GetLocationList()
	assert.Nil(err, "Error getting location list")
	assert.Equal(*locationsIn, locationsOut, "GetLocationList returned different locations")

	return &locationsOut
}

// GetLocationListFailErrMocked test mocked function
func GetLocationListFailErrMocked(t *testing.T, locationsIn *[]types.Location) *[]types.Location {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLocationService(cs)
	assert.Nil(err, "Couldn't load location service")
	assert.NotNil(ds, "Location service not instanced")

	// to json
	dIn, err := json.Marshal(locationsIn)
	assert.Nil(err, "Location test data corrupted")

	// call service
	cs.On("Get", "/v1/wizard/locations").Return(dIn, 200, fmt.Errorf("Mocked error"))
	locationsOut, err := ds.GetLocationList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(locationsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &locationsOut
}

// GetLocationListFailStatusMocked test mocked function
func GetLocationListFailStatusMocked(t *testing.T, locationsIn *[]types.Location) *[]types.Location {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLocationService(cs)
	assert.Nil(err, "Couldn't load location service")
	assert.NotNil(ds, "Location service not instanced")

	// to json
	dIn, err := json.Marshal(locationsIn)
	assert.Nil(err, "Location test data corrupted")

	// call service
	cs.On("Get", "/v1/wizard/locations").Return(dIn, 499, nil)
	locationsOut, err := ds.GetLocationList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(locationsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &locationsOut
}

// GetLocationListFailJSONMocked test mocked function
func GetLocationListFailJSONMocked(t *testing.T, locationsIn *[]types.Location) *[]types.Location {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLocationService(cs)
	assert.Nil(err, "Couldn't load location service")
	assert.NotNil(ds, "Location service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/wizard/locations").Return(dIn, 200, nil)
	locationsOut, err := ds.GetLocationList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(locationsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &locationsOut
}
