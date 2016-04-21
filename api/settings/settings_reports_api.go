package settings

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	// "time"
)

// SettingsReportService manages report operations
type SettingsReportService struct {
	concertoService utils.ConcertoService
}

// NewSettingsReportService returns a Concerto report service
func NewSettingsReportService(concertoService utils.ConcertoService) (*SettingsReportService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &SettingsReportService{
		concertoService: concertoService,
	}, nil
}

// GetSettingsReportList returns the list of reports as an array of Report
func (rs *SettingsReportService) GetSettingsReportList() (reports []types.SettingsReport, err error) {
	log.Debug("GetReportList")

	data, status, err := rs.concertoService.Get("/v1/settings/reports")
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

// GetSettingsReport returns a report by its ID
func (rs *SettingsReportService) GetSettingsReport(ID string) (report *types.SettingsReport, err error) {
	log.Debug("GetReport")

	data, status, err := rs.concertoService.Get(fmt.Sprintf("/v1/settings/reports/%s", ID))
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
