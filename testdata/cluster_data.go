package testdata

import "github.com/flexiant/concerto/api/types"

// GetClusterData loads test data
func GetClusterData() *[]types.Cluster {

	testClusters := []types.Cluster{
		{
			Id:                "fakeID0",
			Name:              "fakeName0",
			State:             "state0",
			MasterCount:       1,
			SlaveCount:        10,
			WorkspaceId:       "fakeWrkID0",
			FirewallProfileId: "fakeFirewallID0",
			MasterTemplateId:  "fakeMasterTemplID0",
			SlaveTemplateId:   "fakeSlaveTemplID0",
			Masters:           []string{"master1"},
		},
		{
			Id:                "fakeID1",
			Name:              "fakeName1",
			State:             "state1",
			MasterCount:       2,
			SlaveCount:        3,
			WorkspaceId:       "fakeWrkID1",
			FirewallProfileId: "fakeFirewallID1",
			MasterTemplateId:  "fakeMasterTemplID1",
			SlaveTemplateId:   "fakeSlaveTemplID1",
			Masters:           []string{"master0", "master2"},
		},
	}

	return &testClusters
}
