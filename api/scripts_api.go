package api

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// ScriptService manages scripts operations
type ScriptService struct {
	concertoService utils.ConcertoService
}

// NewScriptsService returns a Concerto script service
func NewScriptsService(concertoService utils.ConcertoService) (*ScriptService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &ScriptService{
		concertoService: concertoService,
	}, nil
}

// GetScriptsList returns the list of scripts as an array of Scripts
func (sc *ScriptService) GetScriptsList() (scripts []types.Script, err error) {
	log.Debug("GetScriptsList")

	data, status, err := sc.concertoService.Get("/v1/blueprint/scripts")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {

	}

	if err = json.Unmarshal(data, &scripts); err != nil {
		return nil, err
	}

	return scripts, nil
}

// GetScript returns a script by its ID
func (sc *ScriptService) GetScript(ID string) (script *types.Script, err error) {
	log.Debug("GetScript")

	data, status, err := sc.concertoService.Get(fmt.Sprintf("/v1/blueprint/scripts/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}

// CreateScript creates a script
func (sc *ScriptService) CreateScript(scriptVector *map[string]interface{}) (script *types.Script, err error) {
	log.Debug("CreateScript")

	data, status, err := sc.concertoService.Post("/v1/blueprint/scripts", scriptVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}

// UpdateScript updates a script by its ID
func (sc *ScriptService) UpdateScript(scriptVector *map[string]interface{}, ID string) (script *types.Script, err error) {
	log.Debug("UpdateScript")

	data, status, err := sc.concertoService.Put(fmt.Sprintf("/v1/blueprint/scripts/%s", ID), scriptVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &script); err != nil {
		return nil, err
	}

	return script, nil
}

// DeleteScript deletes a script by its ID
func (sc *ScriptService) DeleteScript(ID string) (err error) {
	log.Debug("DeleteScript")

	data, status, err := sc.concertoService.Delete(fmt.Sprintf("/v1/blueprint/scripts/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
