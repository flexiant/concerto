package api

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// WizServerPlanService manages serverPlan operations
type WizServerPlanService struct {
	concertoService utils.ConcertoService
}

// NewWizServerPlanService returns a Concerto serverPlan service
func NewWizServerPlanService(concertoService utils.ConcertoService) (*WizServerPlanService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &WizServerPlanService{
		concertoService: concertoService,
	}, nil
}

// GetWizServerPlanList returns the list of serverPlans as an array of ServerPlan
func (dm *WizServerPlanService) GetWizServerPlanList(AppID string, LocID string, ProviderID string) (serverPlans []types.ServerPlan, err error) {
	log.Debug("GetWizServerPlanList")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/v1/wizard/server_plans?app_id=%s&location_id=%s&cloud_provider_id=%s", AppID, LocID, ProviderID))
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
