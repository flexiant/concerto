package wizard

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWizCloudProviderServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewWizCloudProvidersService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetWizCloudProviderList(t *testing.T) {
	AppID := "fakeAppID"
	LocID := "fakeLocID"
	cloudProvidersIn := testdata.GetCloudProviderData()
	GetWizCloudProviderListMocked(t, cloudProvidersIn, AppID, LocID)
	GetWizCloudProviderListFailErrMocked(t, cloudProvidersIn, AppID, LocID)
	GetWizCloudProviderListFailStatusMocked(t, cloudProvidersIn, AppID, LocID)
	GetWizCloudProviderListFailJSONMocked(t, cloudProvidersIn, AppID, LocID)
}
