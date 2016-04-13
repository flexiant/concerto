package api

import (
	"testing"

	"github.com/flexiant/concerto/testdata"
	"github.com/stretchr/testify/assert"
)

func TestNewCloudAccountServiceNil(t *testing.T) {
	assert := assert.New(t)
	rs, err := NewCloudAccountService(nil)
	assert.Nil(rs, "Uninitialized service should return nil")
	assert.NotNil(err, "Uninitialized service should return error")
}

func TestGetCloudAccountList(t *testing.T) {
	cloudAccountsIn := testdata.GetCloudAccountData()
	GetCloudAccountListMocked(t, cloudAccountsIn)
	GetCloudAccountListFailErrMocked(t, cloudAccountsIn)
	GetCloudAccountListFailStatusMocked(t, cloudAccountsIn)
	GetCloudAccountListFailJSONMocked(t, cloudAccountsIn)
}

func TestCreateCloudAccount(t *testing.T) {
	cloudAccountsIn := testdata.GetCloudAccountData()
	for _, cloudAccountIn := range *cloudAccountsIn {
		CreateCloudAccountMocked(t, &cloudAccountIn)
		CreateCloudAccountFailErrMocked(t, &cloudAccountIn)
		CreateCloudAccountFailStatusMocked(t, &cloudAccountIn)
		CreateCloudAccountFailJSONMocked(t, &cloudAccountIn)
	}
}

func TestUpdateCloudAccount(t *testing.T) {
	cloudAccountsIn := testdata.GetCloudAccountData()
	for _, cloudAccountIn := range *cloudAccountsIn {
		UpdateCloudAccountMocked(t, &cloudAccountIn)
		UpdateCloudAccountFailErrMocked(t, &cloudAccountIn)
		UpdateCloudAccountFailStatusMocked(t, &cloudAccountIn)
		UpdateCloudAccountFailJSONMocked(t, &cloudAccountIn)
	}
}

func TestDeleteCloudAccount(t *testing.T) {
	cloudAccountsIn := testdata.GetCloudAccountData()
	for _, cloudAccountIn := range *cloudAccountsIn {
		DeleteCloudAccountMocked(t, &cloudAccountIn)
		DeleteCloudAccountFailErrMocked(t, &cloudAccountIn)
		DeleteCloudAccountFailStatusMocked(t, &cloudAccountIn)
	}
}
