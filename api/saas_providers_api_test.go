package api

import (
	"github.com/flexiant/concerto/testdata"

	"testing"
)

func TestGetSaasProviderList(t *testing.T) {
	saasProvidersIn := testdata.GetSaasProviderData()
	GetSaasProviderListMocked(t, saasProvidersIn)
}
