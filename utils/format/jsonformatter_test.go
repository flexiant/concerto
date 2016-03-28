package format

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintItemDomainJSON(t *testing.T) {

	assert := assert.New(t)
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range *domainsIn {

		domainOut := api.GetDomainMocked(t, &domainIn)

		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		InitializeFormatter("json", mockOut)
		f := GetFormatter()
		assert.NotNil(f, "Formatter")

		err := f.PrintItem(*domainOut)
		assert.Nil(err, "JSON formatter PrintItem error")
		mockOut.Flush()

		// TODO add more accurate parsing
		assert.Regexp("^\\{\\\"id\\\":.*\\}", b.String(), "JSON output didn't match regular expression")
	}
}

func TestPrintItemTemplateJSON(t *testing.T) {

	assert := assert.New(t)
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range *templatesIn {

		templateOut := api.GetTemplateMocked(t, &templateIn)

		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		InitializeFormatter("json", mockOut)
		f := GetFormatter()
		assert.NotNil(f, "Formatter")

		err := f.PrintItem(*templateOut)
		assert.Nil(err, "Text formatter PrintItem error")
		mockOut.Flush()

		// TODO add more accurate parsing
		assert.Regexp("^\\{\\\"id\\\":.*\\}", b.String(), "JSON output didn't match regular expression")
	}
}

func TestPrintListDomainsJSON(t *testing.T) {

	assert := assert.New(t)
	domainsIn := testdata.GetDomainData()
	domainOut := api.GetDomainListMocked(t, domainsIn)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("json", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(domainOut)
	assert.Nil(err, "JSON formatter PrintItem error")
	mockOut.Flush()

	// TODO add more accurate parsing
	assert.Regexp("^\\[\\{\\\"id\\\":.*\\}\\]", b.String(), "JSON output didn't match regular expression")
}

func TestPrintListTemplateJSON(t *testing.T) {

	assert := assert.New(t)
	templatesIn := testdata.GetTemplateData()
	templatesOut := api.GetTemplateListMocked(t, templatesIn)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("json", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(*templatesOut)
	assert.Nil(err, "Text formatter PrintItem error")
	mockOut.Flush()

	// TODO add more accurate parsing
	assert.Regexp("^\\[\\{\\\"id\\\":.*\\}\\]", b.String(), "JSON output didn't match regular expression")
}

func TestPrintErrorJSON(t *testing.T) {

	assert := assert.New(t)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)

	InitializeFormatter("json", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	f.PrintError("testing errors", fmt.Errorf("this is a test error %s", "TEST"))
	mockOut.Flush()

	assert.Regexp("^\\{\\\"type\\\":\\\"Error\\\",\\\"context\\\":\\\"testing errors\\\",\\\"message\\\":\\\"this is a test error TEST\\\"\\}", b.String(), "JSON output didn't match regular expression")
}
