package cloud

import (
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetCloudProviderListMocked test mocked function
func GetCloudProviderListMocked(t *testing.T, cloudProvidersIn *[]types.CloudProvider) *[]types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "CloudProvider test data corrupted")

	// call service
	cs.On("Get", "/v1/cloud/cloud_providers").Return(dIn, 200, nil)
	cloudProvidersOut, err := ds.GetCloudProviderList()
	assert.Nil(err, "Error getting cloudProvider list")
	assert.Equal(*cloudProvidersIn, cloudProvidersOut, "GetCloudProviderList returned different cloudProviders")

	return &cloudProvidersOut
}

// GetCloudProviderListFailErrMocked test mocked function
func GetCloudProviderListFailErrMocked(t *testing.T, cloudProvidersIn *[]types.CloudProvider) *[]types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "CloudProvider test data corrupted")

	// call service
	cs.On("Get", "/v1/cloud/cloud_providers").Return(dIn, 200, fmt.Errorf("Mocked error"))
	cloudProvidersOut, err := ds.GetCloudProviderList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudProvidersOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &cloudProvidersOut
}

// GetCloudProviderListFailStatusMocked test mocked function
func GetCloudProviderListFailStatusMocked(t *testing.T, cloudProvidersIn *[]types.CloudProvider) *[]types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "CloudProvider test data corrupted")

	// call service
	cs.On("Get", "/v1/cloud/cloud_providers").Return(dIn, 499, nil)
	cloudProvidersOut, err := ds.GetCloudProviderList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudProvidersOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &cloudProvidersOut
}

// GetCloudProviderListFailJSONMocked test mocked function
func GetCloudProviderListFailJSONMocked(t *testing.T, cloudProvidersIn *[]types.CloudProvider) *[]types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewCloudProviderService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "CloudProvider service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/cloud/cloud_providers").Return(dIn, 200, nil)
	cloudProvidersOut, err := ds.GetCloudProviderList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudProvidersOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &cloudProvidersOut
}
