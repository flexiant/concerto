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

// GetSSHProfileListFailErrMocked test mocked function
func GetSSHProfileListFailErrMocked(t *testing.T, sshProfilesIn *[]types.SSHProfile) *[]types.SSHProfile {

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
	cs.On("Get", "/v1/cloud/ssh_profiles").Return(dIn, 200, fmt.Errorf("Mocked error"))
	sshProfilesOut, err := ds.GetSSHProfileList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(sshProfilesOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &sshProfilesOut
}

// GetSSHProfileListFailStatusMocked test mocked function
func GetSSHProfileListFailStatusMocked(t *testing.T, sshProfilesIn *[]types.SSHProfile) *[]types.SSHProfile {

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
	cs.On("Get", "/v1/cloud/ssh_profiles").Return(dIn, 499, nil)
	sshProfilesOut, err := ds.GetSSHProfileList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(sshProfilesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &sshProfilesOut
}

// GetSSHProfileListFailJSONMocked test mocked function
func GetSSHProfileListFailJSONMocked(t *testing.T, sshProfilesIn *[]types.SSHProfile) *[]types.SSHProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSSHProfileService(cs)
	assert.Nil(err, "Couldn't load sshProfile service")
	assert.NotNil(ds, "SSHProfile service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/cloud/ssh_profiles").Return(dIn, 200, nil)
	sshProfilesOut, err := ds.GetSSHProfileList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(sshProfilesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// GetSSHProfileFailErrMocked test mocked function
func GetSSHProfileFailErrMocked(t *testing.T, sshProfile *types.SSHProfile) *types.SSHProfile {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/ssh_profiles/%s", sshProfile.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	sshProfileOut, err := ds.GetSSHProfile(sshProfile.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(sshProfileOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return sshProfileOut
}

// GetSSHProfileFailStatusMocked test mocked function
func GetSSHProfileFailStatusMocked(t *testing.T, sshProfile *types.SSHProfile) *types.SSHProfile {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/ssh_profiles/%s", sshProfile.Id)).Return(dIn, 499, nil)
	sshProfileOut, err := ds.GetSSHProfile(sshProfile.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(sshProfileOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return sshProfileOut
}

// GetSSHProfileFailJSONMocked test mocked function
func GetSSHProfileFailJSONMocked(t *testing.T, sshProfile *types.SSHProfile) *types.SSHProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSSHProfileService(cs)
	assert.Nil(err, "Couldn't load sshProfile service")
	assert.NotNil(ds, "SSHProfile service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/ssh_profiles/%s", sshProfile.Id)).Return(dIn, 200, nil)
	sshProfileOut, err := ds.GetSSHProfile(sshProfile.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(sshProfileOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// CreateSSHProfileFailErrMocked test mocked function
func CreateSSHProfileFailErrMocked(t *testing.T, sshProfileIn *types.SSHProfile) *types.SSHProfile {

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
	cs.On("Post", "/v1/cloud/ssh_profiles/", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	sshProfileOut, err := ds.CreateSSHProfile(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(sshProfileOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return sshProfileOut
}

// CreateSSHProfileFailStatusMocked test mocked function
func CreateSSHProfileFailStatusMocked(t *testing.T, sshProfileIn *types.SSHProfile) *types.SSHProfile {

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
	cs.On("Post", "/v1/cloud/ssh_profiles/", mapIn).Return(dOut, 499, nil)
	sshProfileOut, err := ds.CreateSSHProfile(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(sshProfileOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return sshProfileOut
}

// CreateSSHProfileFailJSONMocked test mocked function
func CreateSSHProfileFailJSONMocked(t *testing.T, sshProfileIn *types.SSHProfile) *types.SSHProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSSHProfileService(cs)
	assert.Nil(err, "Couldn't load sshProfile service")
	assert.NotNil(ds, "SSHProfile service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*sshProfileIn)
	assert.Nil(err, "SSHProfile test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/cloud/ssh_profiles/", mapIn).Return(dIn, 200, nil)
	sshProfileOut, err := ds.CreateSSHProfile(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(sshProfileOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// UpdateSSHProfileFailErrMocked test mocked function
func UpdateSSHProfileFailErrMocked(t *testing.T, sshProfileIn *types.SSHProfile) *types.SSHProfile {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/ssh_profiles/%s", sshProfileIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	sshProfileOut, err := ds.UpdateSSHProfile(mapIn, sshProfileIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(sshProfileOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return sshProfileOut
}

// UpdateSSHProfileFailStatusMocked test mocked function
func UpdateSSHProfileFailStatusMocked(t *testing.T, sshProfileIn *types.SSHProfile) *types.SSHProfile {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/ssh_profiles/%s", sshProfileIn.Id), mapIn).Return(dOut, 499, nil)
	sshProfileOut, err := ds.UpdateSSHProfile(mapIn, sshProfileIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(sshProfileOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return sshProfileOut
}

// UpdateSSHProfileFailJSONMocked test mocked function
func UpdateSSHProfileFailJSONMocked(t *testing.T, sshProfileIn *types.SSHProfile) *types.SSHProfile {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSSHProfileService(cs)
	assert.Nil(err, "Couldn't load sshProfile service")
	assert.NotNil(ds, "SSHProfile service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*sshProfileIn)
	assert.Nil(err, "SSHProfile test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/ssh_profiles/%s", sshProfileIn.Id), mapIn).Return(dIn, 200, nil)
	sshProfileOut, err := ds.UpdateSSHProfile(mapIn, sshProfileIn.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(sshProfileOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// DeleteSSHProfileFailErrMocked test mocked function
func DeleteSSHProfileFailErrMocked(t *testing.T, sshProfileIn *types.SSHProfile) {

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
	cs.On("Delete", fmt.Sprintf("/v1/cloud/ssh_profiles/%s", sshProfileIn.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	err = ds.DeleteSSHProfile(sshProfileIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// DeleteSSHProfileFailStatusMocked test mocked function
func DeleteSSHProfileFailStatusMocked(t *testing.T, sshProfileIn *types.SSHProfile) {

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
	cs.On("Delete", fmt.Sprintf("/v1/cloud/ssh_profiles/%s", sshProfileIn.Id)).Return(dIn, 499, nil)
	err = ds.DeleteSSHProfile(sshProfileIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
