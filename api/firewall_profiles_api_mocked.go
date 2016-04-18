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

// GetFirewallProfileListFailErrMocked test mocked function
func GetFirewallProfileListFailErrMocked(t *testing.T, firewallProfilesIn *[]types.FirewallProfile) *[]types.FirewallProfile {

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
	cs.On("Get", "/v1/network/firewall_profiles").Return(dIn, 200, fmt.Errorf("Mocked error"))
	firewallProfilesOut, err := ds.GetFirewallProfileList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(firewallProfilesOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &firewallProfilesOut
}

// GetFirewallProfileListFailStatusMocked test mocked function
func GetFirewallProfileListFailStatusMocked(t *testing.T, firewallProfilesIn *[]types.FirewallProfile) *[]types.FirewallProfile {

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
	cs.On("Get", "/v1/network/firewall_profiles").Return(dIn, 499, nil)
	firewallProfilesOut, err := ds.GetFirewallProfileList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(firewallProfilesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &firewallProfilesOut
}

// GetFirewallProfileListFailJSONMocked test mocked function
func GetFirewallProfileListFailJSONMocked(t *testing.T, firewallProfilesIn *[]types.FirewallProfile) *[]types.FirewallProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallProfileService(cs)
	assert.Nil(err, "Couldn't load firewallProfile service")
	assert.NotNil(ds, "FirewallProfile service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/network/firewall_profiles").Return(dIn, 200, nil)
	firewallProfilesOut, err := ds.GetFirewallProfileList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(firewallProfilesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// GetFirewallProfileFailErrMocked test mocked function
func GetFirewallProfileFailErrMocked(t *testing.T, firewallProfile *types.FirewallProfile) *types.FirewallProfile {

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
	cs.On("Get", fmt.Sprintf("/v1/network/firewall_profiles/%s", firewallProfile.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	firewallProfileOut, err := ds.GetFirewallProfile(firewallProfile.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(firewallProfileOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return firewallProfileOut
}

// GetFirewallProfileFailStatusMocked test mocked function
func GetFirewallProfileFailStatusMocked(t *testing.T, firewallProfile *types.FirewallProfile) *types.FirewallProfile {

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
	cs.On("Get", fmt.Sprintf("/v1/network/firewall_profiles/%s", firewallProfile.Id)).Return(dIn, 499, nil)
	firewallProfileOut, err := ds.GetFirewallProfile(firewallProfile.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(firewallProfileOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return firewallProfileOut
}

// GetFirewallProfileFailJSONMocked test mocked function
func GetFirewallProfileFailJSONMocked(t *testing.T, firewallProfile *types.FirewallProfile) *types.FirewallProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallProfileService(cs)
	assert.Nil(err, "Couldn't load firewallProfile service")
	assert.NotNil(ds, "FirewallProfile service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/network/firewall_profiles/%s", firewallProfile.Id)).Return(dIn, 200, nil)
	firewallProfileOut, err := ds.GetFirewallProfile(firewallProfile.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(firewallProfileOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// CreateFirewallProfileFailErrMocked test mocked function
func CreateFirewallProfileFailErrMocked(t *testing.T, firewallProfileIn *types.FirewallProfile) *types.FirewallProfile {

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
	cs.On("Post", "/v1/network/firewall_profiles/", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	firewallProfileOut, err := ds.CreateFirewallProfile(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(firewallProfileOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return firewallProfileOut
}

// CreateFirewallProfileFailStatusMocked test mocked function
func CreateFirewallProfileFailStatusMocked(t *testing.T, firewallProfileIn *types.FirewallProfile) *types.FirewallProfile {

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
	cs.On("Post", "/v1/network/firewall_profiles/", mapIn).Return(dOut, 499, nil)
	firewallProfileOut, err := ds.CreateFirewallProfile(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(firewallProfileOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return firewallProfileOut
}

// CreateFirewallProfileFailJSONMocked test mocked function
func CreateFirewallProfileFailJSONMocked(t *testing.T, firewallProfileIn *types.FirewallProfile) *types.FirewallProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallProfileService(cs)
	assert.Nil(err, "Couldn't load firewallProfile service")
	assert.NotNil(ds, "FirewallProfile service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*firewallProfileIn)
	assert.Nil(err, "FirewallProfile test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/network/firewall_profiles/", mapIn).Return(dIn, 200, nil)
	firewallProfileOut, err := ds.CreateFirewallProfile(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(firewallProfileOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// UpdateFirewallProfileFailErrMocked test mocked function
func UpdateFirewallProfileFailErrMocked(t *testing.T, firewallProfileIn *types.FirewallProfile) *types.FirewallProfile {

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
	cs.On("Put", fmt.Sprintf("/v1/network/firewall_profiles/%s", firewallProfileIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	firewallProfileOut, err := ds.UpdateFirewallProfile(mapIn, firewallProfileIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(firewallProfileOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return firewallProfileOut
}

// UpdateFirewallProfileFailStatusMocked test mocked function
func UpdateFirewallProfileFailStatusMocked(t *testing.T, firewallProfileIn *types.FirewallProfile) *types.FirewallProfile {

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
	cs.On("Put", fmt.Sprintf("/v1/network/firewall_profiles/%s", firewallProfileIn.Id), mapIn).Return(dOut, 499, nil)
	firewallProfileOut, err := ds.UpdateFirewallProfile(mapIn, firewallProfileIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(firewallProfileOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return firewallProfileOut
}

// UpdateFirewallProfileFailJSONMocked test mocked function
func UpdateFirewallProfileFailJSONMocked(t *testing.T, firewallProfileIn *types.FirewallProfile) *types.FirewallProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewFirewallProfileService(cs)
	assert.Nil(err, "Couldn't load firewallProfile service")
	assert.NotNil(ds, "FirewallProfile service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*firewallProfileIn)
	assert.Nil(err, "FirewallProfile test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/network/firewall_profiles/%s", firewallProfileIn.Id), mapIn).Return(dIn, 200, nil)
	firewallProfileOut, err := ds.UpdateFirewallProfile(mapIn, firewallProfileIn.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(firewallProfileOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// DeleteFirewallProfileFailErrMocked test mocked function
func DeleteFirewallProfileFailErrMocked(t *testing.T, firewallProfileIn *types.FirewallProfile) {

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
	cs.On("Delete", fmt.Sprintf("/v1/network/firewall_profiles/%s", firewallProfileIn.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	err = ds.DeleteFirewallProfile(firewallProfileIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// DeleteFirewallProfileFailStatusMocked test mocked function
func DeleteFirewallProfileFailStatusMocked(t *testing.T, firewallProfileIn *types.FirewallProfile) {

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
	cs.On("Delete", fmt.Sprintf("/v1/network/firewall_profiles/%s", firewallProfileIn.Id)).Return(dIn, 499, nil)
	err = ds.DeleteFirewallProfile(firewallProfileIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
