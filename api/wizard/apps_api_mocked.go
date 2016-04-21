package wizard

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// GetAppListMocked test mocked function
func GetAppListMocked(t *testing.T, appsIn *[]types.WizardApp) *[]types.WizardApp {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// to json
	dIn, err := json.Marshal(appsIn)
	assert.Nil(err, "App test data corrupted")

	// call service
	cs.On("Get", "/v1/wizard/apps").Return(dIn, 200, nil)
	appsOut, err := ds.GetAppList()
	assert.Nil(err, "Error getting app list")
	assert.Equal(*appsIn, appsOut, "GetAppList returned different apps")

	return &appsOut
}

// GetAppListFailErrMocked test mocked function
func GetAppListFailErrMocked(t *testing.T, appsIn *[]types.WizardApp) *[]types.WizardApp {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// to json
	dIn, err := json.Marshal(appsIn)
	assert.Nil(err, "App test data corrupted")

	// call service
	cs.On("Get", "/v1/wizard/apps").Return(dIn, 200, fmt.Errorf("Mocked error"))
	appsOut, err := ds.GetAppList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(appsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &appsOut
}

// GetAppListFailStatusMocked test mocked function
func GetAppListFailStatusMocked(t *testing.T, appsIn *[]types.WizardApp) *[]types.WizardApp {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// to json
	dIn, err := json.Marshal(appsIn)
	assert.Nil(err, "App test data corrupted")

	// call service
	cs.On("Get", "/v1/wizard/apps").Return(dIn, 499, nil)
	appsOut, err := ds.GetAppList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(appsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &appsOut
}

// GetAppListFailJSONMocked test mocked function
func GetAppListFailJSONMocked(t *testing.T, appsIn *[]types.WizardApp) *[]types.WizardApp {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/wizard/apps").Return(dIn, 200, nil)
	appsOut, err := ds.GetAppList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(appsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &appsOut
}

// DeployAppMocked test mocked function
func DeployAppMocked(t *testing.T, appIn *types.WizardApp) *types.WizardApp {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*appIn)
	assert.Nil(err, "App test data corrupted")

	// to json
	dOut, err := json.Marshal(appIn)
	assert.Nil(err, "App test data corrupted")

	// call service
	cs.On("Post", "/v1/wizard/apps/", mapIn).Return(dOut, 200, nil)
	appOut, err := ds.DeployApp(mapIn)
	assert.Nil(err, "Error creating app list")
	assert.Equal(appIn, appOut, "DeployApp returned different apps")

	return appOut
}

// DeployAppFailErrMocked test mocked function
func DeployAppFailErrMocked(t *testing.T, appIn *types.WizardApp) *types.WizardApp {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*appIn)
	assert.Nil(err, "App test data corrupted")

	// to json
	dOut, err := json.Marshal(appIn)
	assert.Nil(err, "App test data corrupted")

	// call service
	cs.On("Post", "/v1/wizard/apps/", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	appOut, err := ds.DeployApp(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(appOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return appOut
}

// DeployAppFailStatusMocked test mocked function
func DeployAppFailStatusMocked(t *testing.T, appIn *types.WizardApp) *types.WizardApp {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*appIn)
	assert.Nil(err, "App test data corrupted")

	// to json
	dOut, err := json.Marshal(appIn)
	assert.Nil(err, "App test data corrupted")

	// call service
	cs.On("Post", "/v1/wizard/apps/", mapIn).Return(dOut, 499, nil)
	appOut, err := ds.DeployApp(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(appOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return appOut
}

// DeployAppFailJSONMocked test mocked function
func DeployAppFailJSONMocked(t *testing.T, appIn *types.WizardApp) *types.WizardApp {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewAppService(cs)
	assert.Nil(err, "Couldn't load app service")
	assert.NotNil(ds, "App service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*appIn)
	assert.Nil(err, "App test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/wizard/apps/", mapIn).Return(dIn, 200, nil)
	appOut, err := ds.DeployApp(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(appOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return appOut
}
