package api

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	// "time"
)

// LicenseeReportService manages report operations
type LicenseeReportService struct {
	concertoService utils.ConcertoService
}

// NewLicenseeReportService returns a Concerto report service
func NewLicenseeReportService(concertoService utils.ConcertoService) (*LicenseeReportService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &LicenseeReportService{
		concertoService: concertoService,
	}, nil
}

// GetLicenseeReportList returns the list of reports as an array of LicenseeReport
func (rs *LicenseeReportService) GetLicenseeReportList() (reports []types.LicenseeReport, err error) {
	log.Debug("GetReportList")

	data, status, err := rs.concertoService.Get("/v1/licensee/reports")
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

// GetLicenseeReport returns a licensee report by its ID
func (rs *LicenseeReportService) GetLicenseeReport(ID string) (report *types.LicenseeReport, err error) {
	log.Debug("GetReport")

	data, status, err := rs.concertoService.Get(fmt.Sprintf("/v1/licensee/reports/%s", ID))
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
