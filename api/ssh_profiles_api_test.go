package api

import (
	"github.com/flexiant/concerto/testdata"

	"testing"
)

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
