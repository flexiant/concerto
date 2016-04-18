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

// GetNodeListFailErrMocked test mocked function
func GetNodeListFailErrMocked(t *testing.T, nodesIn *[]types.Node) *[]types.Node {

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
	cs.On("Get", "/v1/kaas/ships").Return(dIn, 200, fmt.Errorf("Mocked error"))
	nodesOut, err := ds.GetNodeList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(nodesOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &nodesOut
}

// GetNodeListFailStatusMocked test mocked function
func GetNodeListFailStatusMocked(t *testing.T, nodesIn *[]types.Node) *[]types.Node {

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
	cs.On("Get", "/v1/kaas/ships").Return(dIn, 499, nil)
	nodesOut, err := ds.GetNodeList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(nodesOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &nodesOut
}

// GetNodeListFailJSONMocked test mocked function
func GetNodeListFailJSONMocked(t *testing.T, nodesIn *[]types.Node) *[]types.Node {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodeService(cs)
	assert.Nil(err, "Couldn't load node service")
	assert.NotNil(ds, "Node service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/kaas/ships").Return(dIn, 200, nil)
	nodesOut, err := ds.GetNodeList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(nodesOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// CreateNodeFailErrMocked test mocked function
func CreateNodeFailErrMocked(t *testing.T, nodeIn *types.Node) *types.Node {

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
	cs.On("Post", "/v1/kaas/ships", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	nodeOut, err := ds.CreateNode(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(nodeOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return nodeOut
}

// CreateNodeFailStatusMocked test mocked function
func CreateNodeFailStatusMocked(t *testing.T, nodeIn *types.Node) *types.Node {

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
	cs.On("Post", "/v1/kaas/ships", mapIn).Return(dOut, 499, nil)
	nodeOut, err := ds.CreateNode(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(nodeOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return nodeOut
}

// CreateNodeFailJSONMocked test mocked function
func CreateNodeFailJSONMocked(t *testing.T, nodeIn *types.Node) *types.Node {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewNodeService(cs)
	assert.Nil(err, "Couldn't load node service")
	assert.NotNil(ds, "Node service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*nodeIn)
	assert.Nil(err, "Node test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/kaas/ships", mapIn).Return(dIn, 200, nil)
	nodeOut, err := ds.CreateNode(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(nodeOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// DeleteNodeFailErrMocked test mocked function
func DeleteNodeFailErrMocked(t *testing.T, nodeIn *types.Node) {

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
	cs.On("Delete", fmt.Sprintf("/v1/kaas/ships/%s", nodeIn.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	err = ds.DeleteNode(nodeIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// DeleteNodeFailStatusMocked test mocked function
func DeleteNodeFailStatusMocked(t *testing.T, nodeIn *types.Node) {

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
	cs.On("Delete", fmt.Sprintf("/v1/kaas/ships/%s", nodeIn.Id)).Return(dIn, 499, nil)
	err = ds.DeleteNode(nodeIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
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

// StartNodeFailErrMocked test mocked function
func StartNodeFailErrMocked(t *testing.T, nodeIn *types.Node) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/ships/%s/start", nodeIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	err = ds.StartNode(mapIn, nodeIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

}

// StartNodeFailStatusMocked test mocked function
func StartNodeFailStatusMocked(t *testing.T, nodeIn *types.Node) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/ships/%s/start", nodeIn.Id), mapIn).Return(dOut, 499, nil)
	err = ds.StartNode(mapIn, nodeIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

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

// StopNodeFailErrMocked test mocked function
func StopNodeFailErrMocked(t *testing.T, nodeIn *types.Node) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/ships/%s/stop", nodeIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	err = ds.StopNode(mapIn, nodeIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

}

// StopNodeFailStatusMocked test mocked function
func StopNodeFailStatusMocked(t *testing.T, nodeIn *types.Node) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/ships/%s/stop", nodeIn.Id), mapIn).Return(dOut, 499, nil)
	err = ds.StopNode(mapIn, nodeIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

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

// RestartNodeFailErrMocked test mocked function
func RestartNodeFailErrMocked(t *testing.T, nodeIn *types.Node) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/ships/%s/restart", nodeIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	err = ds.RestartNode(mapIn, nodeIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

}

// RestartNodeFailStatusMocked test mocked function
func RestartNodeFailStatusMocked(t *testing.T, nodeIn *types.Node) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/ships/%s/restart", nodeIn.Id), mapIn).Return(dOut, 499, nil)
	err = ds.RestartNode(mapIn, nodeIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
