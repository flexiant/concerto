package testdata

import (
	"github.com/flexiant/concerto/api/types"
)

// GetServerData loads test data
func GetServerData() *[]types.Server {

	testServers := []types.Server{
		{
			Id:             "fakeID0",
			Name:           "fakeName0",
			Fqdn:           "fakeFqdn0",
			State:          "fakeState0",
			Public_ip:      "fakePublicIP0",
			Workspace_id:   "fakeWorkspaceID0",
			Template_id:    "fakeTemplateID0",
			Server_plan_id: "fakeServerPlanID0",
			Ssh_profile_id: "fakeSSHProfileID0",
		},
		{
			Id:             "fakeID1",
			Name:           "fakeName1",
			Fqdn:           "fakeFqdn1",
			State:          "fakeState1",
			Public_ip:      "fakePublicIP1",
			Workspace_id:   "fakeWorkspaceID1",
			Template_id:    "fakeTemplateID1",
			Server_plan_id: "fakeServerPlanID1",
			Ssh_profile_id: "fakeSSHProfileID1",
		},
	}

	return &testServers
}

// GetDNSData loads test data
func GetDNSData() *[]types.Dns {

	testDnss := []types.Dns{
		{
			Id:        "fakeID0",
			Name:      "fakeName0",
			Content:   "fakeContent0",
			Type:      "fakeType0",
			IsFQDN:    true,
			Domain_id: "fakeDomainID0",
		},
		{
			Id:        "fakeID1",
			Name:      "fakeName1",
			Content:   "fakeContent1",
			Type:      "fakeType1",
			IsFQDN:    false,
			Domain_id: "fakeDomainID1",
		},
	}

	return &testDnss
}

// GetScriptCharData loads test data
func GetScriptCharData() *[]types.ScriptChar {

	testScriptChars := []types.ScriptChar{
		{
			Id:   "fakeID0",
			Type: "fakeType0",
			// Parameter_values: struct{"fakeInst0", "fakeInst1"},
			Template_id: "fakeTemplateID0",
			Script_id:   "fakeScriptID0",
		},
		{
			Id:   "fakeID1",
			Type: "fakeType1",
			// Parameter_values: struct{"fakeInst2", "fakeInst2", "fakeInst3"},
			Template_id: "fakeTemplateID1",
			Script_id:   "fakeScriptID1",
		},
	}

	return &testScriptChars
}
