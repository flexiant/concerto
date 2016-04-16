package testdata

import (
	"github.com/flexiant/concerto/api/types"
	"time"
)

// GetSettingsReportsData loads test data
func GetSettingsReportsData() *[]types.SettingsReport {

	start, _ := time.Parse("2016-01-01T00:00:00.000Z", "2016-01-01T00:00:00.000Z")
	end, _ := time.Parse("2016-02-01T00:00:00.000Z", "2016-02-01T00:00:00.000Z")

	testSettingsReports := []types.SettingsReport{
		{
			ID:            "5687130a4778ef000b000001",
			Year:          2016,
			Month:         1,
			StartTime:     start,
			EndTime:       end,
			ServerSeconds: 9805083.213000001,
			Closed:        true,
		},
		{
			ID:            "5687130a4778ef000b000002",
			Year:          2016,
			Month:         1,
			StartTime:     start,
			EndTime:       end,
			ServerSeconds: 42866259.27299995,
			Closed:        true,
		},
		{
			ID:            "5687130a4778ef000b000003",
			Year:          2016,
			Month:         1,
			StartTime:     start,
			EndTime:       end,
			ServerSeconds: 0.0,
			Closed:        true,
		},
	}

	return &testSettingsReports
}
