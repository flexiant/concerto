package api

import (
	"github.com/flexiant/concerto/testdata"

	"testing"
)

func TestGetClusterList(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	GetClusterListMocked(t, clustersIn)
}

func TestCreateCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range *clustersIn {
		CreateClusterMocked(t, &clusterIn)
	}
}

func TestDeleteCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range *clustersIn {
		DeleteClusterMocked(t, &clusterIn)
	}
}

func TestStartCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range *clustersIn {
		StartClusterMocked(t, &clusterIn)
	}
}

func TestStopCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range *clustersIn {
		StopClusterMocked(t, &clusterIn)
	}
}

func TestEmptyCluster(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	for _, clusterIn := range *clustersIn {
		EmptyClusterMocked(t, &clusterIn)
	}
}
