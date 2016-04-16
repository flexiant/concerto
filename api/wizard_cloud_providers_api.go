package api

import (
	"encoding/json"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// WizCloudProvidersService manages wizCloudProviders operations
type WizCloudProvidersService struct {
	concertoService utils.ConcertoService
}

// NewWizCloudProvidersService returns a Concerto wizCloudProviders service
func NewWizCloudProvidersService(concertoService utils.ConcertoService) (*WizCloudProvidersService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &WizCloudProvidersService{
		concertoService: concertoService,
	}, nil
}

// GetWizCloudProvidersList returns the list of wizCloudProviderss as an array of CloudProvider
func (dm *WizCloudProvidersService) GetWizCloudProviderList(AppID string, LocID string) (wizCloudProviderss []types.CloudProvider, err error) {
	log.Debug("GetWizCloudProvidersList")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/v1/wizard/cloud_providers?app_id=%s&location_id=%s", AppID, LocID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &wizCloudProviderss); err != nil {
		return nil, err
	}

	return wizCloudProviderss, nil
}
