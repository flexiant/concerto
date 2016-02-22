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

func TestPrintItemDomainTXT(t *testing.T) {

	assert := assert.New(t)
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range *domainsIn {

		domainOut := api.GetDomainMocked(t, &domainIn)

		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		InitializeFormatter("text", mockOut)
		f := GetFormatter()
		assert.NotNil(f, "Formatter")

		err := f.PrintItem(*domainOut)
		assert.Nil(err, "Text formatter PrintItem error")
		mockOut.Flush()

		// TODO add more accurate parsing
		assert.Regexp("^ID:\\ *.*\n*.\n", b.String(), "Text output didn't match regular expression")

	}
}

func TestPrintListDomainsTXT(t *testing.T) {

	assert := assert.New(t)
	domainsIn := testdata.GetDomainData()
	domainOut := api.GetDomainListMocked(t, domainsIn)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(*domainOut)
	assert.Nil(err, "Text formatter PrintItem error")
	mockOut.Flush()

	// TODO add more accurate parsing
	assert.Regexp(fmt.Sprintf("^ID.*\n%s.*\n.*", (*domainOut)[0].ID), b.String(), "Text output didn't match regular expression")
}

func TestPrintError(t *testing.T) {

	assert := assert.New(t)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)

	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	f.PrintError("testing errors", fmt.Errorf("this is a test error %s", "TEST"))
	mockOut.Flush()

	assert.Regexp("^ERROR:.*\n -> .*\n", b.String(), "Text output didn't match regular expression")
}
