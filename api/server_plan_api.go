package api

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// ServerPlanService manages serverPlan operations
type ServerPlanService struct {
	concertoService utils.ConcertoService
}

// NewServerPlanService returns a Concerto serverPlan service
func NewServerPlanService(concertoService utils.ConcertoService) (*ServerPlanService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &ServerPlanService{
		concertoService: concertoService,
	}, nil
}

// GetServerPlanList returns the list of serverPlans as an array of ServerPlan
func (dm *ServerPlanService) GetServerPlanList(ProviderID string) (serverPlans []types.ServerPlan, err error) {
	log.Debug("GetServerPlanList")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/v1/cloud/cloud_providers/%s/server_plans", ProviderID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverPlans); err != nil {
		return nil, err
	}

	return serverPlans, nil
}

// GetServerPlan returns a serverPlan by its ID
func (dm *ServerPlanService) GetServerPlan(ID string) (serverPlan *types.ServerPlan, err error) {
	log.Debug("GetServerPlan")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/v1/cloud/server_plans/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &serverPlan); err != nil {
		return nil, err
	}

	return serverPlan, nil
}
