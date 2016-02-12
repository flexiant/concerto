package api

import (
	"github.com/flexiant/concerto/testdata"

	"testing"
)

func TestGetDomainList(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	GetDomainListMocked(t, domainsIn)
}

func TestGetDomain(t *testing.T) {
	domainsIn := testdata.GetDomainData()
	for _, domainIn := range *domainsIn {
		GetDomainMocked(t, &domainIn)
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
