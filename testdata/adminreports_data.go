package testdata

import (
	"github.com/flexiant/concerto/api/types"
	"time"
)

// GetAdminReportsData loads test data
func GetAdminReportsData() *[]types.Report {

	start, _ := time.Parse("2016-01-01T00:00:00.000Z", "2016-01-01T00:00:00.000Z")
	end, _ := time.Parse("2016-02-01T00:00:00.000Z", "2016-02-01T00:00:00.000Z")

	testAdminReports := []types.Report{
		{
			ID:            "5687130a4778ef000b000001",
			Year:          2016,
			Month:         1,
			StartTime:     start,
			EndTime:       end,
			ServerSeconds: 9805083.213000001,
			Closed:        true,
			AccountGroup: types.AccountGroup{
				ID:   "55b0911810c0ecc351000016",
				Name: "Flexiant Test1",
			},
		},
		{
			ID:            "5687130a4778ef000b000002",
			Year:          2016,
			Month:         1,
			StartTime:     start,
			EndTime:       end,
			ServerSeconds: 42866259.27299995,
			Closed:        true,
			AccountGroup: types.AccountGroup{
				ID:   "55b0911810c0ecc351000017",
				Name: "Flexiant Test2",
			},
		},
		{
			ID:            "5687130a4778ef000b000003",
			Year:          2016,
			Month:         1,
			StartTime:     start,
			EndTime:       end,
			ServerSeconds: 0.0,
			Closed:        true,
			AccountGroup: types.AccountGroup{
				ID:   "55b0911810c0ecc351000018",
				Name: "Flexiant Test3",
			},
		},
	}

	return &testAdminReports
}

// type Lines struct {
// 	Id                string    `json:"_id" header:"ID"`
// 	Commissioned_at   time.Time `json:"commissioned_at" header:"COMMISSIONED_AT"`
// 	Decommissioned_at time.Time `json:"decommissioned_at" header:"DECOMMISSIONED_AT"`
// 	Instance_id       string    `json:"instance_id" header:"INSTANCE_ID"`
// 	Instance_name     string    `json:"instance_name" header:"INSTANCE_NAME"`
// 	Instance_fqdn     string    `json:"instance_fqdn" header:"INSTANCE_FQDN"`
// 	Consumption       float32   `json:"consumption" header:"CONSUMPTION"`
// }
