package testdata

import (
	"github.com/flexiant/concerto/api/types"
)

// GetWorkspaceData loads test data
func GetWorkspaceData() *[]types.Workspace {

	testWorkspaces := []types.Workspace{
		{
			Id:                  "fakeID0",
			Name:                "fakeName0",
			Default:             true,
			Domain_id:           "fakeDomID0",
			Ssh_profile_id:      "fakeSSHProfileID0",
			Firewall_profile_id: "fakeFirewallProfileID0",
		},
		{
			Id:                  "fakeID1",
			Name:                "fakeName1",
			Default:             false,
			Domain_id:           "fakeDomID1",
			Ssh_profile_id:      "fakeSSHProfileID1",
			Firewall_profile_id: "fakeFirewallProfileID1",
		},
	}

	return &testWorkspaces
}

// GetWorkspaceServerData loads test data
func GetWorkspaceServerData() *[]types.WorkspaceServer {

	testWorkspaceServers := []types.WorkspaceServer{
		{
			Id:             "fakeID0",
			Name:           "fakeName0",
			Fqdn:           "fakeFQDN0",
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
			Fqdn:           "fakeFQDN1",
			State:          "fakeState1",
			Public_ip:      "fakePublicIP1",
			Workspace_id:   "fakeWorkspaceID1",
			Template_id:    "fakeTemplateID1",
			Server_plan_id: "fakeServerPlanID1",
			Ssh_profile_id: "fakeSSHProfileID1",
		},
	}

	return &testWorkspaceServers
}
