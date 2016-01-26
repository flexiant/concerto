package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestDomainList subcommand function
func TestDomainList(t *testing.T) {

	// TODO load from file
	domainsIn := []api.Domain{
		{
			ID:      "fakeID",
			Name:    "fakeName",
			TTL:     1000,
			Contact: "fakeContact",
			Minimum: 10,
			Enabled: true,
		},
	}

	assert := assert.New(t)

	cs := &utils.MockConcertoService{}
	ds, err := api.NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	dIn, err := json.Marshal(domainsIn)
	assert.Nil(err, "Domain test data corrupted")

	cs.On("Get", "/v1/dns/domains").Return(dIn, 200, nil)
	domainsOut, err := ds.GetDomainList()
	assert.Nil(err, "Error getting domain list")
	assert.Equal(domainsIn, domainsOut, "GetDomainList returned different domains")

	var b bytes.Buffer
	mockOut := bufio.NewWriter(&b)
	f := format.NewJSONFormatter(mockOut)
	assert.NotNil(f, "Formatter")

	err = f.PrintList(domainsOut)
	assert.Nil(err, "JSON Formatter Printlinst error")

	mockOut.Flush()
	assert.Equal(`[{"id":"fakeID","name":"fakeName","ttl":1000,"contact":"fakeContact","minimum":10,"enabled":true}]
`, b.String(), "Wrong JSON output")
}
