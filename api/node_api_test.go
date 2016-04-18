package api

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNodeServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewNodeService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetNodeList(t *testing.T) {
	nodesIn := testdata.GetNodeData()
	GetNodeListMocked(t, nodesIn)
	GetNodeListFailErrMocked(t, nodesIn)
	GetNodeListFailStatusMocked(t, nodesIn)
	GetNodeListFailJSONMocked(t, nodesIn)
}

func TestCreateNode(t *testing.T) {
	nodesIn := testdata.GetNodeData()
	for _, nodeIn := range *nodesIn {
		CreateNodeMocked(t, &nodeIn)
		CreateNodeFailErrMocked(t, &nodeIn)
		CreateNodeFailStatusMocked(t, &nodeIn)
		CreateNodeFailJSONMocked(t, &nodeIn)
	}
}

func TestDeleteNode(t *testing.T) {
	nodesIn := testdata.GetNodeData()
	for _, nodeIn := range *nodesIn {
		DeleteNodeMocked(t, &nodeIn)
		DeleteNodeFailErrMocked(t, &nodeIn)
		DeleteNodeFailStatusMocked(t, &nodeIn)
	}
}

func TestStartNode(t *testing.T) {
	nodesIn := testdata.GetNodeData()
	for _, nodeIn := range *nodesIn {
		StartNodeMocked(t, &nodeIn)
		StartNodeFailErrMocked(t, &nodeIn)
		StartNodeFailStatusMocked(t, &nodeIn)
	}
}

func TestStopNode(t *testing.T) {
	nodesIn := testdata.GetNodeData()
	for _, nodeIn := range *nodesIn {
		StopNodeMocked(t, &nodeIn)
		StopNodeFailErrMocked(t, &nodeIn)
		StopNodeFailStatusMocked(t, &nodeIn)
	}
}

func TestRestartNode(t *testing.T) {
	nodesIn := testdata.GetNodeData()
	for _, nodeIn := range *nodesIn {
		RestartNodeMocked(t, &nodeIn)
		RestartNodeFailErrMocked(t, &nodeIn)
		RestartNodeFailStatusMocked(t, &nodeIn)
	}
}
