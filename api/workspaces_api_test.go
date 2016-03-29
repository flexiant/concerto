package api

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWorkspaceServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewWorkspaceService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetWorkspaceList(t *testing.T) {
	workspacesIn := testdata.GetWorkspaceData()
	GetWorkspaceListMocked(t, workspacesIn)
}

func TestGetWorkspace(t *testing.T) {
	workspacesIn := testdata.GetWorkspaceData()
	for _, workspaceIn := range *workspacesIn {
		GetWorkspaceMocked(t, &workspaceIn)
	}
}

func TestCreateWorkspace(t *testing.T) {
	workspacesIn := testdata.GetWorkspaceData()
	for _, workspaceIn := range *workspacesIn {
		CreateWorkspaceMocked(t, &workspaceIn)
	}
}

func TestUpdateWorkspace(t *testing.T) {
	workspacesIn := testdata.GetWorkspaceData()
	for _, workspaceIn := range *workspacesIn {
		UpdateWorkspaceMocked(t, &workspaceIn)
	}
}

func TestDeleteWorkspace(t *testing.T) {
	workspacesIn := testdata.GetWorkspaceData()
	for _, workspaceIn := range *workspacesIn {
		DeleteWorkspaceMocked(t, &workspaceIn)
	}
}

func TestListWorkspaceServers(t *testing.T) {
	drsIn := testdata.GetWorkspaceServerData()
	for _, drIn := range *drsIn {
		GetWorkspaceServerListMocked(t, drsIn, drIn.Id)
	}
}
