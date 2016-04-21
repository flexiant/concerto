package licensee

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

// GetLicenseeReportListFailErrMocked test mocked function
func GetLicenseeReportListFailErrMocked(t *testing.T, licenseeReportsIn *[]types.LicenseeReport) *[]types.LicenseeReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLicenseeReportService(cs)
	assert.Nil(err, "Couldn't load licenseeReport service")
	assert.NotNil(ds, "LicenseeReport service not instanced")

	// to json
	dIn, err := json.Marshal(licenseeReportsIn)
	assert.Nil(err, "LicenseeReport test data corrupted")

	// call service
	cs.On("Get", "/v1/licensee/reports").Return(dIn, 200, fmt.Errorf("Mocked error"))
	licenseeReportsOut, err := ds.GetLicenseeReportList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(licenseeReportsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &licenseeReportsOut
}

// GetLicenseeReportListFailStatusMocked test mocked function
func GetLicenseeReportListFailStatusMocked(t *testing.T, licenseeReportsIn *[]types.LicenseeReport) *[]types.LicenseeReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLicenseeReportService(cs)
	assert.Nil(err, "Couldn't load licenseeReport service")
	assert.NotNil(ds, "LicenseeReport service not instanced")

	// to json
	dIn, err := json.Marshal(licenseeReportsIn)
	assert.Nil(err, "LicenseeReport test data corrupted")

	// call service
	cs.On("Get", "/v1/licensee/reports").Return(dIn, 499, nil)
	licenseeReportsOut, err := ds.GetLicenseeReportList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(licenseeReportsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &licenseeReportsOut
}

// GetLicenseeReportListFailJSONMocked test mocked function
func GetLicenseeReportListFailJSONMocked(t *testing.T, licenseeReportsIn *[]types.LicenseeReport) *[]types.LicenseeReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLicenseeReportService(cs)
	assert.Nil(err, "Couldn't load licenseeReport service")
	assert.NotNil(ds, "LicenseeReport service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/licensee/reports").Return(dIn, 200, nil)
	licenseeReportsOut, err := ds.GetLicenseeReportList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(licenseeReportsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &licenseeReportsOut
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

// GetLicenseeReportFailErrMocked test mocked function
func GetLicenseeReportFailErrMocked(t *testing.T, licenseeReport *types.LicenseeReport) *types.LicenseeReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLicenseeReportService(cs)
	assert.Nil(err, "Couldn't load licenseeReport service")
	assert.NotNil(ds, "LicenseeReport service not instanced")

	// to json
	dIn, err := json.Marshal(licenseeReport)
	assert.Nil(err, "LicenseeReport test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/licensee/reports/%s", licenseeReport.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	licenseeReportOut, err := ds.GetLicenseeReport(licenseeReport.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(licenseeReportOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return licenseeReportOut
}

// GetLicenseeReportFailStatusMocked test mocked function
func GetLicenseeReportFailStatusMocked(t *testing.T, licenseeReport *types.LicenseeReport) *types.LicenseeReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLicenseeReportService(cs)
	assert.Nil(err, "Couldn't load licenseeReport service")
	assert.NotNil(ds, "LicenseeReport service not instanced")

	// to json
	dIn, err := json.Marshal(licenseeReport)
	assert.Nil(err, "LicenseeReport test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/licensee/reports/%s", licenseeReport.Id)).Return(dIn, 499, nil)
	licenseeReportOut, err := ds.GetLicenseeReport(licenseeReport.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(licenseeReportOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return licenseeReportOut
}

// GetLicenseeReportFailJSONMocked test mocked function
func GetLicenseeReportFailJSONMocked(t *testing.T, licenseeReport *types.LicenseeReport) *types.LicenseeReport {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewLicenseeReportService(cs)
	assert.Nil(err, "Couldn't load licenseeReport service")
	assert.NotNil(ds, "LicenseeReport service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/licensee/reports/%s", licenseeReport.Id)).Return(dIn, 200, nil)
	licenseeReportOut, err := ds.GetLicenseeReport(licenseeReport.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(licenseeReportOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return licenseeReportOut
}
