package api

import (
	"github.com/flexiant/concerto/testdata"

	"testing"
)

func TestGetLicenseeReportList(t *testing.T) {
	licenseeReportsIn := testdata.GetLicenseeReportsData()
	GetLicenseeReportListMocked(t, licenseeReportsIn)
}

func TestGetLicenseeReport(t *testing.T) {
	licenseeReportsIn := testdata.GetLicenseeReportsData()
	for _, licenseeReportIn := range *licenseeReportsIn {
		GetLicenseeReportMocked(t, &licenseeReportIn)
	}
}
