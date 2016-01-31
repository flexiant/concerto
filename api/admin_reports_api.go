package api

import (
	// "encoding/json"
	"fmt"
	// log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/utils"
	"time"
)

type Report struct {
	Id             string       `json:"id" header:"ID"`
	Year           int          `json:"year" header:"YEAR"`
	Month          time.Month   `json:"month" header:"MONTH"`
	Start_time     time.Time    `json:"start_time" header:"START_TIME"`
	End_time       time.Time    `json:"end_time" header:"END_TIME"`
	Server_seconds float32      `json:"server_seconds" header:"SERVER_SECONDS"`
	Closed         bool         `json:"closed" header:"CLOSED"`
	Li             []Lines      `json:"lines" header:"LINES"`
	Account_group  AccountGroup `json:"account_group" header:"ACCOUNT_GROUP"`
}

type Lines struct {
	Id                string    `json:"_id" header:"ID"`
	Commissioned_at   time.Time `json:"commissioned_at" header:"COMMISSIONED_AT"`
	Decommissioned_at time.Time `json:"decommissioned_at" header:"DECOMMISSIONED_AT"`
	Instance_id       string    `json:"instance_id" header:"INSTANCE_ID"`
	Instance_name     string    `json:"instance_name" header:"INSTANCE_NAME"`
	Instance_fqdn     string    `json:"instance_fqdn" header:"INSTANCE_FQDN"`
	Consumption       float32   `json:"consumption" header:"CONSUMPTION"`
}

type AccountGroup struct {
	Id   string `json:"_id" header:"ID"`
	Name string `json:"name" header:"NAME"`
}

// ReportService manages report operations
type ReportService struct {
	concertoService utils.ConcertoService
}

// NewReportService returns a Concerto report service
func NewReportService(concertoService utils.ConcertoService) (*ReportService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &ReportService{
		concertoService: concertoService,
	}, nil
}

// // GetAdminReportList returns the list of reports as an array of Report
// func (cl *ReportService) GetAdminReportList() (reports []Report, err error) {
// 	log.Debug("GetReportList")

// 	data, status, err := cl.concertoService.Get("/v1/admin/reports")
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = utils.CheckStandardStatus(status, data); err != nil {

// 	}

// 	if err = json.Unmarshal(data, &reports); err != nil {
// 		return nil, err
// 	}

// 	return reports, nil
// }

// // GetAdminReport returns the list of reports as an array of Report
// func (cl *ReportService) GetAdminReport(ID string) (report Report, err error) {
// 	log.Debug("GetReport")

// 	data, status, err := cl.concertoService.Get(fmt.Sprintf("/v1/admin/reports/%s", ID))
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err = utils.CheckStandardStatus(status, data); err != nil {

// 	}

// 	if err = json.Unmarshal(data, &report); err != nil {
// 		return nil, err
// 	}

// 	return report, nil
// }
