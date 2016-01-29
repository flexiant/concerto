package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGetDomainList subcommand
func TestGetDomainList(t *testing.T) {

	assert := assert.New(t)

	domainsTest, err := GetDomainData()
	assert.Nil(err, "Couldn't load domain test data")

	// only valid domains
	var domainsIn []api.Domain
	for _, domainTest := range domainsTest {
		if domainTest.Valid {
			domainsIn = append(domainsIn, domainTest.Domain)
		}
	}

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := api.NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domainsIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Get", "/v1/dns/domains").Return(dIn, 200, nil)
	domainsOut, err := ds.GetDomainList()
	assert.Nil(err, "Error getting domain list")
	assert.Equal(domainsIn, domainsOut, "GetDomainList returned different domains")

	// TODO iterate all formatters
	// write output
	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	f := format.NewJSONFormatter(mockOut)
	assert.NotNil(f, "Formatter")
	err = f.PrintList(domainsOut)
	assert.Nil(err, "JSON Formatter Printlinst error")
	mockOut.Flush()

	// TODO add more accurate parsing
	assert.Regexp("\\[\\{\\\"id\\\":.*\\}\\]", b.String(), "JSON Output didn't match regular expression")
}

// TestGetDomain subcommand
func TestGetDomain(t *testing.T) {

	assert := assert.New(t)

	domainsTest, err := GetDomainData()
	assert.Nil(err, "Couldn't load domain test data")

	// only valid domains
	var domainsIn []api.Domain
	for _, domainTest := range domainsTest {
		if domainTest.Valid {
			domainsIn = append(domainsIn, domainTest.Domain)
		}
	}

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := api.NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	for _, domainIn := range domainsIn {
		// to json
		dIn, err := json.Marshal(domainIn)
		assert.Nil(err, "Domain test data corrupted")

		// call service
		cs.On("Get", fmt.Sprintf("/v1/dns/domains/%s", domainIn.ID)).Return(dIn, 200, nil)
		domainOut, err := ds.GetDomain(domainIn.ID)
		assert.Nil(err, "Error getting domain list")
		assert.Equal(domainIn, *domainOut, "GetDomainList returned different domains")

		// TODO iterate all formatters
		// write output
		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		f := format.NewJSONFormatter(mockOut)
		assert.NotNil(f, "Formatter")
		err = f.PrintList(domainOut)
		assert.Nil(err, "JSON Formatter Printlinst error")
		mockOut.Flush()

		// TODO add more accurate parsing
		assert.Regexp("\\{\\\"id\\\":.*\\}", b.String(), "JSON Output didn't match regular expression")
	}
}

// TestDomainCreate subcommand
func TestDomainCreate(t *testing.T) {

	assert := assert.New(t)

	domainsTest, err := GetDomainData()
	assert.Nil(err, "Couldn't load domain test data")

	// only valid domains
	var domainsIn []api.Domain
	for _, domainTest := range domainsTest {
		if domainTest.Valid {
			domainsIn = append(domainsIn, domainTest.Domain)
		}
	}

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := api.NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	for _, domainIn := range domainsIn {

		// convertMap
		mapIn, err := itemConvertParams(domainIn)
		assert.Nil(err, "Domain test data corrupted")

		// to json
		dOut, err := json.Marshal(domainIn)
		assert.Nil(err, "Domain test data corrupted")

		// call service
		cs.On("Post", "/v1/dns/domains/", mapIn).Return(dOut, 200, nil)
		domainOut, err := ds.CreateDomain(mapIn)
		assert.Nil(err, "Error creating domain list")
		assert.Equal(domainIn, *domainOut, "GetDomainList returned different domains")

		// TODO iterate all formatters
		// write output
		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		f := format.NewJSONFormatter(mockOut)
		assert.NotNil(f, "Formatter")
		err = f.PrintList(domainOut)
		assert.Nil(err, "JSON Formatter Printlinst error")
		mockOut.Flush()

		// TODO add more accurate parsing
		assert.Regexp("\\{\\\"id\\\":.*\\}", b.String(), "JSON Output didn't match regular expression")

	}

}

type DomainTest struct {
	Domain api.Domain
	Valid  bool
}

var testDomains []DomainTest

// GetDomainData loads json test data from "./testdata"
func GetDomainData() ([]DomainTest, error) {

	testDomains = []DomainTest{
		{
			Domain: api.Domain{
				ID:      "fakeID0",
				Name:    "fakeName0",
				TTL:     1000,
				Contact: "fakeContact0",
				Minimum: 10,
				Enabled: true,
			},
			Valid: true,
		},
		{
			Domain: api.Domain{ID: "fakeID1",
				Name:    "fakeName1",
				TTL:     1001,
				Contact: "fakeContact1",
				Minimum: 11,
				Enabled: false,
			},
			Valid: true,
		},
	}

	return testDomains, nil
}
