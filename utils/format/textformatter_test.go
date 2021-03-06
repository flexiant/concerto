package format

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/flexiant/concerto/api/admin"
	"github.com/flexiant/concerto/api/blueprint"
	"github.com/flexiant/concerto/api/dns"
	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestPrintItemDomainTXT(t *testing.T) {

	assert := assert.New(t)
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range *domainsIn {

		domainOut := dns.GetDomainMocked(t, &domainIn)

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

func TestPrintItemTemplateTXT(t *testing.T) {

	assert := assert.New(t)
	templatesIn := testdata.GetTemplateData()
	for _, templateIn := range *templatesIn {

		templateOut := blueprint.GetTemplateMocked(t, &templateIn)

		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		InitializeFormatter("text", mockOut)
		f := GetFormatter()
		assert.NotNil(f, "Formatter")

		err := f.PrintItem(*templateOut)
		assert.Nil(err, "Text formatter PrintItem error")
		mockOut.Flush()

		// TODO add more accurate parsing
		assert.Regexp("^ID:\\ *.*\n*.\n", b.String(), "Text output didn't match regular expression")

	}
}

func TestPrintListDomainsTXT(t *testing.T) {

	assert := assert.New(t)
	domainsIn := testdata.GetDomainData()
	domainOut := dns.GetDomainListMocked(t, domainsIn)

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

func TestPrintListTemplateTXT(t *testing.T) {

	assert := assert.New(t)
	templatesIn := testdata.GetTemplateData()
	templatesOut := blueprint.GetTemplateListMocked(t, templatesIn)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(*templatesOut)
	assert.Nil(err, "Text formatter PrintItem error")
	mockOut.Flush()

	// TODO add more accurate parsing
	assert.Regexp(fmt.Sprintf("^ID.*\n%s.*\n.*", (*templatesOut)[0].ID), b.String(), "Text output didn't match regular expression")
}

func TestPrintListTemplateScriptsTXT(t *testing.T) {

	assert := assert.New(t)
	tScriptsIn := testdata.GetTemplateScriptData()

	for _, tsIn := range *tScriptsIn {
		tScriptsOut := blueprint.GetTemplateScriptListMocked(t, tScriptsIn, tsIn.ID, tsIn.Type)

		var b bytes.Buffer
		mockOut := bufio.NewWriter(&b)
		InitializeFormatter("text", mockOut)
		f := GetFormatter()
		assert.NotNil(f, "Formatter")

		err := f.PrintList(*tScriptsOut)
		assert.Nil(err, "Text formatter PrintItem error")
		mockOut.Flush()

		// TODO add more accurate parsing
		assert.Regexp(fmt.Sprintf("^ID.*\n%s.*\n.*", (*tScriptsOut)[0].ID), b.String(), "Text output didn't match regular expression")

	}
}

func TestPrintListNonSliceErrorTXT(t *testing.T) {

	assert := assert.New(t)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList("string")
	assert.Error(err, "A 'non slice' error should have arosen")
	mockOut.Flush()

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

func TestPrintListReportsTXT(t *testing.T) {
	assert := assert.New(t)
	AdminReportsIn := testdata.GetAdminReportsData()
	AdminReportsOut := admin.GetAdminReportListMocked(t, AdminReportsIn)

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	InitializeFormatter("text", mockOut)
	f := GetFormatter()
	assert.NotNil(f, "Formatter")

	err := f.PrintList(*AdminReportsOut)
	assert.Nil(err, "Text formatter PrintItem error")
	mockOut.Flush()

	assert.Regexp(fmt.Sprintf("^REPORT ID.*\n%s.*\n.*", (*AdminReportsOut)[0].ID), b.String(), "Text output didn't match regular expression")

}
