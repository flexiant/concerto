package api

import (
	"github.com/flexiant/concerto/testdata"

	"testing"
)

func TestGetClusterList(t *testing.T) {
	clustersIn := testdata.GetClusterData()
	GetClusterListMocked(t, clustersIn)
}
