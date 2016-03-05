package api

import (
	"encoding/json"
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
