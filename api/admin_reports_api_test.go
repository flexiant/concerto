package api

import (
	"testing"

	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewReportServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewReportService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetAdminReportList(t *testing.T) {
	adminReportsIn := testdata.GetAdminReportsData()
	GetAdminReportListMocked(t, adminReportsIn)
	GetAdminReportListErrMocked(t, adminReportsIn)
	GetAdminReportListFailStatusMocked(t, adminReportsIn)
	GetAdminReportListFailJSONMocked(t, adminReportsIn)
}

func TestGetAdminReport(t *testing.T) {
	adminReportsIn := testdata.GetAdminReportsData()
	for _, adminReportIn := range *adminReportsIn {
		GetAdminReportMocked(t, &adminReportIn)
		GetAdminReportErrMocked(t, &adminReportIn)
		GetAdminReportFailStatusMocked(t, &adminReportIn)
		GetAdminReportFailJSONMocked(t, &adminReportIn)
	}
}
