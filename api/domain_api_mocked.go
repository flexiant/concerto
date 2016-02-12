package api

import (
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO exclude from release compile

// GetDomainListMocked test mocked function
func GetDomainListMocked(t *testing.T, domainsIn *[]types.Domain) *[]types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domainsIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Get", "/v1/dns/domains").Return(dIn, 200, nil)
	domainsOut, err := ds.GetDomainList()
	assert.Nil(err, "Error getting domain list")
	assert.Equal(*domainsIn, domainsOut, "GetDomainList returned different domains")

	return &domainsOut
}

// GetDomainMocked test mocked function
func GetDomainMocked(t *testing.T, domain *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domain)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/dns/domains/%s", domain.ID)).Return(dIn, 200, nil)
	domainOut, err := ds.GetDomain(domain.ID)
	assert.Nil(err, "Error getting domain")
	assert.Equal(*domain, *domainOut, "GetDomain returned different domains")

	return domainOut
}

// CreateDomainMocked test mocked function
func CreateDomainMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// to json
	dOut, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Post", "/v1/dns/domains/", mapIn).Return(dOut, 200, nil)
	domainOut, err := ds.CreateDomain(mapIn)
	assert.Nil(err, "Error creating domain list")
	assert.Equal(domainIn, domainOut, "CreateDomain returned different domains")

	return domainOut
}

// UpdateDomainMocked test mocked function
func UpdateDomainMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// to json
	dOut, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/dns/domains/%s", domainIn.ID), mapIn).Return(dOut, 200, nil)
	domainOut, err := ds.UpdateDomain(mapIn, domainIn.ID)
	assert.Nil(err, "Error updating domain list")
	assert.Equal(domainIn, domainOut, "UpdateDomain returned different domains")

	return domainOut
}

// DeleteDomainMocked test mocked function
func DeleteDomainMocked(t *testing.T, domainIn *types.Domain) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	dIn, err := json.Marshal(domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/dns/domains/%s", domainIn.ID)).Return(dIn, 200, nil)
	err = ds.DeleteDomain(domainIn.ID)
	assert.Nil(err, "Error deleting domain")

}

// GetDomainRecordListMocked test mocked function
func GetDomainRecordListMocked(t *testing.T, domainRecordsIn *[]types.DomainRecord, domainID string) *[]types.DomainRecord {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	drsIn, err := json.Marshal(domainRecordsIn)
	assert.Nil(err, "Domain record test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/dns/domains/%s/records", domainID)).Return(drsIn, 200, nil)
	drsOut, err := ds.GetDomainRecordList(domainID)
	assert.Nil(err, "Error getting domain list")
	assert.Equal(*domainRecordsIn, *drsOut, "GetDomainList returned different domains")

	return drsOut
}

// GetDomainRecordMocked test mocked function
func GetDomainRecordMocked(t *testing.T, dr *types.DomainRecord) *types.DomainRecord {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// to json
	drIn, err := json.Marshal(dr)
	assert.Nil(err, "Domain record test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/dns/domains/%s/records/%s", dr.DomainID, dr.ID)).Return(drIn, 200, nil)
	drOut, err := ds.GetDomainRecord(dr.DomainID, dr.ID)
	assert.Nil(err, "Error getting domain")
	assert.Equal(*dr, *drOut, "GetDomainRecord returned different domain records")

	return drOut

}
