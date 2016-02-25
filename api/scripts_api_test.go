package api

import (
	"github.com/flexiant/concerto/testdata"
	"testing"
)

func TestGetScriptList(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	GetScriptListMocked(t, scriptsIn)
}

func TestGetScript(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	for _, scriptIn := range *scriptsIn {
		GetScriptMocked(t, &scriptIn)
	}
}

func TestCreateScript(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	for _, scriptIn := range *scriptsIn {
		CreateScriptMocked(t, &scriptIn)
	}
}

func TestUpdateScript(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	for _, scriptIn := range *scriptsIn {
		UpdateScriptMocked(t, &scriptIn)
	}
}

func TestDeleteScript(t *testing.T) {
	scriptsIn := testdata.GetScriptData()
	for _, scriptIn := range *scriptsIn {
		DeleteScriptMocked(t, &scriptIn)
	}
}
