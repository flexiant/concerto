package testdata

import "github.com/flexiant/concerto/api/types"

// GetSaasAccountData loads test data
func GetSaasAccountData() *[]types.SaasAccount {

	testSaasAccounts := []types.SaasAccount{
		{
			Id:         "fakeID0",
			SaasProvId: "fakeProvID0",
		},
		{
			Id:         "fakeID1",
			SaasProvId: "fakeProvID1",
		},
	}

	return &testSaasAccounts
}
