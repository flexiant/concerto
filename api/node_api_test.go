package api

import (
	"github.com/flexiant/concerto/testdata"
	"testing"
)

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
