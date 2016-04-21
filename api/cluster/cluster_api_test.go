package cluster

import (
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClusterServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewClusterService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetClusterList(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	GetClusterListMocked(t, clustersIn)
	GetClusterListFailErrMocked(t, clustersIn)
	GetClusterListFailStatusMocked(t, clustersIn)
	GetClusterListFailJSONMocked(t, clustersIn)
}

func TestCreateCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range *clustersIn {
		CreateClusterMocked(t, &clusterIn)
		CreateClusterFailErrMocked(t, &clusterIn)
		CreateClusterFailStatusMocked(t, &clusterIn)
		CreateClusterFailJSONMocked(t, &clusterIn)
	}
}

func TestDeleteCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range *clustersIn {
		DeleteClusterMocked(t, &clusterIn)
		DeleteClusterFailErrMocked(t, &clusterIn)
		DeleteClusterFailStatusMocked(t, &clusterIn)
	}
}

func TestStartCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range *clustersIn {
		StartClusterMocked(t, &clusterIn)
		StartClusterFailErrMocked(t, &clusterIn)
		StartClusterFailStatusMocked(t, &clusterIn)
	}
}

func TestStopCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range *clustersIn {
		StopClusterMocked(t, &clusterIn)
		StopClusterFailErrMocked(t, &clusterIn)
		StopClusterFailStatusMocked(t, &clusterIn)
	}
}

func TestEmptyCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range *clustersIn {
		EmptyClusterMocked(t, &clusterIn)
		EmptyClusterFailErrMocked(t, &clusterIn)
		EmptyClusterFailStatusMocked(t, &clusterIn)
	}
}
