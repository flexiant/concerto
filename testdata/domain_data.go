package testdata

import "github.com/flexiant/concerto/api/types"

// GetDomainData loads loads test data
func GetDomainData() *[]types.Domain {

	testDomains := []types.Domain{
		{
			ID:      "fakeID0",
			Name:    "fakeName0",
			TTL:     1000,
			Contact: "fakeContact0",
			Minimum: 10,
			Enabled: true,
		},
		{
			ID:      "fakeID1",
			Name:    "fakeName1",
			TTL:     1001,
			Contact: "fakeContact1",
			Minimum: 11,
			Enabled: false,
		},
	}

	return &testDomains
}

// GetDomainRecordData loads test data
func GetDomainRecordData() *[]types.DomainRecord {

	testDomainRecords := []types.DomainRecord{
		{
			ID:       "fakeID0.0",
			Type:     "CNAME",
			Name:     "otherserver",
			Content:  "my.server.com",
			TTL:      300,
			Prio:     10,
			ServerID: "server",
			DomainID: "fakeID0",
		},
	}

	return &testDomainRecords
}
