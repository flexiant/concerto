package api

import (
	"github.com/flexiant/concerto/testdata"
	"testing"
)

func TestGetEventList(t *testing.T) {
	eventsIn := testdata.GetEventData()
	GetEventListMocked(t, eventsIn)
}

func TestGetSysEventList(t *testing.T) {
	eventsIn := testdata.GetEventData()
	GetSysEventListMocked(t, eventsIn)
}
