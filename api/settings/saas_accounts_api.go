package settings

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// SaasAccountService manages saasAccount operations
type SaasAccountService struct {
	concertoService utils.ConcertoService
}

// NewSaasAccountService returns a Concerto saasAccount service
func NewSaasAccountService(concertoService utils.ConcertoService) (*SaasAccountService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &SaasAccountService{
		concertoService: concertoService,
	}, nil
}

// GetSaasAccountList returns the list of saasAccounts as an array of SaasAccount
func (dm *SaasAccountService) GetSaasAccountList() (saasAccounts []types.SaasAccount, err error) {
	log.Debug("GetSaasAccountList")

	data, status, err := dm.concertoService.Get("/v1/settings/saas_accounts")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &saasAccounts); err != nil {
		return nil, err
	}

	return saasAccounts, nil
}

// CreateSaasAccount creates a saasAccount
func (dm *SaasAccountService) CreateSaasAccount(saasAccountVector *map[string]interface{}) (saasAccount *types.SaasAccount, err error) {
	log.Debug("CreateSaasAccount")

	data, status, err := dm.concertoService.Post("/v1/settings/saas_accounts/", saasAccountVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &saasAccount); err != nil {
		return nil, err
	}

	return saasAccount, nil
}

// UpdateSaasAccount updates a saasAccount by its ID
func (dm *SaasAccountService) UpdateSaasAccount(saasAccountVector *map[string]interface{}, ID string) (saasAccount *types.SaasAccount, err error) {
	log.Debug("UpdateSaasAccount")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/v1/settings/saas_accounts/%s", ID), saasAccountVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &saasAccount); err != nil {
		return nil, err
	}

	return saasAccount, nil
}

// DeleteSaasAccount deletes a saasAccount by its ID
func (dm *SaasAccountService) DeleteSaasAccount(ID string) (err error) {
	log.Debug("DeleteSaasAccount")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/v1/settings/saas_accounts/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
