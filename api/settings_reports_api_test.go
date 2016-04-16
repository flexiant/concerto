package api

import (
	"testing"

	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewSettingsReportServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewSettingsReportService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetSettingsReportList(t *testing.T) {
	settingsReportsIn := testdata.GetSettingsReportsData()
	GetSettingsReportListMocked(t, settingsReportsIn)
	GetSettingsReportListErrMocked(t, settingsReportsIn)
	GetSettingsReportListFailStatusMocked(t, settingsReportsIn)
	GetSettingsReportListFailJSONMocked(t, settingsReportsIn)
}

func TestGetSettingsReport(t *testing.T) {
	settingsReportsIn := testdata.GetSettingsReportsData()
	for _, settingsReportIn := range *settingsReportsIn {
		GetSettingsReportMocked(t, &settingsReportIn)
		GetSettingsReportErrMocked(t, &settingsReportIn)
		GetSettingsReportFailStatusMocked(t, &settingsReportIn)
		GetSettingsReportFailJSONMocked(t, &settingsReportIn)
	}
}
