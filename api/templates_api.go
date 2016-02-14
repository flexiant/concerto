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
