package settings

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// GetSettingsReportListMocked test mocked function
func GetSettingsReportListMocked(t *testing.T, settingsReportsIn *[]types.SettingsReport) *[]types.SettingsReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	rs, err := NewSettingsReportService(cs)
	assert.Nil(err, "Couldn't load report service")
	assert.NotNil(rs, "Report service not instanced")

	// to json
	arIn, err := json.Marshal(settingsReportsIn)
	assert.Nil(err, "Admin reports test data corrupted")

	// call service
	cs.On("Get", "/v1/settings/reports").Return(arIn, 200, nil)
	arOut, err := rs.GetSettingsReportList()
	assert.Nil(err, "Error getting settings report")

	// exclude time data from comparison
	assert.Equal(len(*settingsReportsIn), len(arOut), "Returned wrong number of elements")

	for i := 0; i < len(*settingsReportsIn); i++ {
		(*settingsReportsIn)[i].StartTime = arOut[i].StartTime
		(*settingsReportsIn)[i].EndTime = arOut[i].EndTime
	}

	assert.Equal(*settingsReportsIn, arOut, "GetSettingsReportList returned different settings reports")

	return &arOut
}

// GetSettingsReportListErrMocked test mocked function
func GetSettingsReportListErrMocked(t *testing.T, settingsReportsIn *[]types.SettingsReport) *[]types.SettingsReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	rs, err := NewSettingsReportService(cs)
	assert.Nil(err, "Couldn't load report service")
	assert.NotNil(rs, "Report service not instanced")

	// to json
	arIn, err := json.Marshal(settingsReportsIn)
	assert.Nil(err, "Admin reports test data corrupted")

	// call service
	cs.On("Get", "/v1/settings/reports").Return(arIn, 200, fmt.Errorf("Mocked error"))
	arOut, err := rs.GetSettingsReportList()
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(arOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &arOut
}

// GetSettingsReportListFailStatusMocked test mocked function
func GetSettingsReportListFailStatusMocked(t *testing.T, settingsReportsIn *[]types.SettingsReport) *[]types.SettingsReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	rs, err := NewSettingsReportService(cs)
	assert.Nil(err, "Couldn't load report service")
	assert.NotNil(rs, "Report service not instanced")

	// to json
	arIn, err := json.Marshal(settingsReportsIn)
	assert.Nil(err, "Admin reports test data corrupted")

	// call service
	cs.On("Get", "/v1/settings/reports").Return(arIn, 499, nil)
	arOut, err := rs.GetSettingsReportList()
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(arOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &arOut
}

// GetSettingsReportListFailJSONMocked test mocked function
func GetSettingsReportListFailJSONMocked(t *testing.T, settingsReportsIn *[]types.SettingsReport) *[]types.SettingsReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	rs, err := NewSettingsReportService(cs)
	assert.Nil(err, "Couldn't load report service")
	assert.NotNil(rs, "Report service not instanced")

	// wrong json
	arIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/settings/reports").Return(arIn, 200, nil)
	arOut, err := rs.GetSettingsReportList()
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &arOut
}

// GetSettingsReportMocked test mocked function
func GetSettingsReportMocked(t *testing.T, settingsReport *types.SettingsReport) *types.SettingsReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ars, err := NewSettingsReportService(cs)
	assert.Nil(err, "Couldn't load settingsReport service")
	assert.NotNil(ars, "SettingsReport service not instanced")

	// to json
	arIn, err := json.Marshal(settingsReport)
	assert.Nil(err, "SettingsReport test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/settings/reports/%s", settingsReport.ID)).Return(arIn, 200, nil)
	arOut, err := ars.GetSettingsReport(settingsReport.ID)
	assert.Nil(err, "Error getting settingsReport")

	// exclude time data from comparison
	settingsReport.StartTime = arOut.StartTime
	settingsReport.EndTime = arOut.EndTime

	assert.Equal(*settingsReport, *arOut, "GetSettingsReport returned different settingsReports")

	return arOut
}

// GetSettingsReportErrMocked test mocked function
func GetSettingsReportErrMocked(t *testing.T, settingsReport *types.SettingsReport) *types.SettingsReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ars, err := NewSettingsReportService(cs)
	assert.Nil(err, "Couldn't load settingsReport service")
	assert.NotNil(ars, "SettingsReport service not instanced")

	// to json
	arIn, err := json.Marshal(settingsReport)
	assert.Nil(err, "SettingsReport test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/settings/reports/%s", settingsReport.ID)).Return(arIn, 200, fmt.Errorf("Mocked error"))
	arOut, err := ars.GetSettingsReport(settingsReport.ID)
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(arOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return arOut
}

// GetSettingsReportFailStatusMocked test mocked function
func GetSettingsReportFailStatusMocked(t *testing.T, settingsReport *types.SettingsReport) *types.SettingsReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ars, err := NewSettingsReportService(cs)
	assert.Nil(err, "Couldn't load settingsReport service")
	assert.NotNil(ars, "SettingsReport service not instanced")

	// to json
	arIn, err := json.Marshal(settingsReport)
	assert.Nil(err, "SettingsReport test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/settings/reports/%s", settingsReport.ID)).Return(arIn, 499, nil)
	arOut, err := ars.GetSettingsReport(settingsReport.ID)
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(arOut, "Expecting nil output")

	return arOut
}

// GetSettingsReportFailJSONMocked test mocked function
func GetSettingsReportFailJSONMocked(t *testing.T, settingsReport *types.SettingsReport) *types.SettingsReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ars, err := NewSettingsReportService(cs)
	assert.Nil(err, "Couldn't load settingsReport service")
	assert.NotNil(ars, "SettingsReport service not instanced")

	// wrong json
	arIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/settings/reports/%s", settingsReport.ID)).Return(arIn, 200, nil)
	arOut, err := ars.GetSettingsReport(settingsReport.ID)
	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return arOut
}
