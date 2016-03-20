package api

import (
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO exclude from release compile

// GetFirewallProfileListMocked test mocked function
func GetFirewallProfileListMocked(t *testing.T, firewallProfilesIn *[]types.FirewallProfile) *[]types.FirewallProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallProfileService(cs)
	assert.Nil(err, "Couldn't load firewallProfile service")
	assert.NotNil(ds, "FirewallProfile service not instanced")

	// to json
	dIn, err := json.Marshal(firewallProfilesIn)
	assert.Nil(err, "FirewallProfile test data corrupted")

	// call service
	cs.On("Get", "/v1/network/firewall_profiles").Return(dIn, 200, nil)
	firewallProfilesOut, err := ds.GetFirewallProfileList()
	assert.Nil(err, "Error getting firewallProfile list")
	assert.Equal(*firewallProfilesIn, firewallProfilesOut, "GetFirewallProfileList returned different firewallProfiles")

	return &firewallProfilesOut
}

// GetFirewallProfileMocked test mocked function
func GetFirewallProfileMocked(t *testing.T, firewallProfile *types.FirewallProfile) *types.FirewallProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallProfileService(cs)
	assert.Nil(err, "Couldn't load firewallProfile service")
	assert.NotNil(ds, "FirewallProfile service not instanced")

	// to json
	dIn, err := json.Marshal(firewallProfile)
	assert.Nil(err, "FirewallProfile test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/network/firewall_profiles/%s", firewallProfile.Id)).Return(dIn, 200, nil)
	firewallProfileOut, err := ds.GetFirewallProfile(firewallProfile.Id)
	assert.Nil(err, "Error getting firewallProfile")
	assert.Equal(*firewallProfile, *firewallProfileOut, "GetFirewallProfile returned different firewallProfiles")

	return firewallProfileOut
}

// CreateFirewallProfileMocked test mocked function
func CreateFirewallProfileMocked(t *testing.T, firewallProfileIn *types.FirewallProfile) *types.FirewallProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallProfileService(cs)
	assert.Nil(err, "Couldn't load firewallProfile service")
	assert.NotNil(ds, "FirewallProfile service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*firewallProfileIn)
	assert.Nil(err, "FirewallProfile test data corrupted")

	// to json
	dOut, err := json.Marshal(firewallProfileIn)
	assert.Nil(err, "FirewallProfile test data corrupted")

	// call service
	cs.On("Post", "/v1/network/firewall_profiles/", mapIn).Return(dOut, 200, nil)
	firewallProfileOut, err := ds.CreateFirewallProfile(mapIn)
	assert.Nil(err, "Error creating firewallProfile list")
	assert.Equal(firewallProfileIn, firewallProfileOut, "CreateFirewallProfile returned different firewallProfiles")

	return firewallProfileOut
}

// UpdateFirewallProfileMocked test mocked function
func UpdateFirewallProfileMocked(t *testing.T, firewallProfileIn *types.FirewallProfile) *types.FirewallProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallProfileService(cs)
	assert.Nil(err, "Couldn't load firewallProfile service")
	assert.NotNil(ds, "FirewallProfile service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*firewallProfileIn)
	assert.Nil(err, "FirewallProfile test data corrupted")

	// to json
	dOut, err := json.Marshal(firewallProfileIn)
	assert.Nil(err, "FirewallProfile test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/network/firewall_profiles/%s", firewallProfileIn.Id), mapIn).Return(dOut, 200, nil)
	firewallProfileOut, err := ds.UpdateFirewallProfile(mapIn, firewallProfileIn.Id)
	assert.Nil(err, "Error updating firewallProfile list")
	assert.Equal(firewallProfileIn, firewallProfileOut, "UpdateFirewallProfile returned different firewallProfiles")

	return firewallProfileOut
}

// DeleteFirewallProfileMocked test mocked function
func DeleteFirewallProfileMocked(t *testing.T, firewallProfileIn *types.FirewallProfile) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallProfileService(cs)
	assert.Nil(err, "Couldn't load firewallProfile service")
	assert.NotNil(ds, "FirewallProfile service not instanced")

	// to json
	dIn, err := json.Marshal(firewallProfileIn)
	assert.Nil(err, "FirewallProfile test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/network/firewall_profiles/%s", firewallProfileIn.Id)).Return(dIn, 200, nil)
	err = ds.DeleteFirewallProfile(firewallProfileIn.Id)
	assert.Nil(err, "Error deleting firewallProfile")

}
