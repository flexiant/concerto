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
