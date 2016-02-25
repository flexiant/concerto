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

// GetNodeListMocked test mocked function
func GetNodeListMocked(t *testing.T, nodesIn *[]types.Node) *[]types.Node {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodeService(cs)
	assert.Nil(err, "Couldn't load node service")
	assert.NotNil(ds, "Node service not instanced")

	// to json
	dIn, err := json.Marshal(nodesIn)
	assert.Nil(err, "Node test data corrupted")

	// call service
	cs.On("Get", "/v1/kaas/ships").Return(dIn, 200, nil)
	nodesOut, err := ds.GetNodeList()
	assert.Nil(err, "Error getting node list")
	assert.Equal(*nodesIn, nodesOut, "GetNodeList returned different nodes")

	return &nodesOut
}

// CreateNodeMocked test mocked function
func CreateNodeMocked(t *testing.T, nodeIn *types.Node) *types.Node {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodeService(cs)
	assert.Nil(err, "Couldn't load node service")
	assert.NotNil(ds, "Node service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodeIn)
	assert.Nil(err, "Node test data corrupted")

	// to json
	dOut, err := json.Marshal(nodeIn)
	assert.Nil(err, "Node test data corrupted")

	// call service
	cs.On("Post", "/v1/kaas/ships", mapIn).Return(dOut, 200, nil)
	nodeOut, err := ds.CreateNode(mapIn)
	assert.Nil(err, "Error creating node list")
	assert.Equal(nodeIn, nodeOut, "CreateNode returned different nodes")

	return nodeOut
}

// DeleteNodeMocked test mocked function
func DeleteNodeMocked(t *testing.T, nodeIn *types.Node) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodeService(cs)
	assert.Nil(err, "Couldn't load node service")
	assert.NotNil(ds, "Node service not instanced")

	// to json
	dIn, err := json.Marshal(nodeIn)
	assert.Nil(err, "Node test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/kaas/ships/%s", nodeIn.Id)).Return(dIn, 200, nil)
	err = ds.DeleteNode(nodeIn.Id)
	assert.Nil(err, "Error deleting node")

}

// StartNodeMocked test mocked function
func StartNodeMocked(t *testing.T, nodeIn *types.Node) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodeService(cs)
	assert.Nil(err, "Couldn't load node service")
	assert.NotNil(ds, "Node service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodeIn)
	assert.Nil(err, "Node test data corrupted")

	// to json
	dOut, err := json.Marshal(nodeIn)
	assert.Nil(err, "Node test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/kaas/ships/%s/start", nodeIn.Id), mapIn).Return(dOut, 200, nil)
	err = ds.StartNode(mapIn, nodeIn.Id)
	assert.Nil(err, "Error updating node list")
}

// StopNodeMocked test mocked function
func StopNodeMocked(t *testing.T, nodeIn *types.Node) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodeService(cs)
	assert.Nil(err, "Couldn't load node service")
	assert.NotNil(ds, "Node service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodeIn)
	assert.Nil(err, "Node test data corrupted")

	// to json
	dOut, err := json.Marshal(nodeIn)
	assert.Nil(err, "Node test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/kaas/ships/%s/stop", nodeIn.Id), mapIn).Return(dOut, 200, nil)
	err = ds.StopNode(mapIn, nodeIn.Id)
	assert.Nil(err, "Error updating node list")
}

// RestartNodeMocked test mocked function
func RestartNodeMocked(t *testing.T, nodeIn *types.Node) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodeService(cs)
	assert.Nil(err, "Couldn't load node service")
	assert.NotNil(ds, "Node service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodeIn)
	assert.Nil(err, "Node test data corrupted")

	// to json
	dOut, err := json.Marshal(nodeIn)
	assert.Nil(err, "Node test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/kaas/ships/%s/restart", nodeIn.Id), mapIn).Return(dOut, 200, nil)
	err = ds.RestartNode(mapIn, nodeIn.Id)
	assert.Nil(err, "Error updating node list")
}
