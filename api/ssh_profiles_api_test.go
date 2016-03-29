package api

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSSHProfileServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewSSHProfileService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetSSHProfileList(t *testing.T) {
	sshProfilesIn := testdata.GetSSHProfileData()
	GetSSHProfileListMocked(t, sshProfilesIn)
}

func TestGetSSHProfile(t *testing.T) {
	sshProfilesIn := testdata.GetSSHProfileData()
	for _, sshProfileIn := range *sshProfilesIn {
		GetSSHProfileMocked(t, &sshProfileIn)
	}
}

func TestCreateSSHProfile(t *testing.T) {
	sshProfilesIn := testdata.GetSSHProfileData()
	for _, sshProfileIn := range *sshProfilesIn {
		CreateSSHProfileMocked(t, &sshProfileIn)
	}
}

func TestUpdateSSHProfile(t *testing.T) {
	sshProfilesIn := testdata.GetSSHProfileData()
	for _, sshProfileIn := range *sshProfilesIn {
		UpdateSSHProfileMocked(t, &sshProfileIn)
	}
}

func TestDeleteSSHProfile(t *testing.T) {
	sshProfilesIn := testdata.GetSSHProfileData()
	for _, sshProfileIn := range *sshProfilesIn {
		DeleteSSHProfileMocked(t, &sshProfileIn)
	}
}
