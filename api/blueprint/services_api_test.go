package blueprint

import (
	"testing"

	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewServicesServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewServicesService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetServiceList(t *testing.T) {
	servicesIn := testdata.GetServiceData()
	GetServiceListMocked(t, servicesIn)
	GetServiceListFailErrMocked(t, servicesIn)
	GetServiceListFailStatusMocked(t, servicesIn)
	GetServiceListFailJSONMocked(t, servicesIn)
}

func TestGetService(t *testing.T) {
	servicesIn := testdata.GetServiceData()
	for _, serviceIn := range *servicesIn {
		GetServiceMocked(t, &serviceIn)
		GetServiceFailErrMocked(t, &serviceIn)
		GetServiceFailStatusMocked(t, &serviceIn)
		GetServiceFailJSONMocked(t, &serviceIn)
	}
}
