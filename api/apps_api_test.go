package api

import (
	"testing"

	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewAppServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewAppService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetAppList(t *testing.T) {
	appsIn := testdata.GetAppData()
	GetAppListMocked(t, appsIn)
	GetAppListFailErrMocked(t, appsIn)
	GetAppListFailStatusMocked(t, appsIn)
	GetAppListFailJSONMocked(t, appsIn)
}

func TestDeployApp(t *testing.T) {
	appsIn := testdata.GetAppData()
	for _, appIn := range *appsIn {
		DeployAppMocked(t, &appIn)
		DeployAppFailErrMocked(t, &appIn)
		DeployAppFailStatusMocked(t, &appIn)
		DeployAppFailJSONMocked(t, &appIn)
	}
}
