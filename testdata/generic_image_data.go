package testdata

import (
	"github.com/flexiant/concerto/api/types"
)

// GetGenericImageData loads test data
func GetGenericImageData() *[]types.GenericImage {

	testGenericImages := []types.GenericImage{
		{
			Id:   "fakeID0",
			Name: "fakeName0",
		},
		{
			Id:   "fakeID1",
			Name: "fakeName1",
		},
	}

	return &testGenericImages
}
