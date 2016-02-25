package testdata

import (
	"github.com/flexiant/concerto/api/types"
)

// GetEventData loads loads test data
func GetNodeData() *[]types.Node {

	testNodes := []types.Node{
		{
			Id:        "fakeID0",
			Name:      "fakeName0",
			Fqdn:      "fakeLevel0",
			PublicIp:  "fakePublicIP0",
			State:     "fakeState0",
			Os:        "fakeOS0",
			Plan:      "fakePlan0",
			FleetName: "fakeFleetName0",
			Master:    true,
		},
		{
			Id:        "fakeID1",
			Name:      "fakeName1",
			Fqdn:      "fakeLevel1",
			PublicIp:  "fakePublicIP1",
			State:     "fakeState1",
			Os:        "fakeOS1",
			Plan:      "fakePlan1",
			FleetName: "fakeFleetName1",
			Master:    false,
		},
	}

	return &testNodes
}
