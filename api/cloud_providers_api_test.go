package api

import (
	"github.com/flexiant/concerto/testdata"

	"testing"
)

func TestGetCloudProviderList(t *testing.T) {
	cloudProvidersIn := testdata.GetCloudProviderData()
	GetCloudProviderListMocked(t, cloudProvidersIn)
}
