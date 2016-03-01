package api

import (
	"github.com/flexiant/concerto/testdata"

	"testing"
)

func TestGetAdminReportList(t *testing.T) {
	adminReportsIn := testdata.GetAdminReportsData()
	GetAdminReportListMocked(t, adminReportsIn)
}
