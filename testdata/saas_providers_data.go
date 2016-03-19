package testdata

import "github.com/flexiant/concerto/api/types"

// GetSaasProviderData loads test data
func GetSaasProviderData() *[]types.SaasProvider {

	testSaasProviders := []types.SaasProvider{
		{
			Id:   "fakeID0",
			Name: "fakeName0",
			Required_account_data: []string{"accData0"},
		},
		{
			Id:   "fakeID1",
			Name: "fakeName1",
			Required_account_data: []string{"accData1", "accData2", "accData3"},
		},
	}

	return &testSaasProviders
}
