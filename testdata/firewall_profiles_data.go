package testdata

import "github.com/flexiant/concerto/api/types"

// GetFirewallProfileData loads test data
func GetFirewallProfileData() *[]types.FirewallProfile {

	testFirewallProfiles := []types.FirewallProfile{
		{
			Id:          "fakeId0",
			Name:        "fakeName0",
			Description: "fakeDescription0",
			Default:     true,
			Rules: []types.Rule{
				{
					Protocol: "fakeProtocol0",
					MinPort:  0,
					MaxPort:  1024,
					CidrIp:   "fakeCidrIp0",
				},
			},
		},
		{
			Id:          "fakeId1",
			Name:        "fakeName1",
			Description: "fakeDescription1",
			Default:     false,
			Rules: []types.Rule{
				{
					Protocol: "fakeProtocol1",
					MinPort:  0,
					MaxPort:  200,
					CidrIp:   "fakeCidrIp1",
				},
				{
					Protocol: "fakeProtocol2",
					MinPort:  0,
					MaxPort:  2048,
					CidrIp:   "fakeCidrIp2",
				},
			},
		},
	}

	return &testFirewallProfiles
}
