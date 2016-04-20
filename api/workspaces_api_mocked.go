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

// GetWorkspaceListFailErrMocked test mocked function
func GetWorkspaceListFailErrMocked(t *testing.T, workspacesIn *[]types.Workspace) *[]types.Workspace {

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
	cs.On("Get", "/v1/cloud/workspaces").Return(dIn, 200, fmt.Errorf("Mocked error"))
	workspacesOut, err := ds.GetWorkspaceList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(workspacesOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &workspacesOut
}

// GetWorkspaceListFailStatusMocked test mocked function
func GetWorkspaceListFailStatusMocked(t *testing.T, workspacesIn *[]types.Workspace) *[]types.Workspace {

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
	cs.On("Get", "/v1/cloud/workspaces").Return(dIn, 499, nil)
	workspacesOut, err := ds.GetWorkspaceList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(workspacesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &workspacesOut
}

// GetWorkspaceListFailJSONMocked test mocked function
func GetWorkspaceListFailJSONMocked(t *testing.T, workspacesIn *[]types.Workspace) *[]types.Workspace {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspace service")
	assert.NotNil(ds, "Workspace service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/cloud/workspaces").Return(dIn, 200, nil)
	workspacesOut, err := ds.GetWorkspaceList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(workspacesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// GetWorkspaceFailErrMocked test mocked function
func GetWorkspaceFailErrMocked(t *testing.T, workspace *types.Workspace) *types.Workspace {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/workspaces/%s", workspace.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	workspaceOut, err := ds.GetWorkspace(workspace.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(workspaceOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return workspaceOut
}

// GetWorkspaceFailStatusMocked test mocked function
func GetWorkspaceFailStatusMocked(t *testing.T, workspace *types.Workspace) *types.Workspace {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/workspaces/%s", workspace.Id)).Return(dIn, 499, nil)
	workspaceOut, err := ds.GetWorkspace(workspace.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(workspaceOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return workspaceOut
}

// GetWorkspaceFailJSONMocked test mocked function
func GetWorkspaceFailJSONMocked(t *testing.T, workspace *types.Workspace) *types.Workspace {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspace service")
	assert.NotNil(ds, "Workspace service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/workspaces/%s", workspace.Id)).Return(dIn, 200, nil)
	workspaceOut, err := ds.GetWorkspace(workspace.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(workspaceOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// CreateWorkspaceFailErrMocked test mocked function
func CreateWorkspaceFailErrMocked(t *testing.T, workspaceIn *types.Workspace) *types.Workspace {

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
	cs.On("Post", "/v1/cloud/workspaces/", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	workspaceOut, err := ds.CreateWorkspace(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(workspaceOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return workspaceOut
}

// CreateWorkspaceFailStatusMocked test mocked function
func CreateWorkspaceFailStatusMocked(t *testing.T, workspaceIn *types.Workspace) *types.Workspace {

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
	cs.On("Post", "/v1/cloud/workspaces/", mapIn).Return(dOut, 499, nil)
	workspaceOut, err := ds.CreateWorkspace(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(workspaceOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return workspaceOut
}

// CreateWorkspaceFailJSONMocked test mocked function
func CreateWorkspaceFailJSONMocked(t *testing.T, workspaceIn *types.Workspace) *types.Workspace {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspace service")
	assert.NotNil(ds, "Workspace service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*workspaceIn)
	assert.Nil(err, "Workspace test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/cloud/workspaces/", mapIn).Return(dIn, 200, nil)
	workspaceOut, err := ds.CreateWorkspace(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(workspaceOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// UpdateWorkspaceFailErrMocked test mocked function
func UpdateWorkspaceFailErrMocked(t *testing.T, workspaceIn *types.Workspace) *types.Workspace {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/workspaces/%s", workspaceIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	workspaceOut, err := ds.UpdateWorkspace(mapIn, workspaceIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(workspaceOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return workspaceOut
}

// UpdateWorkspaceFailStatusMocked test mocked function
func UpdateWorkspaceFailStatusMocked(t *testing.T, workspaceIn *types.Workspace) *types.Workspace {

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
	cs.On("Put", fmt.Sprintf("/v1/cloud/workspaces/%s", workspaceIn.Id), mapIn).Return(dOut, 499, nil)
	workspaceOut, err := ds.UpdateWorkspace(mapIn, workspaceIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(workspaceOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return workspaceOut
}

// UpdateWorkspaceFailJSONMocked test mocked function
func UpdateWorkspaceFailJSONMocked(t *testing.T, workspaceIn *types.Workspace) *types.Workspace {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspace service")
	assert.NotNil(ds, "Workspace service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*workspaceIn)
	assert.Nil(err, "Workspace test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/cloud/workspaces/%s", workspaceIn.Id), mapIn).Return(dIn, 200, nil)
	workspaceOut, err := ds.UpdateWorkspace(mapIn, workspaceIn.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(workspaceOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// DeleteWorkspaceFailErrMocked test mocked function
func DeleteWorkspaceFailErrMocked(t *testing.T, workspaceIn *types.Workspace) {

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
	cs.On("Delete", fmt.Sprintf("/v1/cloud/workspaces/%s", workspaceIn.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	err = ds.DeleteWorkspace(workspaceIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// DeleteWorkspaceFailStatusMocked test mocked function
func DeleteWorkspaceFailStatusMocked(t *testing.T, workspaceIn *types.Workspace) {

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
	cs.On("Delete", fmt.Sprintf("/v1/cloud/workspaces/%s", workspaceIn.Id)).Return(dIn, 499, nil)
	err = ds.DeleteWorkspace(workspaceIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
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

// GetWorkspaceServerListFailErrMocked test mocked function
func GetWorkspaceServerListFailErrMocked(t *testing.T, workspaceServersIn *[]types.WorkspaceServer, workspaceId string) *[]types.WorkspaceServer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspaceServer service")
	assert.NotNil(ds, "WorkspaceServer service not instanced")

	// to json
	dIn, err := json.Marshal(workspaceServersIn)
	assert.Nil(err, "WorkspaceServer test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/workspaces/%s/servers", workspaceId)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	workspaceServersOut, err := ds.GetWorkspaceServerList(workspaceId)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(workspaceServersOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return workspaceServersOut
}

// GetWorkspaceServerListFailStatusMocked test mocked function
func GetWorkspaceServerListFailStatusMocked(t *testing.T, workspaceServersIn *[]types.WorkspaceServer, workspaceId string) *[]types.WorkspaceServer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspaceServer service")
	assert.NotNil(ds, "WorkspaceServer service not instanced")

	// to json
	dIn, err := json.Marshal(workspaceServersIn)
	assert.Nil(err, "WorkspaceServer test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/workspaces/%s/servers", workspaceId)).Return(dIn, 499, nil)
	workspaceServersOut, err := ds.GetWorkspaceServerList(workspaceId)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(workspaceServersOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return workspaceServersOut
}

// GetWorkspaceServerListFailJSONMocked test mocked function
func GetWorkspaceServerListFailJSONMocked(t *testing.T, workspaceServersIn *[]types.WorkspaceServer, workspaceId string) *[]types.WorkspaceServer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWorkspaceService(cs)
	assert.Nil(err, "Couldn't load workspaceServer service")
	assert.NotNil(ds, "WorkspaceServer service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/workspaces/%s/servers", workspaceId)).Return(dIn, 200, nil)
	workspaceServersOut, err := ds.GetWorkspaceServerList(workspaceId)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(workspaceServersOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return workspaceServersOut
}
