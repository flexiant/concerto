package testdata

import (
	"encoding/json"
	"github.com/flexiant/concerto/api/types"
)

// GetAppData loads test data
func GetAppData() *[]types.WizardApp {

	param0 := json.RawMessage(`{"fakeFlavour01":"x","fakeFlavour02":"y"}`)
	param1 := json.RawMessage(`{"fakeFlavour11":"a","fakeFlavour12":"b"}`)

	testApps := []types.WizardApp{
		{
			Id:                   "fakeID0",
			Name:                 "fakeName0",
			Flavour_requirements: param0,
			Generic_image_id:     "fakeGenericImageID0",
		},
		{
			Id:                   "fakeID1",
			Name:                 "fakeName1",
			Flavour_requirements: param1,
			Generic_image_id:     "fakeGenericImageID1",
		},
	}

	return &testApps
}
