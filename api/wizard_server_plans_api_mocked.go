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

// GetWizServerPlanListFailErrMocked test mocked function
func GetWizServerPlanListFailErrMocked(t *testing.T, serverPlansIn *[]types.ServerPlan, AppID string, LocID string, ProviderID string) *[]types.ServerPlan {

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
	cs.On("Get", fmt.Sprintf("/v1/wizard/server_plans?app_id=%s&location_id=%s&cloud_provider_id=%s", AppID, LocID, ProviderID)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	serverPlansOut, err := ds.GetWizServerPlanList(AppID, LocID, ProviderID)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverPlansOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &serverPlansOut
}

// GetWizServerPlanListFailStatusMocked test mocked function
func GetWizServerPlanListFailStatusMocked(t *testing.T, serverPlansIn *[]types.ServerPlan, AppID string, LocID string, ProviderID string) *[]types.ServerPlan {

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
	cs.On("Get", fmt.Sprintf("/v1/wizard/server_plans?app_id=%s&location_id=%s&cloud_provider_id=%s", AppID, LocID, ProviderID)).Return(dIn, 499, nil)
	serverPlansOut, err := ds.GetWizServerPlanList(AppID, LocID, ProviderID)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &serverPlansOut
}

// GetWizServerPlanListFailJSONMocked test mocked function
func GetWizServerPlanListFailJSONMocked(t *testing.T, serverPlansIn *[]types.ServerPlan, AppID string, LocID string, ProviderID string) *[]types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewWizServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "WizServerPlan service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/wizard/server_plans?app_id=%s&location_id=%s&cloud_provider_id=%s", AppID, LocID, ProviderID)).Return(dIn, 200, nil)
	serverPlansOut, err := ds.GetWizServerPlanList(AppID, LocID, ProviderID)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &serverPlansOut
}
