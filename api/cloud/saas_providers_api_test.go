package cloud

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSaasProviderServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewSaasProviderService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetSaasProviderList(t *testing.T) {
	saasProvidersIn := testdata.GetSaasProviderData()
	GetSaasProviderListMocked(t, saasProvidersIn)
	GetSaasProviderListFailErrMocked(t, saasProvidersIn)
	GetSaasProviderListFailStatusMocked(t, saasProvidersIn)
	GetSaasProviderListFailJSONMocked(t, saasProvidersIn)
}
