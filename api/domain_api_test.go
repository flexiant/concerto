package api

import (
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDomainList(t *testing.T) {
	domainsIn := getOKDomains(t)
	GetTestDomainList(t, domainsIn)
}

func TestGetDomain(t *testing.T) {
	domainsIn := getOKDomains(t)
	for _, domainIn := range *domainsIn {
		GetTestDomain(t, &domainIn)
	}
}

func getOKDomains(t *testing.T) *[]types.Domain {
	assert := assert.New(t)
	domainsTest, err := testdata.GetDomainData()
	assert.Nil(err, "Couldn't load domain test data")
	var domains []types.Domain
	for _, domainTest := range domainsTest {
		if domainTest.FieldsOK {
			domains = append(domains, domainTest.Domain)
		}
	}
	return &domains
}
