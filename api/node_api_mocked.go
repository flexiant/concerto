package api

import (
	"encoding/json"
	// "fmt"
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
