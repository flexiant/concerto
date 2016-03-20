package api

import (
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetLicenseeReportListMocked test mocked function
func GetLicenseeReportListMocked(t *testing.T, licenseeReportsIn *[]types.LicenseeReport) *[]types.LicenseeReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	rs, err := NewLicenseeReportService(cs)
	assert.Nil(err, "Couldn't load report service")
	assert.NotNil(rs, "Report service not instanced")

	// to json
	lrIn, err := json.Marshal(licenseeReportsIn)
	assert.Nil(err, "Licensee report test data corrupted")

	// call service
	cs.On("Get", "/v1/licensee/reports").Return(lrIn, 200, nil)
	lrOut, err := rs.GetLicenseeReportList()
	assert.Nil(err, "Error getting licensee report")

	// exclude time data from comparison
	assert.Equal(len(*licenseeReportsIn), len(lrOut), "Returned wrong number of elements")

	for i := 0; i < len(*licenseeReportsIn); i++ {
		(*licenseeReportsIn)[i].StartTime = lrOut[i].StartTime
		(*licenseeReportsIn)[i].EndTime = lrOut[i].EndTime
	}

	assert.Equal(*licenseeReportsIn, lrOut, "GetLicenseeReportList returned different licensee reports")

	return &lrOut
}

// GetLicenseeReportMocked test mocked function
func GetLicenseeReportMocked(t *testing.T, licenseeReport *types.LicenseeReport) *types.LicenseeReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lrs, err := NewLicenseeReportService(cs)
	assert.Nil(err, "Couldn't load licenseeReport service")
	assert.NotNil(lrs, "LicenseeReport service not instanced")

	// to json
	lrIn, err := json.Marshal(licenseeReport)
	assert.Nil(err, "LicenseeReport test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/licensee/reports/%s", licenseeReport.Id)).Return(lrIn, 200, nil)
	lrOut, err := lrs.GetLicenseeReport(licenseeReport.Id)
	assert.Nil(err, "Error getting licenseeReport")

	// exclude time data from comparison
	licenseeReport.StartTime = lrOut.StartTime
	licenseeReport.EndTime = lrOut.EndTime

	assert.Equal(*licenseeReport, *lrOut, "GetLicenseeReport returned different licenseeReports")

	return lrOut
}
