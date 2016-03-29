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
	eventsIn := testdata.GetNodeData()
	GetNodeListMocked(t, eventsIn)
}

func TestCreateNode(t *testing.T) {
	nodesIn := testdata.GetNodeData()
	for _, nodeIn := range *nodesIn {
		CreateNodeMocked(t, &nodeIn)
	}
}

func TestDeleteNode(t *testing.T) {
	nodesIn := testdata.GetNodeData()
	for _, nodeIn := range *nodesIn {
		DeleteNodeMocked(t, &nodeIn)
	}
}

func TestStartNode(t *testing.T) {
	nodesIn := testdata.GetNodeData()
	for _, nodeIn := range *nodesIn {
		StartNodeMocked(t, &nodeIn)
	}
}

func TestStopNode(t *testing.T) {
	nodesIn := testdata.GetNodeData()
	for _, nodeIn := range *nodesIn {
		StopNodeMocked(t, &nodeIn)
	}
}

func TestRestartNode(t *testing.T) {
	nodesIn := testdata.GetNodeData()
	for _, nodeIn := range *nodesIn {
		RestartNodeMocked(t, &nodeIn)
	}
}
