package testdata

import (
	"github.com/flexiant/concerto/api/types"
)

// GetLocationData loads test data
func GetLocationData() *[]types.Location {

	testLocations := []types.Location{
		{
			Id:   "fakeID0",
			Name: "fakeName0",
		},
		{
			Id:   "fakeID1",
			Name: "fakeName1",
		},
	}

	return &testLocations
}
