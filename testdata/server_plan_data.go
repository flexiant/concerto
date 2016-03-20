package testdata

import "github.com/flexiant/concerto/api/types"

// GetServerPlanData loads test data
func GetServerPlanData() *[]types.ServerPlan {

	testServerPlans := []types.ServerPlan{
		{
			Id:              "fakeID0",
			Name:            "fakeName0",
			Memory:          512,
			CPUs:            2,
			Storage:         2048,
			LocationId:      "fakeLocationID0",
			CloudProviderId: "fakeCloudProviderID0",
		},
		{
			Id:              "fakeID1",
			Name:            "fakeName1",
			Memory:          256,
			CPUs:            3,
			Storage:         1024,
			LocationId:      "fakeLocationID1",
			CloudProviderId: "fakeCloudProviderID1",
		},
	}

	return &testServerPlans
}
