package testdata

import (
	"github.com/flexiant/concerto/api/types"
)

// GetSSHProfileData loads test data
func GetSSHProfileData() *[]types.SSHProfile {

	testSSHProfiles := []types.SSHProfile{
		{
			Id:          "fakeID0",
			Name:        "fakeName0",
			Public_key:  "fakePubKey0",
			Private_key: "fakePrivKey0",
		},
		{
			Id:          "fakeID1",
			Name:        "fakeName1",
			Public_key:  "fakePubKey1",
			Private_key: "fakePrivKey1",
		},
	}

	return &testSSHProfiles
}
