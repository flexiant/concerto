package api

import (
	"testing"

	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewSaasAccountServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewSaasAccountService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetSaasAccountList(t *testing.T) {
	saasAccountsIn := testdata.GetSaasAccountData()
	GetSaasAccountListMocked(t, saasAccountsIn)
	GetSaasAccountListFailErrMocked(t, saasAccountsIn)
	GetSaasAccountListFailStatusMocked(t, saasAccountsIn)
	GetSaasAccountListFailJSONMocked(t, saasAccountsIn)
}

func TestCreateSaasAccount(t *testing.T) {
	saasAccountsIn := testdata.GetSaasAccountData()
	for _, saasAccountIn := range *saasAccountsIn {
		CreateSaasAccountMocked(t, &saasAccountIn)
		CreateSaasAccountFailErrMocked(t, &saasAccountIn)
		CreateSaasAccountFailStatusMocked(t, &saasAccountIn)
		CreateSaasAccountFailJSONMocked(t, &saasAccountIn)
	}
}

func TestUpdateSaasAccount(t *testing.T) {
	saasAccountsIn := testdata.GetSaasAccountData()
	for _, saasAccountIn := range *saasAccountsIn {
		UpdateSaasAccountMocked(t, &saasAccountIn)
		UpdateSaasAccountFailErrMocked(t, &saasAccountIn)
		UpdateSaasAccountFailStatusMocked(t, &saasAccountIn)
		UpdateSaasAccountFailJSONMocked(t, &saasAccountIn)
	}
}

func TestDeleteSaasAccount(t *testing.T) {
	saasAccountsIn := testdata.GetSaasAccountData()
	for _, saasAccountIn := range *saasAccountsIn {
		DeleteSaasAccountMocked(t, &saasAccountIn)
		DeleteSaasAccountFailErrMocked(t, &saasAccountIn)
		DeleteSaasAccountFailStatusMocked(t, &saasAccountIn)
	}
}
