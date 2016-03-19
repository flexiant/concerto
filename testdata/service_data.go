package testdata

import (
	"github.com/flexiant/concerto/api/types"
)

// GetServiceData loads test data
func GetServiceData() *[]types.Service {

	testServices := []types.Service{
		{
			Id:          "fakeID0",
			Name:        "fakeName0",
			Description: "fakeDescription0",
			Public:      true,
			License:     "fakeLicense0",
			Recipes:     []string{"fakeRecipe01", "fakeRecipe02"},
		},
		{
			Id:          "fakeID1",
			Name:        "fakeName1",
			Description: "fakeDescription1",
			Public:      true,
			License:     "fakeLicense1",
			Recipes:     []string{"fakeRecipe11", "fakeRecipe12"},
		},
	}

	return &testServices
}
