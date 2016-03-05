package api

import (
	"encoding/json"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetSaasProviderListMocked test mocked function
func GetSaasProviderListMocked(t *testing.T, saasProvidersIn *[]types.SaasProvider) *[]types.SaasProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasProviderService(cs)
	assert.Nil(err, "Couldn't load saasProvider service")
	assert.NotNil(ds, "SaasProvider service not instanced")

	// to json
	dIn, err := json.Marshal(saasProvidersIn)
	assert.Nil(err, "SaasProvider test data corrupted")

	// call service
	cs.On("Get", "/v1/cloud/saas_providers").Return(dIn, 200, nil)
	saasProvidersOut, err := ds.GetSaasProviderList()
	assert.Nil(err, "Error getting saasProvider list")
	assert.Equal(*saasProvidersIn, saasProvidersOut, "GetSaasProviderList returned different saasProviders")

	return &saasProvidersOut
}
