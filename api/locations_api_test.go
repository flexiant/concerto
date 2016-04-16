package api

import (
	"testing"

	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewLocationServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewLocationService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetLocationList(t *testing.T) {
	locationsIn := testdata.GetLocationData()
	GetLocationListMocked(t, locationsIn)
	GetLocationListFailErrMocked(t, locationsIn)
	GetLocationListFailStatusMocked(t, locationsIn)
	GetLocationListFailJSONMocked(t, locationsIn)
}
