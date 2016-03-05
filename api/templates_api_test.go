package api

import (
	"github.com/flexiant/concerto/testdata"

	"testing"
)

func TestGetTemplateList(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	GetTemplateListMocked(t, templatesIn)
}

func TestGetTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range *templatesIn {
		GetTemplateMocked(t, &templateIn)
	}
}

func TestCreateTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range *templatesIn {
		CreateTemplateMocked(t, &templateIn)
	}
}

func TestUpdateTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range *templatesIn {
		UpdateTemplateMocked(t, &templateIn)
	}
}

func TestDeleteTemplate(t *testing.T) {
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range *templatesIn {
		DeleteTemplateMocked(t, &templateIn)
	}
}

func TestListTemplateScripts(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range *drsIn {
		GetTemplateScriptListMocked(t, drsIn, drIn.ID, drIn.Type)
	}
}

func TestGetTemplateScript(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range *drsIn {
		GetTemplateScriptMocked(t, &drIn)
	}
}

func TestCreateTemplateScript(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range *drsIn {
		CreateTemplateScriptMocked(t, &drIn)
	}
}

func TestUpdateTemplateScript(t *testing.T) {
	drsIn := testdata.GetTemplateScriptData()
	for _, drIn := range *drsIn {
		UpdateTemplateScriptMocked(t, &drIn)
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
