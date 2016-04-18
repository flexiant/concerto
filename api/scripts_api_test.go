package api

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewScriptServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewScriptService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetScriptList(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	GetScriptListMocked(t, scriptsIn)
	GetScriptListFailErrMocked(t, scriptsIn)
	GetScriptListFailStatusMocked(t, scriptsIn)
	GetScriptListFailJSONMocked(t, scriptsIn)
}

func TestGetScript(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	for _, scriptIn := range *scriptsIn {
		GetScriptMocked(t, &scriptIn)
		GetScriptFailErrMocked(t, &scriptIn)
		GetScriptFailStatusMocked(t, &scriptIn)
		GetScriptFailJSONMocked(t, &scriptIn)
	}
}

func TestCreateScript(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	for _, scriptIn := range *scriptsIn {
		CreateScriptMocked(t, &scriptIn)
		CreateScriptFailErrMocked(t, &scriptIn)
		CreateScriptFailStatusMocked(t, &scriptIn)
		CreateScriptFailJSONMocked(t, &scriptIn)
	}
}

func TestUpdateScript(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	for _, scriptIn := range *scriptsIn {
		UpdateScriptMocked(t, &scriptIn)
		UpdateScriptFailErrMocked(t, &scriptIn)
		UpdateScriptFailStatusMocked(t, &scriptIn)
		UpdateScriptFailJSONMocked(t, &scriptIn)
	}
}

func TestDeleteScript(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	for _, scriptIn := range *scriptsIn {
		DeleteScriptMocked(t, &scriptIn)
		DeleteScriptFailErrMocked(t, &scriptIn)
		DeleteScriptFailStatusMocked(t, &scriptIn)
	}
}
