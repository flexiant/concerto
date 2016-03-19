package api

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// WorkspaceService manages workspace operations
type WorkspaceService struct {
	concertoService utils.ConcertoService
}

// NewWorkspaceService returns a Concerto workspace service
func NewWorkspaceService(concertoService utils.ConcertoService) (*WorkspaceService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &WorkspaceService{
		concertoService: concertoService,
	}, nil
}

// GetWorkspaceList returns the list of workspaces as an array of Workspace
func (dm *WorkspaceService) GetWorkspaceList() (workspaces []types.Workspace, err error) {
	log.Debug("GetWorkspaceList")

	data, status, err := dm.concertoService.Get("/v1/cloud/workspaces")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {

	}

	if err = json.Unmarshal(data, &workspaces); err != nil {
		return nil, err
	}

	return workspaces, nil
}

// GetWorkspace returns a workspace by its ID
func (dm *WorkspaceService) GetWorkspace(ID string) (workspace *types.Workspace, err error) {
	log.Debug("GetWorkspace")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/v1/cloud/workspaces/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &workspace); err != nil {
		return nil, err
	}

	return workspace, nil
}

// CreateWorkspace creates a workspace
func (dm *WorkspaceService) CreateWorkspace(workspaceVector *map[string]interface{}) (workspace *types.Workspace, err error) {
	log.Debug("CreateWorkspace")

	data, status, err := dm.concertoService.Post("/v1/cloud/workspaces/", workspaceVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &workspace); err != nil {
		return nil, err
	}

	return workspace, nil
}

// UpdateWorkspace updates a workspace by its ID
func (dm *WorkspaceService) UpdateWorkspace(workspaceVector *map[string]interface{}, ID string) (workspace *types.Workspace, err error) {
	log.Debug("UpdateWorkspace")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/v1/cloud/workspaces/%s", ID), workspaceVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &workspace); err != nil {
		return nil, err
	}

	return workspace, nil
}

// DeleteWorkspace deletes a workspace by its ID
func (dm *WorkspaceService) DeleteWorkspace(ID string) (err error) {
	log.Debug("DeleteWorkspace")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/v1/cloud/workspaces/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// GetWorkspaceServerList returns a list of workspaceServer by workspace ID
func (dm *WorkspaceService) GetWorkspaceServerList(workspaceID string) (workspaceServer *[]types.WorkspaceServer, err error) {
	log.Debug("ListWorkspaceServers")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/v1/cloud/workspaces/%s/servers", workspaceID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &workspaceServer); err != nil {
		return nil, err
	}

	return workspaceServer, nil
}
