package api

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFirewallProfileServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewFirewallProfileService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetFirewallProfileList(t *testing.T) {
	firewallProfilesIn := testdata.GetFirewallProfileData()
	GetFirewallProfileListMocked(t, firewallProfilesIn)
}

func TestGetFirewallProfile(t *testing.T) {
	firewallProfilesIn := testdata.GetFirewallProfileData()
	for _, firewallProfileIn := range *firewallProfilesIn {
		GetFirewallProfileMocked(t, &firewallProfileIn)
	}
}

func TestCreateFirewallProfile(t *testing.T) {
	firewallProfilesIn := testdata.GetFirewallProfileData()
	for _, firewallProfileIn := range *firewallProfilesIn {
		CreateFirewallProfileMocked(t, &firewallProfileIn)
	}
}

func TestUpdateFirewallProfile(t *testing.T) {
	firewallProfilesIn := testdata.GetFirewallProfileData()
	for _, firewallProfileIn := range *firewallProfilesIn {
		UpdateFirewallProfileMocked(t, &firewallProfileIn)
	}
}

func TestDeleteFirewallProfile(t *testing.T) {
	firewallProfilesIn := testdata.GetFirewallProfileData()
	for _, firewallProfileIn := range *firewallProfilesIn {
		DeleteFirewallProfileMocked(t, &firewallProfileIn)
	}
}
