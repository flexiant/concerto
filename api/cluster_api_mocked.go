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

// GetClusterListFailErrMocked test mocked function
func GetClusterListFailErrMocked(t *testing.T, clustersIn *[]types.Cluster) *[]types.Cluster {

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
	cs.On("Get", "/v1/kaas/fleets").Return(dIn, 200, fmt.Errorf("Mocked error"))
	clustersOut, err := ds.GetClusterList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(clustersOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &clustersOut
}

// GetClusterListFailStatusMocked test mocked function
func GetClusterListFailStatusMocked(t *testing.T, clustersIn *[]types.Cluster) *[]types.Cluster {

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
	cs.On("Get", "/v1/kaas/fleets").Return(dIn, 499, nil)
	clustersOut, err := ds.GetClusterList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(clustersOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &clustersOut
}

// GetClusterListFailJSONMocked test mocked function
func GetClusterListFailJSONMocked(t *testing.T, clustersIn *[]types.Cluster) *[]types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/kaas/fleets").Return(dIn, 200, nil)
	clustersOut, err := ds.GetClusterList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(clustersOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// CreateClusterFailErrMocked test mocked function
func CreateClusterFailErrMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

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
	cs.On("Post", "/v1/kaas/fleets", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	clusterOut, err := ds.CreateCluster(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return clusterOut
}

// CreateClusterFailStatusMocked test mocked function
func CreateClusterFailStatusMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

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
	cs.On("Post", "/v1/kaas/fleets", mapIn).Return(dOut, 499, nil)
	clusterOut, err := ds.CreateCluster(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return clusterOut
}

// CreateClusterFailJSONMocked test mocked function
func CreateClusterFailJSONMocked(t *testing.T, clusterIn *types.Cluster) *types.Cluster {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewClusterService(cs)
	assert.Nil(err, "Couldn't load cluster service")
	assert.NotNil(ds, "Cluster service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*clusterIn)
	assert.Nil(err, "Cluster test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/kaas/fleets", mapIn).Return(dIn, 200, nil)
	clusterOut, err := ds.CreateCluster(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(clusterOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// DeleteClusterFailErrMocked test mocked function
func DeleteClusterFailErrMocked(t *testing.T, clusterIn *types.Cluster) {

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
	cs.On("Delete", fmt.Sprintf("/v1/kaas/fleets/%s", clusterIn.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	err = ds.DeleteCluster(clusterIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// DeleteClusterFailStatusMocked test mocked function
func DeleteClusterFailStatusMocked(t *testing.T, clusterIn *types.Cluster) {

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
	cs.On("Delete", fmt.Sprintf("/v1/kaas/fleets/%s", clusterIn.Id)).Return(dIn, 499, nil)
	err = ds.DeleteCluster(clusterIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// StartClusterMocked test mocked function
func StartClusterMocked(t *testing.T, clusterIn *types.Cluster) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/fleets/%s/start", clusterIn.Id), mapIn).Return(dOut, 200, nil)
	err = ds.StartCluster(mapIn, clusterIn.Id)
	assert.Nil(err, "Error updating cluster list")
}

// StartClusterFailErrMocked test mocked function
func StartClusterFailErrMocked(t *testing.T, clusterIn *types.Cluster) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/fleets/%s/start", clusterIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	err = ds.StartCluster(mapIn, clusterIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

}

// StartClusterFailStatusMocked test mocked function
func StartClusterFailStatusMocked(t *testing.T, clusterIn *types.Cluster) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/fleets/%s/start", clusterIn.Id), mapIn).Return(dOut, 499, nil)
	err = ds.StartCluster(mapIn, clusterIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

}

// StopClusterMocked test mocked function
func StopClusterMocked(t *testing.T, clusterIn *types.Cluster) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/fleets/%s/stop", clusterIn.Id), mapIn).Return(dOut, 200, nil)
	err = ds.StopCluster(mapIn, clusterIn.Id)
	assert.Nil(err, "Error updating cluster list")
}

// StopClusterFailErrMocked test mocked function
func StopClusterFailErrMocked(t *testing.T, clusterIn *types.Cluster) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/fleets/%s/stop", clusterIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	err = ds.StopCluster(mapIn, clusterIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

}

// StopClusterFailStatusMocked test mocked function
func StopClusterFailStatusMocked(t *testing.T, clusterIn *types.Cluster) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/fleets/%s/stop", clusterIn.Id), mapIn).Return(dOut, 499, nil)
	err = ds.StopCluster(mapIn, clusterIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

}

// EmptyClusterMocked test mocked function
func EmptyClusterMocked(t *testing.T, clusterIn *types.Cluster) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/fleets/%s/empty", clusterIn.Id), mapIn).Return(dOut, 200, nil)
	err = ds.EmptyCluster(mapIn, clusterIn.Id)
	assert.Nil(err, "Error updating cluster list")

}

// EmptyClusterFailErrMocked test mocked function
func EmptyClusterFailErrMocked(t *testing.T, clusterIn *types.Cluster) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/fleets/%s/empty", clusterIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	err = ds.EmptyCluster(mapIn, clusterIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

}

// EmptyClusterFailStatusMocked test mocked function
func EmptyClusterFailStatusMocked(t *testing.T, clusterIn *types.Cluster) {

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
	cs.On("Put", fmt.Sprintf("/v1/kaas/fleets/%s/empty", clusterIn.Id), mapIn).Return(dOut, 499, nil)
	err = ds.EmptyCluster(mapIn, clusterIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
