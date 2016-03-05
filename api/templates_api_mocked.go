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

// GetTemplateListMocked test mocked function
func GetTemplateListMocked(t *testing.T, templatesIn *[]types.Template) *[]types.Template {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemplateService(cs)
	assert.Nil(err, "Couldn't load template service")
	assert.NotNil(ds, "Template service not instanced")

	// to json
	dIn, err := json.Marshal(templatesIn)
	assert.Nil(err, "Template test data corrupted")

	// call service
	cs.On("Get", "/v1/blueprint/templates").Return(dIn, 200, nil)
	templatesOut, err := ds.GetTemplateList()
	assert.Nil(err, "Error getting template list")
	assert.Equal(*templatesIn, templatesOut, "GetTemplateList returned different templates")

	return &templatesOut
}

// GetTemplateMocked test mocked function
func GetTemplateMocked(t *testing.T, template *types.Template) *types.Template {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemplateService(cs)
	assert.Nil(err, "Couldn't load template service")
	assert.NotNil(ds, "Template service not instanced")

	// to json
	dIn, err := json.Marshal(template)
	assert.Nil(err, "Template test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/blueprint/templates/%s", template.ID)).Return(dIn, 200, nil)
	templateOut, err := ds.GetTemplate(template.ID)
	assert.Nil(err, "Error getting template")
	assert.Equal(*template, *templateOut, "GetTemplate returned different templates")

	return templateOut
}

// CreateTemplateMocked test mocked function
func CreateTemplateMocked(t *testing.T, templateIn *types.Template) *types.Template {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemplateService(cs)
	assert.Nil(err, "Couldn't load template service")
	assert.NotNil(ds, "Template service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*templateIn)
	assert.Nil(err, "Template test data corrupted")

	// to json
	dOut, err := json.Marshal(templateIn)
	assert.Nil(err, "Template test data corrupted")

	// call service
	cs.On("Post", "/v1/blueprint/templates/", mapIn).Return(dOut, 200, nil)
	templateOut, err := ds.CreateTemplate(mapIn)
	assert.Nil(err, "Error creating template list")
	assert.Equal(templateIn, templateOut, "CreateTemplate returned different templates")

	return templateOut
}

// UpdateTemplateMocked test mocked function
func UpdateTemplateMocked(t *testing.T, templateIn *types.Template) *types.Template {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemplateService(cs)
	assert.Nil(err, "Couldn't load template service")
	assert.NotNil(ds, "Template service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*templateIn)
	assert.Nil(err, "Template test data corrupted")

	// to json
	dOut, err := json.Marshal(templateIn)
	assert.Nil(err, "Template test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/blueprint/templates/%s", templateIn.ID), mapIn).Return(dOut, 200, nil)
	templateOut, err := ds.UpdateTemplate(mapIn, templateIn.ID)
	assert.Nil(err, "Error updating template list")
	assert.Equal(templateIn, templateOut, "UpdateTemplate returned different templates")

	return templateOut
}

// DeleteTemplateMocked test mocked function
func DeleteTemplateMocked(t *testing.T, templateIn *types.Template) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemplateService(cs)
	assert.Nil(err, "Couldn't load template service")
	assert.NotNil(ds, "Template service not instanced")

	// to json
	dIn, err := json.Marshal(templateIn)
	assert.Nil(err, "Template test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/blueprint/templates/%s", templateIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteTemplate(templateIn.ID)
	assert.Nil(err, "Error deleting template")

}

// GetTemplateScriptListMocked test mocked function
func GetTemplateScriptListMocked(t *testing.T, templateScriptsIn *[]types.TemplateScript, templateID string, scriptType string) *[]types.TemplateScript {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemplateService(cs)
	assert.Nil(err, "Couldn't load template service")
	assert.NotNil(ds, "Template service not instanced")

	// to json
	drsIn, err := json.Marshal(templateScriptsIn)
	assert.Nil(err, "Template script test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/blueprint/templates/%s/scripts?type=%s", templateID, scriptType)).Return(drsIn, 200, nil)
	drsOut, err := ds.GetTemplateScriptList(templateID, scriptType)
	assert.Nil(err, "Error getting template list")
	assert.Equal(*templateScriptsIn, *drsOut, "GetTemplateList returned different templates")

	return drsOut
}

// GetTemplateScriptMocked test mocked function
func GetTemplateScriptMocked(t *testing.T, dr *types.TemplateScript) *types.TemplateScript {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemplateService(cs)
	assert.Nil(err, "Couldn't load template service")
	assert.NotNil(ds, "Template service not instanced")

	// to json
	drIn, err := json.Marshal(dr)
	assert.Nil(err, "Template script test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s", dr.TemplateID, dr.ID)).Return(drIn, 200, nil)
	drOut, err := ds.GetTemplateScript(dr.TemplateID, dr.ID)
	assert.Nil(err, "Error getting template")
	assert.Equal(*dr, *drOut, "GetTemplateScript returned different template scripts")

	return drOut
}

// CreateTemplateScriptMocked test mocked function
func CreateTemplateScriptMocked(t *testing.T, dr *types.TemplateScript) *types.TemplateScript {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemplateService(cs)
	assert.Nil(err, "Couldn't load template service")
	assert.NotNil(ds, "Template service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*dr)
	assert.Nil(err, "Template script test data corrupted")

	// to json
	drIn, err := json.Marshal(dr)
	assert.Nil(err, "Template script test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/v1/blueprint/templates/%s/scripts", dr.TemplateID), mapIn).Return(drIn, 200, nil)
	drOut, err := ds.CreateTemplateScript(mapIn, dr.TemplateID)
	assert.Nil(err, "Error getting template")
	assert.Equal(*dr, *drOut, "CreateTemplateScript returned different template scripts")

	return drOut
}

// UpdateTemplateScriptMocked test mocked function
func UpdateTemplateScriptMocked(t *testing.T, dr *types.TemplateScript) *types.TemplateScript {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemplateService(cs)
	assert.Nil(err, "Couldn't load template service")
	assert.NotNil(ds, "Template service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*dr)
	assert.Nil(err, "Template script test data corrupted")

	// to json
	drIn, err := json.Marshal(dr)
	assert.Nil(err, "Template script test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s", dr.TemplateID, dr.ID), mapIn).Return(drIn, 200, nil)
	drOut, err := ds.UpdateTemplateScript(mapIn, dr.TemplateID, dr.ID)
	assert.Nil(err, "Error updating template list")
	assert.Equal(*dr, *drOut, "UpdateTemplateScript returned different template scripts")

	return drOut
}

// ReorderTemplateScriptMocked test mocked function
func ReorderTemplateScriptMocked(t *testing.T, dr *types.TemplateScript) *[]types.TemplateScript {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemplateService(cs)
	assert.Nil(err, "Couldn't load template service")
	assert.NotNil(ds, "Template service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*dr)
	assert.Nil(err, "Template script test data corrupted")

	// to json
	drIn, err := json.Marshal(dr)
	assert.Nil(err, "Template script test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/blueprint/templates/%s/scripts/reorder", dr.TemplateID), mapIn).Return(drIn, 200, nil)
	drOut, err := ds.ReorderTemplateScript(mapIn, dr.TemplateID)
	assert.Nil(err, "Error updating template list")
	assert.Equal(*dr, *drOut, "ReorderTemplateScript returned different template scripts")

	return drOut
}

// DeleteTemplateScriptMocked test mocked function
func DeleteTemplateScriptMocked(t *testing.T, dr *types.TemplateScript) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemplateService(cs)
	assert.Nil(err, "Couldn't load template service")
	assert.NotNil(ds, "Template service not instanced")

	// to json
	drIn, err := json.Marshal(dr)
	assert.Nil(err, "Template script test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s", dr.TemplateID, dr.ID)).Return(drIn, 200, nil)
	err = ds.DeleteTemplateScript(dr.TemplateID, dr.ID)
	assert.Nil(err, "Error deleting template script")
}

// GetTemplateServerListMocked test mocked function
func GetTemplateServerListMocked(t *testing.T, templateServersIn *[]types.TemplateServer, templateID string) *[]types.TemplateServer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewTemplateService(cs)
	assert.Nil(err, "Couldn't load template service")
	assert.NotNil(ds, "Template service not instanced")

	// to json
	drsIn, err := json.Marshal(templateServersIn)
	assert.Nil(err, "Template server test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/blueprint/templates/%s/servers", templateID)).Return(drsIn, 200, nil)
	drsOut, err := ds.GetTemplateServerList(templateID)
	assert.Nil(err, "Error getting template list")
	assert.Equal(*templateServersIn, *drsOut, "GetTemplateList returned different templates")

	return drsOut
}
