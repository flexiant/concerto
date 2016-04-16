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

// GetSaasAccountListMocked test mocked function
func GetSaasAccountListMocked(t *testing.T, saasAccountsIn *[]types.SaasAccount) *[]types.SaasAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// to json
	dIn, err := json.Marshal(saasAccountsIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// call service
	cs.On("Get", "/v1/settings/saas_accounts").Return(dIn, 200, nil)
	saasAccountsOut, err := ds.GetSaasAccountList()
	assert.Nil(err, "Error getting saasAccount list")
	assert.Equal(*saasAccountsIn, saasAccountsOut, "GetSaasAccountList returned different saasAccounts")

	return &saasAccountsOut
}

// GetSaasAccountListFailErrMocked test mocked function
func GetSaasAccountListFailErrMocked(t *testing.T, saasAccountsIn *[]types.SaasAccount) *[]types.SaasAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// to json
	dIn, err := json.Marshal(saasAccountsIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// call service
	cs.On("Get", "/v1/settings/saas_accounts").Return(dIn, 200, fmt.Errorf("Mocked error"))
	saasAccountsOut, err := ds.GetSaasAccountList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(saasAccountsOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &saasAccountsOut
}

// GetSaasAccountListFailStatusMocked test mocked function
func GetSaasAccountListFailStatusMocked(t *testing.T, saasAccountsIn *[]types.SaasAccount) *[]types.SaasAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// to json
	dIn, err := json.Marshal(saasAccountsIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// call service
	cs.On("Get", "/v1/settings/saas_accounts").Return(dIn, 499, nil)
	saasAccountsOut, err := ds.GetSaasAccountList()

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(saasAccountsOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &saasAccountsOut
}

// GetSaasAccountListFailJSONMocked test mocked function
func GetSaasAccountListFailJSONMocked(t *testing.T, saasAccountsIn *[]types.SaasAccount) *[]types.SaasAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/settings/saas_accounts").Return(dIn, 200, nil)
	saasAccountsOut, err := ds.GetSaasAccountList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(saasAccountsOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &saasAccountsOut
}

// CreateSaasAccountMocked test mocked function
func CreateSaasAccountMocked(t *testing.T, saasAccountIn *types.SaasAccount) *types.SaasAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// to json
	dOut, err := json.Marshal(saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// call service
	cs.On("Post", "/v1/settings/saas_accounts/", mapIn).Return(dOut, 200, nil)
	saasAccountOut, err := ds.CreateSaasAccount(mapIn)
	assert.Nil(err, "Error creating saasAccount list")
	assert.Equal(saasAccountIn, saasAccountOut, "CreateSaasAccount returned different saasAccounts")

	return saasAccountOut
}

// CreateSaasAccountFailErrMocked test mocked function
func CreateSaasAccountFailErrMocked(t *testing.T, saasAccountIn *types.SaasAccount) *types.SaasAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// to json
	dOut, err := json.Marshal(saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// call service
	cs.On("Post", "/v1/settings/saas_accounts/", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	saasAccountOut, err := ds.CreateSaasAccount(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(saasAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return saasAccountOut
}

// CreateSaasAccountFailStatusMocked test mocked function
func CreateSaasAccountFailStatusMocked(t *testing.T, saasAccountIn *types.SaasAccount) *types.SaasAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// to json
	dOut, err := json.Marshal(saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// call service
	cs.On("Post", "/v1/settings/saas_accounts/", mapIn).Return(dOut, 499, nil)
	saasAccountOut, err := ds.CreateSaasAccount(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(saasAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return saasAccountOut
}

// CreateSaasAccountFailJSONMocked test mocked function
func CreateSaasAccountFailJSONMocked(t *testing.T, saasAccountIn *types.SaasAccount) *types.SaasAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/settings/saas_accounts/", mapIn).Return(dIn, 200, nil)
	saasAccountOut, err := ds.CreateSaasAccount(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(saasAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return saasAccountOut
}

// UpdateSaasAccountMocked test mocked function
func UpdateSaasAccountMocked(t *testing.T, saasAccountIn *types.SaasAccount) *types.SaasAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// to json
	dOut, err := json.Marshal(saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/settings/saas_accounts/%s", saasAccountIn.Id), mapIn).Return(dOut, 200, nil)
	saasAccountOut, err := ds.UpdateSaasAccount(mapIn, saasAccountIn.Id)
	assert.Nil(err, "Error updating saasAccount list")
	assert.Equal(saasAccountIn, saasAccountOut, "UpdateSaasAccount returned different saasAccounts")

	return saasAccountOut
}

// UpdateSaasAccountFailErrMocked test mocked function
func UpdateSaasAccountFailErrMocked(t *testing.T, saasAccountIn *types.SaasAccount) *types.SaasAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// to json
	dOut, err := json.Marshal(saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/settings/saas_accounts/%s", saasAccountIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	saasAccountOut, err := ds.UpdateSaasAccount(mapIn, saasAccountIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(saasAccountOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return saasAccountOut
}

// UpdateSaasAccountFailStatusMocked test mocked function
func UpdateSaasAccountFailStatusMocked(t *testing.T, saasAccountIn *types.SaasAccount) *types.SaasAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// to json
	dOut, err := json.Marshal(saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/settings/saas_accounts/%s", saasAccountIn.Id), mapIn).Return(dOut, 499, nil)
	saasAccountOut, err := ds.UpdateSaasAccount(mapIn, saasAccountIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(saasAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return saasAccountOut
}

// UpdateSaasAccountFailJSONMocked test mocked function
func UpdateSaasAccountFailJSONMocked(t *testing.T, saasAccountIn *types.SaasAccount) *types.SaasAccount {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/settings/saas_accounts/%s", saasAccountIn.Id), mapIn).Return(dIn, 200, nil)
	saasAccountOut, err := ds.UpdateSaasAccount(mapIn, saasAccountIn.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(saasAccountOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return saasAccountOut
}

// DeleteSaasAccountMocked test mocked function
func DeleteSaasAccountMocked(t *testing.T, saasAccountIn *types.SaasAccount) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// to json
	dIn, err := json.Marshal(saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/settings/saas_accounts/%s", saasAccountIn.Id)).Return(dIn, 200, nil)
	err = ds.DeleteSaasAccount(saasAccountIn.Id)
	assert.Nil(err, "Error deleting saasAccount")
}

// DeleteSaasAccountFailErrMocked test mocked function
func DeleteSaasAccountFailErrMocked(t *testing.T, saasAccountIn *types.SaasAccount) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// to json
	dIn, err := json.Marshal(saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/settings/saas_accounts/%s", saasAccountIn.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	err = ds.DeleteSaasAccount(saasAccountIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// DeleteSaasAccountFailStatusMocked test mocked function
func DeleteSaasAccountFailStatusMocked(t *testing.T, saasAccountIn *types.SaasAccount) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewSaasAccountService(cs)
	assert.Nil(err, "Couldn't load saasAccount service")
	assert.NotNil(ds, "SaasAccount service not instanced")

	// to json
	dIn, err := json.Marshal(saasAccountIn)
	assert.Nil(err, "SaasAccount test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/settings/saas_accounts/%s", saasAccountIn.Id)).Return(dIn, 499, nil)
	err = ds.DeleteSaasAccount(saasAccountIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}
