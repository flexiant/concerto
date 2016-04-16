package api

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCloudProviderServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewCloudProviderService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetCloudProviderList(t *testing.T) {
	cloudProvidersIn := testdata.GetCloudProviderData()
	GetCloudProviderListMocked(t, cloudProvidersIn)
	GetCloudProviderListFailErrMocked(t, cloudProvidersIn)
	GetCloudProviderListFailStatusMocked(t, cloudProvidersIn)
	GetCloudProviderListFailJSONMocked(t, cloudProvidersIn)
}
