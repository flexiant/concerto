package api

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	// "time"
)

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

// GetAdminReportList returns the list of reports as an array of Report
func (rs *ReportService) GetAdminReportList() (reports []types.Report, err error) {
	log.Debug("GetReportList")

	data, status, err := rs.concertoService.Get("/v1/admin/reports")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &reports); err != nil {
		return nil, err
	}

	return reports, nil
}

// GetAdminReport returns a report by its ID
func (rs *ReportService) GetAdminReport(ID string) (report *types.Report, err error) {
	log.Debug("GetReport")

	data, status, err := rs.concertoService.Get(fmt.Sprintf("/v1/admin/reports/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &report); err != nil {
		return nil, err
	}

	return report, nil
}
