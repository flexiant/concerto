package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/testdata"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestDomainList subcommand
func TestDomainList(t *testing.T) {

	assert := assert.New(t)

	domainsTest, err := testdata.GetDomainData()
	assert.Nil(err, "Couldn't load domain test data")

	// only valid domains
	var domainsIn []types.Domain
	for _, domainTest := range domainsTest {
		if domainTest.FieldsOK {
			domainsIn = append(domainsIn, domainTest.Domain)
		}
	}

	domainsOut := api.GetTestDomainList(t, &domainsIn)

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

// TestDomainShow subcommand
func TestDomainShow(t *testing.T) {

	assert := assert.New(t)

	domainsTest, err := testdata.GetDomainData()
	assert.Nil(err, "Couldn't load domain test data")

	// only valid domains
	var domainsIn []types.Domain
	for _, domainTest := range domainsTest {
		if domainTest.FieldsOK {
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

	domainsTest, err := testdata.GetDomainData()
	assert.Nil(err, "Couldn't load domain test data")

	// only valid domains
	var domainsIn []types.Domain
	for _, domainTest := range domainsTest {
		if domainTest.FieldsOK {
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

// TestDomainUpdate subcommand
func TestDomainUpdate(t *testing.T) {

	assert := assert.New(t)

	domainsTest, err := testdata.GetDomainData()
	assert.Nil(err, "Couldn't load domain test data")

	// only valid domains
	var domainsIn []types.Domain
	for _, domainTest := range domainsTest {
		if domainTest.FieldsOK {
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
		cs.On("Put", fmt.Sprintf("/v1/dns/domains/%s", domainIn.ID), mapIn).Return(dOut, 200, nil)
		domainOut, err := ds.UpdateDomain(mapIn, domainIn.ID)
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

// TestDomainDelete subcommand
func TestDomainDelete(t *testing.T) {

	assert := assert.New(t)

	domainsTest, err := testdata.GetDomainData()
	assert.Nil(err, "Couldn't load domain test data")

	// only valid domains
	var domainsIn []types.Domain
	for _, domainTest := range domainsTest {
		if domainTest.FieldsOK {
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
		cs.On("Delete", fmt.Sprintf("/v1/dns/domains/%s", domainIn.ID)).Return(dIn, 200, nil)
		err = ds.DeleteDomain(domainIn.ID)
		assert.Nil(err, "Error getting domain list")
	}
}

func TestDomainRecordsList(t *testing.T) {

	assert := assert.New(t)

	domainsTest, err := testdata.GetDomainData()
	assert.Nil(err, "Couldn't load domain test data")

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := api.NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// only valid domains
	var domainsIn []types.Domain
	for _, domainTest := range domainsTest {
		if domainTest.FieldsOK {
			domainsIn = append(domainsIn, domainTest.Domain)
		}
	}

	for _, domainIn := range domainsIn {
		// to json
		dIn, err := json.Marshal(domainIn)
		assert.Nil(err, "Domain test data corrupted")

		// call service
		cs.On("Delete", fmt.Sprintf("/v1/dns/domains/%s", domainIn.ID)).Return(dIn, 200, nil)
		err = ds.DeleteDomain(domainIn.ID)
		assert.Nil(err, "Error getting domain list")
	}

	/*
		domainSvc, formatter := WireUpDomain(c)

		checkRequiredFlags(c, []string{"domain_id"}, formatter)
		domainRecords, err := domainSvc.ListDomainRecords(c.String("domain_id"))
		if err != nil {
			formatter.PrintFatal("Couldn't list domain records", err)
		}
		if err = formatter.PrintList(*domainRecords); err != nil {
			formatter.PrintFatal("Couldn't print/format result", err)
		}
	*/

}
