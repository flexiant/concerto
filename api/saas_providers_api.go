package api

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// SaasProviderService manages saasProvider operations
type SaasProviderService struct {
	concertoService utils.ConcertoService
}

// NewSaasProviderService returns a Concerto saasProvider service
func NewSaasProviderService(concertoService utils.ConcertoService) (*SaasProviderService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &SaasProviderService{
		concertoService: concertoService,
	}, nil
}

// GetSaasProviderList returns the list of saasProviders as an array of SaasProvider
func (cl *SaasProviderService) GetSaasProviderList() (saasProviders []types.SaasProvider, err error) {
	log.Debug("GetSaasProviderList")

	data, status, err := cl.concertoService.Get("/v1/cloud/saas_providers")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &saasProviders); err != nil {
		return nil, err
	}

	return saasProviders, nil
}
