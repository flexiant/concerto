package testdata

import "github.com/flexiant/concerto/api/types"

// GetCloudAccountData loads test data
func GetCloudAccountData() *[]types.CloudAccount {

	testCloudAccounts := []types.CloudAccount{
		{
			Id:          "fakeID0",
			CloudProvId: "fakeProvID0",
		},
		{
			Id:          "fakeID1",
			CloudProvId: "fakeProvID1",
		},
	}

	return &testCloudAccounts
}
