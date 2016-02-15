package api

import (
	"encoding/json"
	// "fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetClusterListMocked test mocked function
func GetClusterListMocked(t *testing.T, clustersIn *[]types.Cluster) *[]types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clustersIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Get", "/v1/kaas/fleets").Return(dIn, 200, nil)
	clustersOut, err := ds.GetClusterList()
	assert.Nil(err, "Error getting cluster list")
	assert.Equal(*clustersIn, clustersOut, "GetClusterList returned different clusters")

	return &clustersOut
}
