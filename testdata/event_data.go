package testdata

import (
	"github.com/flexiant/concerto/api/types"
	// "time"
)

// GetEventData loads loads test data
func GetEventData() *[]types.Event {

	testEvents := []types.Event{
		{
			Id: "fakeID0",
			// Timestamp:   time.Time.Clock(10),
			Level:       "fakeLevel0",
			Header:      "fakeHeader0",
			Description: "fakeDescription0",
		},
		{
			Id: "fakeID1",
			// Timestamp:   time.Time.Clock(20),
			Level:       "fakeLevel1",
			Header:      "fakeHeader1",
			Description: "fakeDescription1",
		},
	}

	return &testEvents
}
