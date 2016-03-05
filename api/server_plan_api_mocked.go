package api

import (
	"encoding/json"
	"fmt"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO exclude from release compile

// GetServerPlanListMocked test mocked function
func GetServerPlanListMocked(t *testing.T, serverPlansIn *[]types.ServerPlan, cloudProviderID string) *[]types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "ServerPlan service not instanced")

	// to json
	dIn, err := json.Marshal(serverPlansIn)
	assert.Nil(err, "ServerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/cloud_providers/%s/server_plans", cloudProviderID)).Return(dIn, 200, nil)
	serverPlansOut, err := ds.GetServerPlanList(cloudProviderID)
	assert.Nil(err, "Error getting serverPlan list")
	assert.Equal(*serverPlansIn, serverPlansOut, "GetServerPlanList returned different serverPlans")

	return &serverPlansOut
}

// GetServerPlanMocked test mocked function
func GetServerPlanMocked(t *testing.T, serverPlan *types.ServerPlan) *types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "ServerPlan service not instanced")

	// to json
	dIn, err := json.Marshal(serverPlan)
	assert.Nil(err, "ServerPlan test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/server_plans/%s", serverPlan.Id)).Return(dIn, 200, nil)
	serverPlanOut, err := ds.GetServerPlan(serverPlan.Id)
	assert.Nil(err, "Error getting serverPlan")
	assert.Equal(*serverPlan, *serverPlanOut, "GetServerPlan returned different serverPlans")

	return serverPlanOut
}
