package api

import (
	"testing"

	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewTemplateServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewTemplateService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetTemplateList(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	GetTemplateListMocked(t, templatesIn)
	GetTemplateListFailErrMocked(t, templatesIn)
	GetTemplateListFailStatusMocked(t, templatesIn)
	GetTemplateListFailJSONMocked(t, templatesIn)
}

func TestGetTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range *templatesIn {
		GetTemplateMocked(t, &templateIn)
		GetTemplateFailErrMocked(t, &templateIn)
		GetTemplateFailStatusMocked(t, &templateIn)
		GetTemplateFailJSONMocked(t, &templateIn)
	}
}

func TestCreateTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range *templatesIn {
		CreateTemplateMocked(t, &templateIn)
		CreateTemplateFailErrMocked(t, &templateIn)
		CreateTemplateFailStatusMocked(t, &templateIn)
		CreateTemplateFailJSONMocked(t, &templateIn)
	}
}

func TestUpdateTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range *templatesIn {
		UpdateTemplateMocked(t, &templateIn)
		UpdateTemplateFailErrMocked(t, &templateIn)
		UpdateTemplateFailStatusMocked(t, &templateIn)
		UpdateTemplateFailJSONMocked(t, &templateIn)
	}
}

func TestDeleteTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range *templatesIn {
		DeleteTemplateMocked(t, &templateIn)
		DeleteTemplateFailErrMocked(t, &templateIn)
		DeleteTemplateFailStatusMocked(t, &templateIn)
	}
}

func TestListTemplateScripts(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range *drsIn {
		GetTemplateScriptListMocked(t, drsIn, drIn.ID, drIn.Type)
		GetTemplateScriptListFailErrMocked(t, drsIn, drIn.ID, drIn.Type)
		GetTemplateScriptListFailStatusMocked(t, drsIn, drIn.ID, drIn.Type)
		GetTemplateScriptListFailJSONMocked(t, drsIn, drIn.ID, drIn.Type)
	}
}

func TestGetTemplateScript(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range *drsIn {
		GetTemplateScriptMocked(t, &drIn)
		GetTemplateScriptFailErrMocked(t, &drIn)
		GetTemplateScriptFailStatusMocked(t, &drIn)
		GetTemplateScriptFailJSONMocked(t, &drIn)
	}
}

func TestCreateTemplateScript(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range *drsIn {
		CreateTemplateScriptMocked(t, &drIn)
		CreateTemplateScriptFailErrMocked(t, &drIn)
		CreateTemplateScriptFailStatusMocked(t, &drIn)
		CreateTemplateScriptFailJSONMocked(t, &drIn)
	}
}

func TestUpdateTemplateScript(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range *drsIn {
		UpdateTemplateScriptMocked(t, &drIn)
		UpdateTemplateScriptFailErrMocked(t, &drIn)
		UpdateTemplateScriptFailStatusMocked(t, &drIn)
		UpdateTemplateScriptFailJSONMocked(t, &drIn)

	}
}

func TestDeleteTemplateScripts(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range *drsIn {
		DeleteTemplateScriptMocked(t, &drIn)
	}
}

func TestListTemplateServers(t *testing.T) {
	drsIn := testdata.GetTemplateServerData()
	for _, drIn := range *drsIn {
		GetTemplateServerListMocked(t, drsIn, drIn.ID)
	}
}

func TestReorderTemplateScript(t *testing.T) {
	tsIn := testdata.GetTemplateScriptData()

	// get template
	templateID := (*tsIn)[0].TemplateID

	// reorder
	num := len(*tsIn)
	reorder := make([]string, num, num)
	tsOut := make([]types.TemplateScript, num, num)

	num--
	for i, ts := range *tsIn {
		reorder[num-i] = ts.ID
		tsOut[num-i] = ts
	}

	ReorderTemplateScriptMocked(t, &tsOut, templateID, reorder)
}
