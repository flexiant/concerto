package api

// package api
//
// import (
// 	"encoding/json"
// 	"github.com/flexiant/concerto/api/types"
// 	"github.com/flexiant/concerto/testdata"
// 	"github.com/flexiant/concerto/utils"
// 	"github.com/stretchr/testify/assert"
// 	"testing"
// )
//
// func TestGetDomainList(t *testing.T) {
//
// 	assert := assert.New(t)
//
// 	domainsTest, err := testdata.GetDomainData()
// 	assert.Nil(err, "Couldn't load domain test data")
//
// 	var domainsIn []types.Domain
// 	for _, domainTest := range domainsTest {
// 		if domainTest.FieldsOK {
// 			domainsIn = append(domainsIn, domainTest.Domain)
// 		}
// 	}
//
// 	getDomainList(t, &domainsIn)
//
// }
//
// func getDomainList(t *testing.T, domainsIn *[]types.Domain) {
// 	assert := assert.New(t)
//
// 	// wire up
// 	cs := &utils.MockConcertoService{}
// 	ds, err := NewDomainService(cs)
// 	assert.Nil(err, "Couldn't load domain service")
// 	assert.NotNil(ds, "Domain service not instanced")
//
// 	// to json
// 	dIn, err := json.Marshal(domainsIn)
// 	assert.Nil(err, "Domain test data corrupted")
//
// 	// call service
// 	cs.On("Get", "/v1/dns/domains").Return(dIn, 200, nil)
// 	domainsOut, err := ds.GetDomainList()
// 	assert.Nil(err, "Error getting domain list")
// 	assert.Equal(*domainsIn, domainsOut, "GetDomainList returned different domains")
// }
