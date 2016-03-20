package testdata

import (
	"github.com/flexiant/concerto/api/types"
	"time"
)

// GetLicenseeReportsData loads test data
func GetLicenseeReportsData() *[]types.LicenseeReport {

	start, _ := time.Parse("2016-01-01T00:00:00.000Z", "2016-01-01T00:00:00.000Z")
	end, _ := time.Parse("2016-02-01T00:00:00.000Z", "2016-02-01T00:00:00.000Z")

	testLicenseeReports := []types.LicenseeReport{
		{
			Id:            "fakeID00",
			Year:          2016,
			Month:         1,
			StartTime:     start,
			EndTime:       end,
			ServerSeconds: 9805083.213000001,
			Closed:        true,
			Lines:         []types.Lines{},
			// Lines: []types.Lines{
			// 	{
			// 		ID:               "fakeID0",
			// 		CommissionedAt:   start,
			// 		DecommissionedAt: end,
			// 		InstanceID:       "fakeInstanceID0",
			// 		InstanceName:     "fakeInstanceName0",
			// 		InstanceFQDN:     "fakeInstanceFQDN0",
			// 		Consumption:      1024,
			// 	},
			// 	{
			// 		ID:               "fakeID1",
			// 		CommissionedAt:   start,
			// 		DecommissionedAt: end,
			// 		InstanceID:       "fakeInstanceID1",
			// 		InstanceName:     "fakeInstanceName1",
			// 		InstanceFQDN:     "fakeInstanceFQDN1",
			// 		Consumption:      2048,
			// 	},
			// },
		},
		{
			Id:            "fakeID01",
			Year:          2016,
			Month:         1,
			StartTime:     start,
			EndTime:       end,
			ServerSeconds: 42866259.27299995,
			Closed:        false,
			Lines:         []types.Lines{},
		},
	}

	return &testLicenseeReports
}
