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
			Name:     "test.zero",
			Content:  "another.server.com",
			TTL:      300,
			Prio:     0,
			ServerID: "56ae7bfd47ac3d0008000066",
			DomainID: "fakeID0",
		},
		{
			ID:       "fakeID0.1",
			Type:     "A",
			Name:     "test.one",
			Content:  "10.20.30.40",
			TTL:      1800,
			Prio:     10,
			ServerID: "56ae7bfd47ac3d0008000066",
			DomainID: "fakeID0",
		},
		{
			ID:       "fakeID0.2",
			Type:     "A",
			Name:     "test.two",
			Content:  "20.30.40.50",
			TTL:      18800,
			Prio:     20,
			ServerID: "56ae7bfd47ac3d0008000066",
			DomainID: "fakeID0",
		},
		{
			ID:       "fakeID1.0",
			Type:     "A",
			Name:     "test.1.zero",
			Content:  "1.1.1.1",
			TTL:      100,
			Prio:     0,
			ServerID: "56ae7bfd47ac3d0008000066",
			DomainID: "fakeID1",
		},
	}

	return &testDomainRecords
}
