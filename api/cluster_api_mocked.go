package api

import (
	"encoding/json"
	"fmt"
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

// CreateClusterMocked test mocked function
func CreateClusterMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// to json
	dOut, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Post", "/v1/kaas/fleets", mapIn).Return(dOut, 200, nil)
	clusterOut, err := ds.CreateCluster(mapIn)
	assert.Nil(err, "Error creating cluster list")
	assert.Equal(clusterIn, clusterOut, "CreateCluster returned different clusters")

	return clusterOut
}

// DeleteClusterMocked test mocked function
func DeleteClusterMocked(t *testing.T, clusterIn *types.Cluster) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// to json
	dIn, err := json.Marshal(clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/kaas/fleets/%s", clusterIn.Id)).Return(dIn, 200, nil)
	err = ds.DeleteCluster(clusterIn.Id)
	assert.Nil(err, "Error deleting cluster")

}
