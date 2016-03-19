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

// GetWorkspaceListMocked test mocked function
func GetWorkspaceListMocked(t *testing.T, workspacesIn *[]types.Workspace) *[]types.Workspace {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspace service")
	assert.NotNil(ds, "Workspace service not instanced")

	// to json
	dIn, err := json.Marshal(workspacesIn)
	assert.Nil(err, "Workspace test data corrupted")

	// call service
	cs.On("Get", "/v1/cloud/workspaces").Return(dIn, 200, nil)
	workspacesOut, err := ds.GetWorkspaceList()
	assert.Nil(err, "Error getting workspace list")
	assert.Equal(*workspacesIn, workspacesOut, "GetWorkspaceList returned different workspaces")

	return &workspacesOut
}

// GetWorkspaceMocked test mocked function
func GetWorkspaceMocked(t *testing.T, workspace *types.Workspace) *types.Workspace {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspace service")
	assert.NotNil(ds, "Workspace service not instanced")

	// to json
	dIn, err := json.Marshal(workspace)
	assert.Nil(err, "Workspace test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/workspaces/%s", workspace.Id)).Return(dIn, 200, nil)
	workspaceOut, err := ds.GetWorkspace(workspace.Id)
	assert.Nil(err, "Error getting workspace")
	assert.Equal(*workspace, *workspaceOut, "GetWorkspace returned different workspaces")

	return workspaceOut
}

// CreateWorkspaceMocked test mocked function
func CreateWorkspaceMocked(t *testing.T, workspaceIn *types.Workspace) *types.Workspace {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspace service")
	assert.NotNil(ds, "Workspace service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*workspaceIn)
	assert.Nil(err, "Workspace test data corrupted")

	// to json
	dOut, err := json.Marshal(workspaceIn)
	assert.Nil(err, "Workspace test data corrupted")

	// call service
	cs.On("Post", "/v1/cloud/workspaces/", mapIn).Return(dOut, 200, nil)
	workspaceOut, err := ds.CreateWorkspace(mapIn)
	assert.Nil(err, "Error creating workspace list")
	assert.Equal(workspaceIn, workspaceOut, "CreateWorkspace returned different workspaces")

	return workspaceOut
}

// UpdateWorkspaceMocked test mocked function
func UpdateWorkspaceMocked(t *testing.T, workspaceIn *types.Workspace) *types.Workspace {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspace service")
	assert.NotNil(ds, "Workspace service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*workspaceIn)
	assert.Nil(err, "Workspace test data corrupted")

	// to json
	dOut, err := json.Marshal(workspaceIn)
	assert.Nil(err, "Workspace test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/workspaces/%s", workspaceIn.Id), mapIn).Return(dOut, 200, nil)
	workspaceOut, err := ds.UpdateWorkspace(mapIn, workspaceIn.Id)
	assert.Nil(err, "Error updating workspace list")
	assert.Equal(workspaceIn, workspaceOut, "UpdateWorkspace returned different workspaces")

	return workspaceOut
}

// DeleteWorkspaceMocked test mocked function
func DeleteWorkspaceMocked(t *testing.T, workspaceIn *types.Workspace) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspace service")
	assert.NotNil(ds, "Workspace service not instanced")

	// to json
	dIn, err := json.Marshal(workspaceIn)
	assert.Nil(err, "Workspace test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/cloud/workspaces/%s", workspaceIn.Id)).Return(dIn, 200, nil)
	err = ds.DeleteWorkspace(workspaceIn.Id)
	assert.Nil(err, "Error deleting workspace")

}

// GetWorkspaceServerListMocked test mocked function
func GetWorkspaceServerListMocked(t *testing.T, workspaceServersIn *[]types.WorkspaceServer, workspaceId string) *[]types.WorkspaceServer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspace service")
	assert.NotNil(ds, "Workspace service not instanced")

	// to json
	drsIn, err := json.Marshal(workspaceServersIn)
	assert.Nil(err, "Workspace record test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/workspaces/%s/servers", workspaceId)).Return(drsIn, 200, nil)
	drsOut, err := ds.GetWorkspaceServerList(workspaceId)
	assert.Nil(err, "Error getting workspace list")
	assert.Equal(*workspaceServersIn, *drsOut, "GetWorkspaceList returned different workspaces")

	return drsOut
}
