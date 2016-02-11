package testdata

import "github.com/flexiant/concerto/api/types"

// DomainTest holds test data domains
type DomainTest struct {
	Domain   types.Domain
	FieldsOK bool // true if all mandatory fields are informed
}

// DomainRecordTest holds test data domains records
type DomainRecordTest struct {
	DomainRecord types.DomainRecord
	FieldsOK     bool // true if all mandatory fields are informed
}

var testDomains []DomainTest
var testDomainRecords []DomainRecordTest

// GetDomainData loads loads test data
func GetDomainData() ([]DomainTest, error) {

	testDomains = []DomainTest{
		{
			Domain: types.Domain{
				ID:      "fakeID0",
				Name:    "fakeName0",
				TTL:     1000,
				Contact: "fakeContact0",
				Minimum: 10,
				Enabled: true,
			},
			FieldsOK: true,
		},
		{
			Domain: types.Domain{
				ID:      "fakeID1",
				Name:    "fakeName1",
				TTL:     1001,
				Contact: "fakeContact1",
				Minimum: 11,
				Enabled: false,
			},
			FieldsOK: true,
		},
	}

	return testDomains, nil
}

// GetDomainRecordData loads test data
func GetDomainRecordData() ([]DomainRecordTest, error) {

	testDomainRecords = []DomainRecordTest{
		{
			DomainRecord: types.DomainRecord{
				ID:       "fakeID0.0",
				Type:     "CNAME",
				Name:     "otherserver",
				Content:  "my.server.com",
				TTL:      300,
				Prio:     10,
				ServerID: "server",
				DomainID: "fakeID0",
			},
			FieldsOK: true,
		},
	}

	return testDomainRecords, nil
}
