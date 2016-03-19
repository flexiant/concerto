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

// GetSSHProfileListMocked test mocked function
func GetSSHProfileListMocked(t *testing.T, sshProfilesIn *[]types.SSHProfile) *[]types.SSHProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSSHProfileService(cs)
	assert.Nil(err, "Couldn't load sshProfile service")
	assert.NotNil(ds, "SSHProfile service not instanced")

	// to json
	dIn, err := json.Marshal(sshProfilesIn)
	assert.Nil(err, "SSHProfile test data corrupted")

	// call service
	cs.On("Get", "/v1/cloud/ssh_profiles").Return(dIn, 200, nil)
	sshProfilesOut, err := ds.GetSSHProfileList()
	assert.Nil(err, "Error getting sshProfile list")
	assert.Equal(*sshProfilesIn, sshProfilesOut, "GetSSHProfileList returned different sshProfiles")

	return &sshProfilesOut
}

// GetSSHProfileMocked test mocked function
func GetSSHProfileMocked(t *testing.T, sshProfile *types.SSHProfile) *types.SSHProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSSHProfileService(cs)
	assert.Nil(err, "Couldn't load sshProfile service")
	assert.NotNil(ds, "SSHProfile service not instanced")

	// to json
	dIn, err := json.Marshal(sshProfile)
	assert.Nil(err, "SSHProfile test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/ssh_profiles/%s", sshProfile.Id)).Return(dIn, 200, nil)
	sshProfileOut, err := ds.GetSSHProfile(sshProfile.Id)
	assert.Nil(err, "Error getting sshProfile")
	assert.Equal(*sshProfile, *sshProfileOut, "GetSSHProfile returned different sshProfiles")

	return sshProfileOut
}

// CreateSSHProfileMocked test mocked function
func CreateSSHProfileMocked(t *testing.T, sshProfileIn *types.SSHProfile) *types.SSHProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSSHProfileService(cs)
	assert.Nil(err, "Couldn't load sshProfile service")
	assert.NotNil(ds, "SSHProfile service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*sshProfileIn)
	assert.Nil(err, "SSHProfile test data corrupted")

	// to json
	dOut, err := json.Marshal(sshProfileIn)
	assert.Nil(err, "SSHProfile test data corrupted")

	// call service
	cs.On("Post", "/v1/cloud/ssh_profiles/", mapIn).Return(dOut, 200, nil)
	sshProfileOut, err := ds.CreateSSHProfile(mapIn)
	assert.Nil(err, "Error creating sshProfile list")
	assert.Equal(sshProfileIn, sshProfileOut, "CreateSSHProfile returned different sshProfiles")

	return sshProfileOut
}

// UpdateSSHProfileMocked test mocked function
func UpdateSSHProfileMocked(t *testing.T, sshProfileIn *types.SSHProfile) *types.SSHProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSSHProfileService(cs)
	assert.Nil(err, "Couldn't load sshProfile service")
	assert.NotNil(ds, "SSHProfile service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*sshProfileIn)
	assert.Nil(err, "SSHProfile test data corrupted")

	// to json
	dOut, err := json.Marshal(sshProfileIn)
	assert.Nil(err, "SSHProfile test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/ssh_profiles/%s", sshProfileIn.Id), mapIn).Return(dOut, 200, nil)
	sshProfileOut, err := ds.UpdateSSHProfile(mapIn, sshProfileIn.Id)
	assert.Nil(err, "Error updating sshProfile list")
	assert.Equal(sshProfileIn, sshProfileOut, "UpdateSSHProfile returned different sshProfiles")

	return sshProfileOut
}

// DeleteSSHProfileMocked test mocked function
func DeleteSSHProfileMocked(t *testing.T, sshProfileIn *types.SSHProfile) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSSHProfileService(cs)
	assert.Nil(err, "Couldn't load sshProfile service")
	assert.NotNil(ds, "SSHProfile service not instanced")

	// to json
	dIn, err := json.Marshal(sshProfileIn)
	assert.Nil(err, "SSHProfile test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/cloud/ssh_profiles/%s", sshProfileIn.Id)).Return(dIn, 200, nil)
	err = ds.DeleteSSHProfile(sshProfileIn.Id)
	assert.Nil(err, "Error deleting sshProfile")

}
