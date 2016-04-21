package settings

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// CloudAccountService manages cloudAccount operations
type CloudAccountService struct {
	concertoService utils.ConcertoService
}

// NewCloudAccountService returns a Concerto cloudAccount service
func NewCloudAccountService(concertoService utils.ConcertoService) (*CloudAccountService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &CloudAccountService{
		concertoService: concertoService,
	}, nil
}

// GetCloudAccountList returns the list of cloudAccounts as an array of CloudAccount
func (ca *CloudAccountService) GetCloudAccountList() (cloudAccounts []types.CloudAccount, err error) {
	log.Debug("GetCloudAccountList")

	data, status, err := ca.concertoService.Get("/v1/settings/cloud_accounts")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccounts); err != nil {
		return nil, err
	}

	return cloudAccounts, nil
}

// CreateCloudAccount creates a cloudAccount
func (ca *CloudAccountService) CreateCloudAccount(cloudAccountVector *map[string]interface{}) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("CreateCloudAccount")

	data, status, err := ca.concertoService.Post("/v1/settings/cloud_accounts/", cloudAccountVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccount); err != nil {
		return nil, err
	}

	return cloudAccount, nil
}

// UpdateCloudAccount updates a cloudAccount by its ID
func (ca *CloudAccountService) UpdateCloudAccount(cloudAccountVector *map[string]interface{}, ID string) (cloudAccount *types.CloudAccount, err error) {
	log.Debug("UpdateCloudAccount")

	data, status, err := ca.concertoService.Put(fmt.Sprintf("/v1/settings/cloud_accounts/%s", ID), cloudAccountVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cloudAccount); err != nil {
		return nil, err
	}

	return cloudAccount, nil
}

// DeleteCloudAccount deletes a cloudAccount by its ID
func (ca *CloudAccountService) DeleteCloudAccount(ID string) (err error) {
	log.Debug("DeleteCloudAccount")

	data, status, err := ca.concertoService.Delete(fmt.Sprintf("/v1/settings/cloud_accounts/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
