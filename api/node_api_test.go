package api

import (
	"github.com/flexiant/concerto/testdata"
	"testing"
)

func TestGetNodeList(t *testing.T) {
	eventsIn := testdata.GetNodeData()
	GetNodeListMocked(t, eventsIn)
}
