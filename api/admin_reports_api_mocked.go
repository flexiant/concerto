package api

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// GetAdminReportListMocked test mocked function
func GetAdminReportListMocked(t *testing.T, adminReportsIn *[]types.Report) *[]types.Report {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	rs, err := NewReportService(cs)
	assert.Nil(err, "Couldn't load report service")
	assert.NotNil(rs, "Report service not instanced")

	// to json
	arIn, err := json.Marshal(adminReportsIn)
	assert.Nil(err, "Admin reports test data corrupted")

	// call service
	cs.On("Get", "/v1/admin/reports").Return(arIn, 200, nil)
	arOut, err := rs.GetAdminReportList()
	assert.Nil(err, "Error getting admin report")

	// exclude time data from comparison
	assert.Equal(len(*adminReportsIn), len(arOut), "Returned wrong number of elements")

	for i := 0; i < len(*adminReportsIn); i++ {
		(*adminReportsIn)[i].StartTime = arOut[i].StartTime
		(*adminReportsIn)[i].EndTime = arOut[i].EndTime
	}

	assert.Equal(*adminReportsIn, arOut, "GetAdminReportList returned different admin reports")

	return &arOut
}

// GetAdminReportListErrMocked test mocked function
func GetAdminReportListErrMocked(t *testing.T, adminReportsIn *[]types.Report) *[]types.Report {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	rs, err := NewReportService(cs)
	assert.Nil(err, "Couldn't load report service")
	assert.NotNil(rs, "Report service not instanced")

	// to json
	arIn, err := json.Marshal(adminReportsIn)
	assert.Nil(err, "Admin reports test data corrupted")

	// call service
	cs.On("Get", "/v1/admin/reports").Return(arIn, 200, fmt.Errorf("Mocked error"))
	arOut, err := rs.GetAdminReportList()
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(arOut, "Expecting nil output")

	return &arOut
}

// GetAdminReportListFailStatusMocked test mocked function
func GetAdminReportListFailStatusMocked(t *testing.T, adminReportsIn *[]types.Report) *[]types.Report {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	rs, err := NewReportService(cs)
	assert.Nil(err, "Couldn't load report service")
	assert.NotNil(rs, "Report service not instanced")

	// to json
	arIn, err := json.Marshal(adminReportsIn)
	assert.Nil(err, "Admin reports test data corrupted")

	// call service
	cs.On("Get", "/v1/admin/reports").Return(arIn, 499, nil)
	arOut, err := rs.GetAdminReportList()
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(arOut, "Expecting nil output")

	return &arOut
}

// GetAdminReportListFailJSONMocked test mocked function
func GetAdminReportListFailJSONMocked(t *testing.T, adminReportsIn *[]types.Report) *[]types.Report {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	rs, err := NewReportService(cs)
	assert.Nil(err, "Couldn't load report service")
	assert.NotNil(rs, "Report service not instanced")

	// wrong json
	arIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/admin/reports").Return(arIn, 200, nil)
	arOut, err := rs.GetAdminReportList()
	assert.NotNil(err, "We are expecting a marshalling error")

	return &arOut
}

// GetAdminReportMocked test mocked function
func GetAdminReportMocked(t *testing.T, adminReport *types.Report) *types.Report {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ars, err := NewReportService(cs)
	assert.Nil(err, "Couldn't load adminReport service")
	assert.NotNil(ars, "AdminReport service not instanced")

	// to json
	arIn, err := json.Marshal(adminReport)
	assert.Nil(err, "AdminReport test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/admin/reports/%s", adminReport.ID)).Return(arIn, 200, nil)
	arOut, err := ars.GetAdminReport(adminReport.ID)
	assert.Nil(err, "Error getting adminReport")

	// exclude time data from comparison
	adminReport.StartTime = arOut.StartTime
	adminReport.EndTime = arOut.EndTime

	assert.Equal(*adminReport, *arOut, "GetAdminReport returned different adminReports")

	return arOut
}

// GetAdminReportErrMocked test mocked function
func GetAdminReportErrMocked(t *testing.T, adminReport *types.Report) *types.Report {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ars, err := NewReportService(cs)
	assert.Nil(err, "Couldn't load adminReport service")
	assert.NotNil(ars, "AdminReport service not instanced")

	// to json
	arIn, err := json.Marshal(adminReport)
	assert.Nil(err, "AdminReport test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/admin/reports/%s", adminReport.ID)).Return(arIn, 200, fmt.Errorf("Mocked error"))
	arOut, err := ars.GetAdminReport(adminReport.ID)
	assert.NotNil(err, "We are expecting an error")
	assert.Nil(arOut, "Expecting nil output")

	return arOut
}

// GetAdminReportFailStatusMocked test mocked function
func GetAdminReportFailStatusMocked(t *testing.T, adminReport *types.Report) *types.Report {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ars, err := NewReportService(cs)
	assert.Nil(err, "Couldn't load adminReport service")
	assert.NotNil(ars, "AdminReport service not instanced")

	// to json
	arIn, err := json.Marshal(adminReport)
	assert.Nil(err, "AdminReport test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/admin/reports/%s", adminReport.ID)).Return(arIn, 499, nil)
	arOut, err := ars.GetAdminReport(adminReport.ID)
	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(arOut, "Expecting nil output")

	return arOut
}

// GetAdminReportFailJSONMocked test mocked function
func GetAdminReportFailJSONMocked(t *testing.T, adminReport *types.Report) *types.Report {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ars, err := NewReportService(cs)
	assert.Nil(err, "Couldn't load adminReport service")
	assert.NotNil(ars, "AdminReport service not instanced")

	// wrong json
	arIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/admin/reports/%s", adminReport.ID)).Return(arIn, 499, nil)
	arOut, err := ars.GetAdminReport(adminReport.ID)
	assert.NotNil(err, "We are expecting a marshalling error")

	return arOut
}
