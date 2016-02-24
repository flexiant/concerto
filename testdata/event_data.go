package testdata

import (
	"github.com/flexiant/concerto/api/types"
	"time"
)

// GetEventData loads loads test data
func GetEventData() *[]types.Event {

	testEvents := []types.Event{
		{
			Id:          "fakeID0",
			Timestamp:   time.Date(2014, 1, 1, 12, 0, 0, 0, time.UTC),
			Level:       "fakeLevel0",
			Header:      "fakeHeader0",
			Description: "fakeDescription0",
		},
		{
			Id:          "fakeID1",
			Timestamp:   time.Date(2015, 1, 10, 11, 0, 0, 0, time.UTC),
			Level:       "fakeLevel1",
			Header:      "fakeHeader1",
			Description: "fakeDescription1",
		},
	}

	return &testEvents
}
