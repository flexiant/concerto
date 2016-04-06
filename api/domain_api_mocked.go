package api

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
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

// GetDomainListFailErrMocked test mocked function
func GetDomainListFailErrMocked(t *testing.T, domainsIn *[]types.Domain) *[]types.Domain {

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
	cs.On("Get", "/v1/dns/domains").Return(dIn, 200, fmt.Errorf("Mocked error"))
	domainsOut, err := ds.GetDomainList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(domainsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &domainsOut
}

// GetDomainListFailStatusMocked test mocked function
func GetDomainListFailStatusMocked(t *testing.T, domainsIn *[]types.Domain) *[]types.Domain {

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
	cs.On("Get", "/v1/dns/domains").Return(dIn, 499, nil)
	domainsOut, err := ds.GetDomainList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(domainsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &domainsOut
}

// GetDomainListFailJSONMocked test mocked function
func GetDomainListFailJSONMocked(t *testing.T, domainsIn *[]types.Domain) *[]types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/dns/domains").Return(dIn, 200, nil)
	domainsOut, err := ds.GetDomainList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(domainsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// GetDomainFailErrMocked test mocked function
func GetDomainFailErrMocked(t *testing.T, domain *types.Domain) *types.Domain {

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
	cs.On("Get", fmt.Sprintf("/v1/dns/domains/%s", domain.ID)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	domainOut, err := ds.GetDomain(domain.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return domainOut
}

// GetDomainFailStatusMocked test mocked function
func GetDomainFailStatusMocked(t *testing.T, domain *types.Domain) *types.Domain {

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
	cs.On("Get", fmt.Sprintf("/v1/dns/domains/%s", domain.ID)).Return(dIn, 499, nil)
	domainOut, err := ds.GetDomain(domain.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return domainOut
}

// GetDomainFailJSONMocked test mocked function
func GetDomainFailJSONMocked(t *testing.T, domain *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/dns/domains/%s", domain.ID)).Return(dIn, 200, nil)
	domainOut, err := ds.GetDomain(domain.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// CreateDomainFailErrMocked test mocked function
func CreateDomainFailErrMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

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
	cs.On("Post", "/v1/dns/domains/", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	domainOut, err := ds.CreateDomain(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return domainOut
}

// CreateDomainFailStatusMocked test mocked function
func CreateDomainFailStatusMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

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
	cs.On("Post", "/v1/dns/domains/", mapIn).Return(dOut, 499, nil)
	domainOut, err := ds.CreateDomain(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return domainOut
}

// CreateDomainFailJSONMocked test mocked function
func CreateDomainFailJSONMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/dns/domains/", mapIn).Return(dIn, 200, nil)
	domainOut, err := ds.CreateDomain(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// UpdateDomainFailErrMocked test mocked function
func UpdateDomainFailErrMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

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
	cs.On("Put", fmt.Sprintf("/v1/dns/domains/%s", domainIn.ID), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	domainOut, err := ds.UpdateDomain(mapIn, domainIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return domainOut
}

// UpdateDomainFailStatusMocked test mocked function
func UpdateDomainFailStatusMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

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
	cs.On("Put", fmt.Sprintf("/v1/dns/domains/%s", domainIn.ID), mapIn).Return(dOut, 499, nil)
	domainOut, err := ds.UpdateDomain(mapIn, domainIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return domainOut
}

// UpdateDomainFailJSONMocked test mocked function
func UpdateDomainFailJSONMocked(t *testing.T, domainIn *types.Domain) *types.Domain {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*domainIn)
	assert.Nil(err, "Domain test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/dns/domains/%s", domainIn.ID), mapIn).Return(dIn, 200, nil)
	domainOut, err := ds.UpdateDomain(mapIn, domainIn.ID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(domainOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// DeleteDomainFailErrMocked test mocked function
func DeleteDomainFailErrMocked(t *testing.T, domainIn *types.Domain) {

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
	cs.On("Delete", fmt.Sprintf("/v1/dns/domains/%s", domainIn.ID)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	err = ds.DeleteDomain(domainIn.ID)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// DeleteDomainFailStatusMocked test mocked function
func DeleteDomainFailStatusMocked(t *testing.T, domainIn *types.Domain) {

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
	cs.On("Delete", fmt.Sprintf("/v1/dns/domains/%s", domainIn.ID)).Return(dIn, 499, nil)
	err = ds.DeleteDomain(domainIn.ID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
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

// CreateDomainRecordMocked test mocked function
func CreateDomainRecordMocked(t *testing.T, dr *types.DomainRecord) *types.DomainRecord {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*dr)
	assert.Nil(err, "Domain record test data corrupted")

	// to json
	drIn, err := json.Marshal(dr)
	assert.Nil(err, "Domain record test data corrupted")

	// call service
	cs.On("Post", fmt.Sprintf("/v1/dns/domains/%s/records", dr.DomainID), mapIn).Return(drIn, 200, nil)
	drOut, err := ds.CreateDomainRecord(mapIn, dr.DomainID)
	assert.Nil(err, "Error getting domain")
	assert.Equal(*dr, *drOut, "CreateDomainRecord returned different domain records")

	return drOut
}

// UpdateDomainRecordMocked test mocked function
func UpdateDomainRecordMocked(t *testing.T, dr *types.DomainRecord) *types.DomainRecord {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewDomainService(cs)
	assert.Nil(err, "Couldn't load domain service")
	assert.NotNil(ds, "Domain service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*dr)
	assert.Nil(err, "Domain record test data corrupted")

	// to json
	drIn, err := json.Marshal(dr)
	assert.Nil(err, "Domain record test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/dns/domains/%s/records/%s", dr.DomainID, dr.ID), mapIn).Return(drIn, 200, nil)
	drOut, err := ds.UpdateDomainRecord(mapIn, dr.DomainID, dr.ID)
	assert.Nil(err, "Error updating domain list")
	assert.Equal(*dr, *drOut, "UpdateDomainRecord returned different domain records")

	return drOut
}

// DeleteDomainRecordMocked test mocked function
func DeleteDomainRecordMocked(t *testing.T, dr *types.DomainRecord) {

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
	cs.On("Delete", fmt.Sprintf("/v1/dns/domains/%s/records/%s", dr.DomainID, dr.ID)).Return(drIn, 200, nil)
	err = ds.DeleteDomainRecord(dr.DomainID, dr.ID)
	assert.Nil(err, "Error deleting domain record")
}
