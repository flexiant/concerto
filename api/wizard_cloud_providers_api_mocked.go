package api

import (
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetWizCloudProviderListMocked test mocked function
func GetWizCloudProviderListMocked(t *testing.T, cloudProvidersIn *[]types.CloudProvider, AppID string, LocID string) *[]types.CloudProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWizCloudProvidersService(cs)
	assert.Nil(err, "Couldn't load cloudProvider service")
	assert.NotNil(ds, "WizCloudProvider service not instanced")

	// to json
	dIn, err := json.Marshal(cloudProvidersIn)
	assert.Nil(err, "WizCloudProvider test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/wizard/cloud_providers?app_id=%s&location_id=%s", AppID, LocID)).Return(dIn, 200, nil)
	cloudProvidersOut, err := ds.GetWizCloudProviderList(AppID, LocID)
	assert.Nil(err, "Error getting cloudProvider list")
	assert.Equal(*cloudProvidersIn, cloudProvidersOut, "GetWizCloudProviderList returned different cloudProviders")

	return &cloudProvidersOut
}
