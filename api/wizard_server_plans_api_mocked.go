package api

import (
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// GetWizServerPlanListMocked test mocked function
func GetWizServerPlanListMocked(t *testing.T, serverPlansIn *[]types.ServerPlan, AppID string, LocID string, ProviderID string) *[]types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWizServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "WizServerPlan service not instanced")

	// to json
	dIn, err := json.Marshal(serverPlansIn)
	assert.Nil(err, "WizServerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/wizard/server_plans?app_id=%s&location_id=%s&cloud_provider_id=%s", AppID, LocID, ProviderID)).Return(dIn, 200, nil)
	serverPlansOut, err := ds.GetWizServerPlanList(AppID, LocID, ProviderID)
	assert.Nil(err, "Error getting serverPlan list")
	assert.Equal(*serverPlansIn, serverPlansOut, "GetWizServerPlanList returned different serverPlans")

	return &serverPlansOut
}
