package testdata

import (
	"github.com/flexiant/concerto/api/types"
)

// GetSSHProfileData loads test data
func GetSSHProfileData() *[]types.SSHProfile {

	testSSHProfiles := []types.SSHProfile{
		{
			Id:         "fakeID0",
			Name:       "fakeName0",
			PublicKey:  "fakePubKey0",
			PrivateKey: "fakePrivKey0",
		},
		{
			Id:         "fakeID1",
			Name:       "fakeName1",
			PublicKey:  "fakePubKey1",
			PrivateKey: "fakePrivKey1",
		},
	}

	return &testSSHProfiles
}
