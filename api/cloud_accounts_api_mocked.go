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

// GetCloudAccountListMocked test mocked function
func GetCloudAccountListMocked(t *testing.T, cloudAccountsIn *[]types.CloudAccount) *[]types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Get", "/v1/settings/cloud_accounts").Return(dIn, 200, nil)
	cloudAccountsOut, err := clAccService.GetCloudAccountList()
	assert.Nil(err, "Error getting cloudAccount list")
	assert.Equal(*cloudAccountsIn, cloudAccountsOut, "GetCloudAccountList returned different cloudAccounts")

	return &cloudAccountsOut
}

// GetCloudAccountListFailErrMocked test mocked function
func GetCloudAccountListFailErrMocked(t *testing.T, cloudAccountsIn *[]types.CloudAccount) *[]types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Get", "/v1/settings/cloud_accounts").Return(dIn, 200, fmt.Errorf("Mocked error"))
	cloudAccountsOut, err := clAccService.GetCloudAccountList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &cloudAccountsOut
}

// GetCloudAccountListFailStatusMocked test mocked function
func GetCloudAccountListFailStatusMocked(t *testing.T, cloudAccountsIn *[]types.CloudAccount) *[]types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountsIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Get", "/v1/settings/cloud_accounts").Return(dIn, 499, nil)
	cloudAccountsOut, err := clAccService.GetCloudAccountList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &cloudAccountsOut
}

// GetCloudAccountListFailJSONMocked test mocked function
func GetCloudAccountListFailJSONMocked(t *testing.T, cloudAccountsIn *[]types.CloudAccount) *[]types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/settings/cloud_accounts").Return(dIn, 200, nil)
	cloudAccountsOut, err := clAccService.GetCloudAccountList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &cloudAccountsOut
}

// CreateCloudAccountMocked test mocked function
func CreateCloudAccountMocked(t *testing.T, cloudAccountIn *types.CloudAccount) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Post", "/v1/settings/cloud_accounts/", mapIn).Return(dOut, 200, nil)
	cloudAccountOut, err := clAccService.CreateCloudAccount(mapIn)
	assert.Nil(err, "Error creating cloudAccount list")
	assert.Equal(cloudAccountIn, cloudAccountOut, "CreateCloudAccount returned different cloudAccounts")

	return cloudAccountOut
}

// CreateCloudAccountFailErrMocked test mocked function
func CreateCloudAccountFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Post", "/v1/settings/cloud_accounts/", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	cloudAccountOut, err := clAccService.CreateCloudAccount(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return cloudAccountOut
}

// CreateCloudAccountFailStatusMocked test mocked function
func CreateCloudAccountFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Post", "/v1/settings/cloud_accounts/", mapIn).Return(dOut, 499, nil)
	cloudAccountOut, err := clAccService.CreateCloudAccount(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return cloudAccountOut
}

// CreateCloudAccountFailJSONMocked test mocked function
func CreateCloudAccountFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/settings/cloud_accounts/", mapIn).Return(dIn, 200, nil)
	cloudAccountOut, err := clAccService.CreateCloudAccount(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}

// UpdateCloudAccountMocked test mocked function
func UpdateCloudAccountMocked(t *testing.T, cloudAccountIn *types.CloudAccount) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/settings/cloud_accounts/%s", cloudAccountIn.Id), mapIn).Return(dOut, 200, nil)
	cloudAccountOut, err := clAccService.UpdateCloudAccount(mapIn, cloudAccountIn.Id)
	assert.Nil(err, "Error updating cloudAccount list")
	assert.Equal(cloudAccountIn, cloudAccountOut, "UpdateCloudAccount returned different cloudAccounts")

	return cloudAccountOut
}

// UpdateCloudAccountFailErrMocked test mocked function
func UpdateCloudAccountFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/settings/cloud_accounts/%s", cloudAccountIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	cloudAccountOut, err := clAccService.UpdateCloudAccount(mapIn, cloudAccountIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return cloudAccountOut
}

// UpdateCloudAccountFailStatusMocked test mocked function
func UpdateCloudAccountFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// to json
	dOut, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/settings/cloud_accounts/%s", cloudAccountIn.Id), mapIn).Return(dOut, 499, nil)
	cloudAccountOut, err := clAccService.UpdateCloudAccount(mapIn, cloudAccountIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return cloudAccountOut
}

// UpdateCloudAccountFailJSONMocked test mocked function
func UpdateCloudAccountFailJSONMocked(t *testing.T, cloudAccountIn *types.CloudAccount) *types.CloudAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/settings/cloud_accounts/%s", cloudAccountIn.Id), mapIn).Return(dIn, 200, nil)
	cloudAccountOut, err := clAccService.UpdateCloudAccount(mapIn, cloudAccountIn.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(cloudAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return cloudAccountOut
}

// DeleteCloudAccountMocked test mocked function
func DeleteCloudAccountMocked(t *testing.T, cloudAccountIn *types.CloudAccount) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/settings/cloud_accounts/%s", cloudAccountIn.Id)).Return(dIn, 200, nil)
	err = clAccService.DeleteCloudAccount(cloudAccountIn.Id)
	assert.Nil(err, "Error deleting cloudAccount")
}

// DeleteCloudAccountFailErrMocked test mocked function
func DeleteCloudAccountFailErrMocked(t *testing.T, cloudAccountIn *types.CloudAccount) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/settings/cloud_accounts/%s", cloudAccountIn.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	err = clAccService.DeleteCloudAccount(cloudAccountIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// DeleteCloudAccountFailStatusMocked test mocked function
func DeleteCloudAccountFailStatusMocked(t *testing.T, cloudAccountIn *types.CloudAccount) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	clAccService, err := NewCloudAccountService(cs)
	assert.Nil(err, "Couldn't load cloudAccount service")
	assert.NotNil(clAccService, "CloudAccount service not instanced")

	// to json
	dIn, err := json.Marshal(cloudAccountIn)
	assert.Nil(err, "CloudAccount test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/settings/cloud_accounts/%s", cloudAccountIn.Id)).Return(dIn, 499, nil)
	err = clAccService.DeleteCloudAccount(cloudAccountIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
