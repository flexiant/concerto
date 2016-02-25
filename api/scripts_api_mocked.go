package api

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
	scriptsOut, err := ds.GetScriptsList()
	assert.Nil(err, "Error getting script list")
	assert.Equal(*scriptsIn, scriptsOut, "GetScriptList returned different scripts")

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
