package api

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// TemplateService manages template operations
type TemplateService struct {
	concertoService utils.ConcertoService
}

// NewTemplateService returns a Concerto template service
func NewTemplateService(concertoService utils.ConcertoService) (*TemplateService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &TemplateService{
		concertoService: concertoService,
	}, nil
}

// GetTemplateList returns the list of templates as an array of Template
func (tp *TemplateService) GetTemplateList() (templates []types.Template, err error) {
	log.Debug("GetTemplateList")

	data, status, err := tp.concertoService.Get("/v1/blueprint/templates")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {

	}

	if err = json.Unmarshal(data, &templates); err != nil {
		return nil, err
	}

	return templates, nil
}

// GetTemplate returns a template by its ID
func (tp *TemplateService) GetTemplate(ID string) (template *types.Template, err error) {
	log.Debug("GetTemplate")

	data, status, err := tp.concertoService.Get(fmt.Sprintf("/v1/blueprint/templates/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &template); err != nil {
		return nil, err
	}

	return template, nil
}

// CreateTemplate creates a template
func (tp *TemplateService) CreateTemplate(templateVector *map[string]string) (template *types.Template, err error) {
	log.Debug("CreateTemplate")

	data, status, err := tp.concertoService.Post("/v1/blueprint/templates/", templateVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &template); err != nil {
		return nil, err
	}

	return template, nil
}

// UpdateTemplate updates a template by its ID
func (tp *TemplateService) UpdateTemplate(templateVector *map[string]string, ID string) (template *types.Template, err error) {
	log.Debug("UpdateTemplate")

	data, status, err := tp.concertoService.Put(fmt.Sprintf("/v1/blueprint/templates/%s", ID), templateVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &template); err != nil {
		return nil, err
	}

	return template, nil
}

// DeleteTemplate deletes a template by its ID
func (tp *TemplateService) DeleteTemplate(ID string) (err error) {
	log.Debug("DeleteTemplate")

	data, status, err := tp.concertoService.Delete(fmt.Sprintf("/v1/blueprint/templates/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ================ Template Script =================

// GetTemplateScriptList returns a list of templateScript by template ID
func (tp *TemplateService) GetTemplateScriptList(templateID string, scriptType string) (templateScript *[]types.TemplateScript, err error) {
	log.Debug("ListTemplateScripts")

	data, status, err := tp.concertoService.Get(fmt.Sprintf("/v1/blueprint/templates/%s/scripts?type=%s", templateID, scriptType))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templateScript); err != nil {
		return nil, err
	}

	return templateScript, nil
}

// GetTemplateScript returns a templateScript
func (tp *TemplateService) GetTemplateScript(templateID string, ID string) (templateScript *types.TemplateScript, err error) {
	log.Debug("GetTemplateScript")

	data, status, err := tp.concertoService.Get(fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s", templateID, ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templateScript); err != nil {
		return nil, err
	}

	return templateScript, nil
}

// CreateTemplateScript returns a list of templateScript
func (tp *TemplateService) CreateTemplateScript(templateScriptVector *map[string]string, templateID string) (templateScript *types.TemplateScript, err error) {
	log.Debug("CreateTemplateScript")

	data, status, err := tp.concertoService.Post(fmt.Sprintf("/v1/blueprint/templates/%s/scripts", templateID), templateScriptVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templateScript); err != nil {
		return nil, err
	}

	return templateScript, nil
}

// UpdateTemplateScript returns a list of templateScript
func (tp *TemplateService) UpdateTemplateScript(templateScriptVector *map[string]string, templateID string, ID string) (templateScript *types.TemplateScript, err error) {
	log.Debug("UpdateTemplateScript")

	data, status, err := tp.concertoService.Put(fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s", templateID, ID), templateScriptVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templateScript); err != nil {
		return nil, err
	}

	return templateScript, nil
}

// DeleteTemplateScript deletes a template record
func (tp *TemplateService) DeleteTemplateScript(templateID string, ID string) (err error) {
	log.Debug("DeleteTemplateScript")

	data, status, err := tp.concertoService.Delete(fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s", templateID, ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// ReorderTemplateScript returns a list of templateScript
func (tp *TemplateService) ReorderTemplateScript(templateID string) (templateScript *types.TemplateScript, err error) {
	log.Debug("UpdateTemplateScript")

	data, status, err := tp.concertoService.Put(fmt.Sprintf("/v1/blueprint/templates/%s/scripts/reorder", templateID), nil)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templateScript); err != nil {
		return nil, err
	}

	return templateScript, nil
}

// ================ Template Servers =================

// GetTemplateServersList returns a list of templateServers by template ID
func (tp *TemplateService) GetTemplateServersList(templateID string) (templateScript *[]types.TemplateScript, err error) {
	log.Debug("ListTemplateScripts")

	data, status, err := tp.concertoService.Get(fmt.Sprintf("/v1/blueprint/templates/%s/servers", templateID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &templateScript); err != nil {
		return nil, err
	}

	return templateScript, nil
}
