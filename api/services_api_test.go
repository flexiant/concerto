package api

import (
	"github.com/flexiant/concerto/testdata"

	"testing"
)

func TestGetServiceList(t *testing.T) {
	servicesIn := testdata.GetServiceData()
	GetServiceListMocked(t, servicesIn)
}

func TestGetService(t *testing.T) {
	servicesIn := testdata.GetServiceData()
	for _, serviceIn := range *servicesIn {
		GetServiceMocked(t, &serviceIn)
	}
}
