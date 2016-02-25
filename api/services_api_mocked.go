package api

import (
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO exclude from release compile

// GetServiceListMocked test mocked function
func GetServiceListMocked(t *testing.T, servicesIn *[]types.Service) *[]types.Service {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServicesService(cs)
	assert.Nil(err, "Couldn't load service service")
	assert.NotNil(ds, "Service service not instanced")

	// to json
	dIn, err := json.Marshal(servicesIn)
	assert.Nil(err, "Service test data corrupted")

	// call service
	cs.On("Get", "/v1/blueprint/services").Return(dIn, 200, nil)
	servicesOut, err := ds.GetServiceList()
	assert.Nil(err, "Error getting service list")
	assert.Equal(*servicesIn, servicesOut, "GetServiceList returned different services")

	return &servicesOut
}

// GetServiceMocked test mocked function
func GetServiceMocked(t *testing.T, service *types.Service) *types.Service {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServicesService(cs)
	assert.Nil(err, "Couldn't load service service")
	assert.NotNil(ds, "Service service not instanced")

	// to json
	dIn, err := json.Marshal(service)
	assert.Nil(err, "Service test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/blueprint/services/%s", service.Id)).Return(dIn, 200, nil)
	serviceOut, err := ds.GetService(service.Id)
	assert.Nil(err, "Error getting service")
	assert.Equal(*service, *serviceOut, "GetService returned different services")

	return serviceOut
}
