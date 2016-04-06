package api

import (
	"testing"

	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewDomainServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewDomainService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetDomainList(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	GetDomainListMocked(t, domainsIn)
	GetDomainListFailErrMocked(t, domainsIn)
	GetDomainListFailStatusMocked(t, domainsIn)
	GetDomainListFailJSONMocked(t, domainsIn)
}

func TestGetDomain(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range *domainsIn {
		GetDomainMocked(t, &domainIn)
		GetDomainFailErrMocked(t, &domainIn)
		GetDomainFailStatusMocked(t, &domainIn)
		GetDomainFailJSONMocked(t, &domainIn)
	}
}

func TestCreateDomain(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range *domainsIn {
		CreateDomainMocked(t, &domainIn)
	}
}

func TestUpdateDomain(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range *domainsIn {
		UpdateDomainMocked(t, &domainIn)
	}
}

func TestDeleteDomain(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range *domainsIn {
		DeleteDomainMocked(t, &domainIn)
	}
}

func TestListDomainRecords(t *testing.T) {
	drsIn := testdata.GetDomainRecordData()
	for _, drIn := range *drsIn {
		GetDomainRecordListMocked(t, drsIn, drIn.ID)
	}
}

func TestGetDomainRecord(t *testing.T) {
	drsIn := testdata.GetDomainRecordData()
	for _, drIn := range *drsIn {
		GetDomainRecordMocked(t, &drIn)
	}
}

func TestCreateDomainRecord(t *testing.T) {
	drsIn := testdata.GetDomainRecordData()
	for _, drIn := range *drsIn {
		CreateDomainRecordMocked(t, &drIn)
	}
}

func TestUpdateDomainRecord(t *testing.T) {
	drsIn := testdata.GetDomainRecordData()
	for _, drIn := range *drsIn {
		UpdateDomainRecordMocked(t, &drIn)
	}
}

func TestDeleteDomainRecords(t *testing.T) {
	drsIn := testdata.GetDomainRecordData()
	for _, drIn := range *drsIn {
		DeleteDomainRecordMocked(t, &drIn)
	}
}
