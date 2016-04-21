package cloud

import (
	"encoding/json"
	"fmt"
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

// GetSaasProviderListFailErrMocked test mocked function
func GetSaasProviderListFailErrMocked(t *testing.T, saasProvidersIn *[]types.SaasProvider) *[]types.SaasProvider {

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
	cs.On("Get", "/v1/cloud/saas_providers").Return(dIn, 200, fmt.Errorf("Mocked error"))
	saasProvidersOut, err := ds.GetSaasProviderList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(saasProvidersOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &saasProvidersOut
}

// GetSaasProviderListFailStatusMocked test mocked function
func GetSaasProviderListFailStatusMocked(t *testing.T, saasProvidersIn *[]types.SaasProvider) *[]types.SaasProvider {

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
	cs.On("Get", "/v1/cloud/saas_providers").Return(dIn, 499, nil)
	saasProvidersOut, err := ds.GetSaasProviderList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(saasProvidersOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &saasProvidersOut
}

// GetSaasProviderListFailJSONMocked test mocked function
func GetSaasProviderListFailJSONMocked(t *testing.T, saasProvidersIn *[]types.SaasProvider) *[]types.SaasProvider {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasProviderService(cs)
	assert.Nil(err, "Couldn't load saasProvider service")
	assert.NotNil(ds, "SaasProvider service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/cloud/saas_providers").Return(dIn, 200, nil)
	saasProvidersOut, err := ds.GetSaasProviderList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(saasProvidersOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &saasProvidersOut
}
