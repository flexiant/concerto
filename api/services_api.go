package api

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// ServicesService manages service operations
type ServicesService struct {
	concertoService utils.ConcertoService
}

// NewServicesService returns a Concerto service service
func NewServicesService(concertoService utils.ConcertoService) (*ServicesService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &ServicesService{
		concertoService: concertoService,
	}, nil
}

// GetServiceList returns the list of services as an array of Service
func (ss *ServicesService) GetServiceList() (services []types.Service, err error) {
	log.Debug("GetServiceList")

	data, status, err := ss.concertoService.Get("/v1/blueprint/services")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {

	}

	if err = json.Unmarshal(data, &services); err != nil {
		return nil, err
	}

	return services, nil
}

// GetService returns a service by its ID
func (ss *ServicesService) GetService(ID string) (service *types.Service, err error) {
	log.Debug("GetService")

	data, status, err := ss.concertoService.Get(fmt.Sprintf("/v1/blueprint/services/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &service); err != nil {
		return nil, err
	}

	return service, nil
}
