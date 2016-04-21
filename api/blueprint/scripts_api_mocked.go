package blueprint

import (
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO exclude from release compile

// GetScriptListMocked test mocked function
func GetScriptListMocked(t *testing.T, scriptsIn *[]types.Script) *[]types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(scriptsIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", "/v1/blueprint/scripts").Return(dIn, 200, nil)
	scriptsOut, err := ds.GetScriptList()
	assert.Nil(err, "Error getting script list")
	assert.Equal(*scriptsIn, scriptsOut, "GetScriptList returned different scripts")

	return &scriptsOut
}

// GetScriptListFailErrMocked test mocked function
func GetScriptListFailErrMocked(t *testing.T, scriptsIn *[]types.Script) *[]types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(scriptsIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", "/v1/blueprint/scripts").Return(dIn, 200, fmt.Errorf("Mocked error"))
	scriptsOut, err := ds.GetScriptList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scriptsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &scriptsOut
}

// GetScriptListFailStatusMocked test mocked function
func GetScriptListFailStatusMocked(t *testing.T, scriptsIn *[]types.Script) *[]types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(scriptsIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", "/v1/blueprint/scripts").Return(dIn, 499, nil)
	scriptsOut, err := ds.GetScriptList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scriptsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &scriptsOut
}

// GetScriptListFailJSONMocked test mocked function
func GetScriptListFailJSONMocked(t *testing.T, scriptsIn *[]types.Script) *[]types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/blueprint/scripts").Return(dIn, 200, nil)
	scriptsOut, err := ds.GetScriptList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scriptsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &scriptsOut
}

// GetScriptMocked test mocked function
func GetScriptMocked(t *testing.T, script *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(script)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/blueprint/scripts/%s", script.ID)).Return(dIn, 200, nil)
	scriptOut, err := ds.GetScript(script.ID)
	assert.Nil(err, "Error getting script")
	assert.Equal(*script, *scriptOut, "GetScript returned different scripts")

	return scriptOut
}

// GetScriptFailErrMocked test mocked function
func GetScriptFailErrMocked(t *testing.T, script *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(script)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/blueprint/scripts/%s", script.ID)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	scriptOut, err := ds.GetScript(script.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return scriptOut
}

// GetScriptFailStatusMocked test mocked function
func GetScriptFailStatusMocked(t *testing.T, script *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(script)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/blueprint/scripts/%s", script.ID)).Return(dIn, 499, nil)
	scriptOut, err := ds.GetScript(script.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return scriptOut
}

// GetScriptFailJSONMocked test mocked function
func GetScriptFailJSONMocked(t *testing.T, script *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/blueprint/scripts/%s", script.ID)).Return(dIn, 200, nil)
	scriptOut, err := ds.GetScript(script.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return scriptOut
}

// CreateScriptMocked test mocked function
func CreateScriptMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dOut, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Post", "/v1/blueprint/scripts", mapIn).Return(dOut, 200, nil)
	scriptOut, err := ds.CreateScript(mapIn)
	assert.Nil(err, "Error creating script list")
	assert.Equal(scriptIn, scriptOut, "CreateScript returned different scripts")

	return scriptOut
}

// CreateScriptFailErrMocked test mocked function
func CreateScriptFailErrMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dOut, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Post", "/v1/blueprint/scripts", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	scriptOut, err := ds.CreateScript(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return scriptOut
}

// CreateScriptFailStatusMocked test mocked function
func CreateScriptFailStatusMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dOut, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Post", "/v1/blueprint/scripts", mapIn).Return(dOut, 499, nil)
	scriptOut, err := ds.CreateScript(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return scriptOut
}

// CreateScriptFailJSONMocked test mocked function
func CreateScriptFailJSONMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/blueprint/scripts", mapIn).Return(dIn, 200, nil)
	scriptOut, err := ds.CreateScript(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return scriptOut
}

// UpdateScriptMocked test mocked function
func UpdateScriptMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dOut, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/blueprint/scripts/%s", scriptIn.ID), mapIn).Return(dOut, 200, nil)
	scriptOut, err := ds.UpdateScript(mapIn, scriptIn.ID)
	assert.Nil(err, "Error updating script list")
	assert.Equal(scriptIn, scriptOut, "UpdateScript returned different scripts")

	return scriptOut
}

// UpdateScriptFailErrMocked test mocked function
func UpdateScriptFailErrMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dOut, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/blueprint/scripts/%s", scriptIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	scriptOut, err := ds.UpdateScript(mapIn, scriptIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return scriptOut
}

// UpdateScriptFailStatusMocked test mocked function
func UpdateScriptFailStatusMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// to json
	dOut, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/blueprint/scripts/%s", scriptIn.ID), mapIn).Return(dOut, 499, nil)
	scriptOut, err := ds.UpdateScript(mapIn, scriptIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return scriptOut
}

// UpdateScriptFailJSONMocked test mocked function
func UpdateScriptFailJSONMocked(t *testing.T, scriptIn *types.Script) *types.Script {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/blueprint/scripts/%s", scriptIn.ID), mapIn).Return(dIn, 200, nil)
	scriptOut, err := ds.UpdateScript(mapIn, scriptIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(scriptOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return scriptOut
}

// DeleteScriptMocked test mocked function
func DeleteScriptMocked(t *testing.T, scriptIn *types.Script) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/blueprint/scripts/%s", scriptIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteScript(scriptIn.ID)
	assert.Nil(err, "Error deleting script")

}

// DeleteScriptFailErrMocked test mocked function
func DeleteScriptFailErrMocked(t *testing.T, scriptIn *types.Script) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/blueprint/scripts/%s", scriptIn.ID)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	err = ds.DeleteScript(scriptIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// DeleteScriptFailStatusMocked test mocked function
func DeleteScriptFailStatusMocked(t *testing.T, scriptIn *types.Script) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewScriptService(cs)
	assert.Nil(err, "Couldn't load script service")
	assert.NotNil(ds, "Script service not instanced")

	// to json
	dIn, err := json.Marshal(scriptIn)
	assert.Nil(err, "Script test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/blueprint/scripts/%s", scriptIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteScript(scriptIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
