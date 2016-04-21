package licensee

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLicenseeReportServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewLicenseeReportService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetLicenseeReportList(t *testing.T) {
	licenseeReportsIn := testdata.GetLicenseeReportsData()
	GetLicenseeReportListMocked(t, licenseeReportsIn)
	GetLicenseeReportListFailErrMocked(t, licenseeReportsIn)
	GetLicenseeReportListFailStatusMocked(t, licenseeReportsIn)
	GetLicenseeReportListFailJSONMocked(t, licenseeReportsIn)
}

func TestGetLicenseeReport(t *testing.T) {
	licenseeReportsIn := testdata.GetLicenseeReportsData()
	for _, licenseeReportIn := range *licenseeReportsIn {
		GetLicenseeReportMocked(t, &licenseeReportIn)
		GetLicenseeReportFailErrMocked(t, &licenseeReportIn)
		GetLicenseeReportFailStatusMocked(t, &licenseeReportIn)
		GetLicenseeReportFailJSONMocked(t, &licenseeReportIn)
	}
}
