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
func GetServerPlanListMocked(t *testing.T, serverPlansIn *[]types.ServerPlan, cloudProviderId string) *[]types.ServerPlan {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/cloud_providers/%s/server_plans", cloudProviderId)).Return(dIn, 200, nil)
	serverPlansOut, err := ds.GetServerPlanList(cloudProviderId)
	assert.Nil(err, "Error getting serverPlan list")
	assert.Equal(*serverPlansIn, serverPlansOut, "GetServerPlanList returned different serverPlans")

	return &serverPlansOut
}

// GetServerPlanListFailErrMocked test mocked function
func GetServerPlanListFailErrMocked(t *testing.T, serverPlansIn *[]types.ServerPlan, cloudProviderId string) *[]types.ServerPlan {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/cloud_providers/%s/server_plans", cloudProviderId)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	serverPlansOut, err := ds.GetServerPlanList(cloudProviderId)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverPlansOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &serverPlansOut
}

// GetServerPlanListFailStatusMocked test mocked function
func GetServerPlanListFailStatusMocked(t *testing.T, serverPlansIn *[]types.ServerPlan, cloudProviderId string) *[]types.ServerPlan {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/cloud_providers/%s/server_plans", cloudProviderId)).Return(dIn, 499, nil)
	serverPlansOut, err := ds.GetServerPlanList(cloudProviderId)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &serverPlansOut
}

// GetServerPlanListFailJSONMocked test mocked function
func GetServerPlanListFailJSONMocked(t *testing.T, serverPlansIn *[]types.ServerPlan, cloudProviderId string) *[]types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "ServerPlan service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/cloud_providers/%s/server_plans", cloudProviderId)).Return(dIn, 200, nil)
	serverPlansOut, err := ds.GetServerPlanList(cloudProviderId)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverPlansOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

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

// GetServerPlanFailErrMocked test mocked function
func GetServerPlanFailErrMocked(t *testing.T, serverPlan *types.ServerPlan) *types.ServerPlan {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/server_plans/%s", serverPlan.Id)).Return(dIn, 200, fmt.Errorf("Mocked error"))
	serverPlanOut, err := ds.GetServerPlan(serverPlan.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(serverPlanOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return serverPlanOut
}

// GetServerPlanFailStatusMocked test mocked function
func GetServerPlanFailStatusMocked(t *testing.T, serverPlan *types.ServerPlan) *types.ServerPlan {

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
	cs.On("Get", fmt.Sprintf("/v1/cloud/server_plans/%s", serverPlan.Id)).Return(dIn, 499, nil)
	serverPlanOut, err := ds.GetServerPlan(serverPlan.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(serverPlanOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return serverPlanOut
}

// GetServerPlanFailJSONMocked test mocked function
func GetServerPlanFailJSONMocked(t *testing.T, serverPlan *types.ServerPlan) *types.ServerPlan {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	ds, err := NewServerPlanService(cs)
	assert.Nil(err, "Couldn't load serverPlan service")
	assert.NotNil(ds, "ServerPlan service not instanced")

	// wrong json
	dIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/cloud/server_plans/%s", serverPlan.Id)).Return(dIn, 200, nil)
	serverPlanOut, err := ds.GetServerPlan(serverPlan.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(serverPlanOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return serverPlanOut
}
