package api

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/stretchr/testify/assert"
)

// TODO exclude from release compile

// GetLoadBalancerListMocked test mocked function
func GetLoadBalancerListMocked(t *testing.T, loadBalancersIn *[]types.LoadBalancer) *[]types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// to json
	lbIn, err := json.Marshal(loadBalancersIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Get", "/v1/network/load_balancers").Return(lbIn, 200, nil)
	loadBalancersOut, err := lbs.GetLoadBalancerList()
	assert.Nil(err, "Error getting loadBalancer list")
	assert.Equal(*loadBalancersIn, loadBalancersOut, "GetLoadBalancerList returned different loadBalancers")

	return &loadBalancersOut
}

// GetLoadBalancerListFailErrMocked test mocked function
func GetLoadBalancerListFailErrMocked(t *testing.T, loadBalancersIn *[]types.LoadBalancer) *[]types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// to json
	lbIn, err := json.Marshal(loadBalancersIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Get", "/v1/network/load_balancers").Return(lbIn, 200, fmt.Errorf("Mocked error"))
	loadBalancersOut, err := lbs.GetLoadBalancerList()

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(loadBalancersOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return &loadBalancersOut
}

// GetLoadBalancerListFailStatusMocked test mocked function
func GetLoadBalancerListFailStatusMocked(t *testing.T, loadBalancersIn *[]types.LoadBalancer) *[]types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// to json
	lbIn, err := json.Marshal(loadBalancersIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Get", "/v1/network/load_balancers").Return(lbIn, 499, nil)
	loadBalancersOut, err := lbs.GetLoadBalancerList()

	assert.NotNil(err, "We are expecting a status code error")
	assert.Nil(loadBalancersOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return &loadBalancersOut
}

// GetLoadBalancerListFailJSONMocked test mocked function
func GetLoadBalancerListFailJSONMocked(t *testing.T, loadBalancersIn *[]types.LoadBalancer) *[]types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// wrong json
	lbIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", "/v1/network/load_balancers").Return(lbIn, 200, nil)
	loadBalancersOut, err := lbs.GetLoadBalancerList()

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(loadBalancersOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return &loadBalancersOut
}

// GetLoadBalancerMocked test mocked function
func GetLoadBalancerMocked(t *testing.T, loadBalancer *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// to json
	lbIn, err := json.Marshal(loadBalancer)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/network/load_balancers/%s", loadBalancer.Id)).Return(lbIn, 200, nil)
	loadBalancerOut, err := lbs.GetLoadBalancer(loadBalancer.Id)
	assert.Nil(err, "Error getting loadBalancer")
	assert.Equal(*loadBalancer, *loadBalancerOut, "GetLoadBalancer returned different loadBalancers")

	return loadBalancerOut
}

// GetLoadBalancerFailErrMocked test mocked function
func GetLoadBalancerFailErrMocked(t *testing.T, loadBalancer *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// to json
	lbIn, err := json.Marshal(loadBalancer)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/network/load_balancers/%s", loadBalancer.Id)).Return(lbIn, 200, fmt.Errorf("Mocked error"))
	loadBalancerOut, err := lbs.GetLoadBalancer(loadBalancer.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return loadBalancerOut
}

// GetLoadBalancerFailStatusMocked test mocked function
func GetLoadBalancerFailStatusMocked(t *testing.T, loadBalancer *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// to json
	lbIn, err := json.Marshal(loadBalancer)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/network/load_balancers/%s", loadBalancer.Id)).Return(lbIn, 499, nil)
	loadBalancerOut, err := lbs.GetLoadBalancer(loadBalancer.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return loadBalancerOut
}

// GetLoadBalancerFailJSONMocked test mocked function
func GetLoadBalancerFailJSONMocked(t *testing.T, loadBalancer *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// wrong json
	lbIn := []byte{10, 20, 30}

	// call service
	cs.On("Get", fmt.Sprintf("/v1/network/load_balancers/%s", loadBalancer.Id)).Return(lbIn, 200, nil)
	loadBalancerOut, err := lbs.GetLoadBalancer(loadBalancer.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return loadBalancerOut
}

// CreateLoadBalancerMocked test mocked function
func CreateLoadBalancerMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Post", "/v1/network/load_balancers/", mapIn).Return(dOut, 200, nil)
	loadBalancerOut, err := lbs.CreateLoadBalancer(mapIn)
	assert.Nil(err, "Error creating loadBalancer list")
	assert.Equal(loadBalancerIn, loadBalancerOut, "CreateLoadBalancer returned different loadBalancers")

	return loadBalancerOut
}

// CreateLoadBalancerFailErrMocked test mocked function
func CreateLoadBalancerFailErrMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Post", "/v1/network/load_balancers/", mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	loadBalancerOut, err := lbs.CreateLoadBalancer(mapIn)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return loadBalancerOut
}

// CreateLoadBalancerFailStatusMocked test mocked function
func CreateLoadBalancerFailStatusMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Post", "/v1/network/load_balancers/", mapIn).Return(dOut, 499, nil)
	loadBalancerOut, err := lbs.CreateLoadBalancer(mapIn)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")

	return loadBalancerOut
}

// CreateLoadBalancerFailJSONMocked test mocked function
func CreateLoadBalancerFailJSONMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// wrong json
	lbIn := []byte{10, 20, 30}

	// call service
	cs.On("Post", "/v1/network/load_balancers/", mapIn).Return(lbIn, 200, nil)
	loadBalancerOut, err := lbs.CreateLoadBalancer(mapIn)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return loadBalancerOut
}

// UpdateLoadBalancerMocked test mocked function
func UpdateLoadBalancerMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/network/load_balancers/%s", loadBalancerIn.Id), mapIn).Return(dOut, 200, nil)
	loadBalancerOut, err := lbs.UpdateLoadBalancer(mapIn, loadBalancerIn.Id)
	assert.Nil(err, "Error updating loadBalancer list")
	assert.Equal(loadBalancerIn, loadBalancerOut, "UpdateLoadBalancer returned different loadBalancers")

	return loadBalancerOut
}

// UpdateLoadBalancerFailErrMocked test mocked function
func UpdateLoadBalancerFailErrMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/network/load_balancers/%s", loadBalancerIn.Id), mapIn).Return(dOut, 200, fmt.Errorf("Mocked error"))
	loadBalancerOut, err := lbs.UpdateLoadBalancer(mapIn, loadBalancerIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")

	return loadBalancerOut
}

// UpdateLoadBalancerFailStatusMocked test mocked function
func UpdateLoadBalancerFailStatusMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// to json
	dOut, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Put", fmt.Sprintf("/v1/network/load_balancers/%s", loadBalancerIn.Id), mapIn).Return(dOut, 499, nil)
	loadBalancerOut, err := lbs.UpdateLoadBalancer(mapIn, loadBalancerIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
	return loadBalancerOut
}

// UpdateLoadBalancerFailJSONMocked test mocked function
func UpdateLoadBalancerFailJSONMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) *types.LoadBalancer {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// convertMap
	mapIn, err := utils.ItemConvertParams(*loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// wrong json
	lbIn := []byte{10, 20, 30}

	// call service
	cs.On("Put", fmt.Sprintf("/v1/network/load_balancers/%s", loadBalancerIn.Id), mapIn).Return(lbIn, 200, nil)
	loadBalancerOut, err := lbs.UpdateLoadBalancer(mapIn, loadBalancerIn.Id)

	assert.NotNil(err, "We are expecting a marshalling error")
	assert.Nil(loadBalancerOut, "Expecting nil output")
	assert.Contains(err.Error(), "invalid character", "Error message should include the string 'invalid character'")

	return loadBalancerOut
}

// DeleteLoadBalancerMocked test mocked function
func DeleteLoadBalancerMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// to json
	lbIn, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/network/load_balancers/%s", loadBalancerIn.Id)).Return(lbIn, 200, nil)
	err = lbs.DeleteLoadBalancer(loadBalancerIn.Id)
	assert.Nil(err, "Error deleting loadBalancer")
}

// DeleteLoadBalancerFailErrMocked test mocked function
func DeleteLoadBalancerFailErrMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// to json
	lbIn, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/network/load_balancers/%s", loadBalancerIn.Id)).Return(lbIn, 200, fmt.Errorf("Mocked error"))
	err = lbs.DeleteLoadBalancer(loadBalancerIn.Id)

	assert.NotNil(err, "We are expecting an error")
	assert.Equal(err.Error(), "Mocked error", "Error should be 'Mocked error'")
}

// DeleteLoadBalancerFailStatusMocked test mocked function
func DeleteLoadBalancerFailStatusMocked(t *testing.T, loadBalancerIn *types.LoadBalancer) {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// to json
	lbIn, err := json.Marshal(loadBalancerIn)
	assert.Nil(err, "LoadBalancer test data corrupted")

	// call service
	cs.On("Delete", fmt.Sprintf("/v1/network/load_balancers/%s", loadBalancerIn.Id)).Return(lbIn, 499, nil)
	err = lbs.DeleteLoadBalancer(loadBalancerIn.Id)

	assert.NotNil(err, "We are expecting an status code error")
	assert.Contains(err.Error(), "499", "Error should contain http code 499")
}

// GetLBNodeListMocked test mocked function
func GetLBNodeListMocked(t *testing.T, lbnodesIn *[]types.LoadBalancer, lbID string) *[]types.LBNode {

	assert := assert.New(t)

	// wire up
	cs := &utils.MockConcertoService{}
	lbs, err := NewLoadBalancerService(cs)
	assert.Nil(err, "Couldn't load loadBalancer service")
	assert.NotNil(lbs, "LoadBalancer service not instanced")

	// to json
	lbnsIn, err := json.Marshal(lbnodesIn)
	assert.Nil(err, "lbNode test data corrupted")

	// call service
	cs.On("Get", fmt.Sprintf("/v1/network/load_balancers/%s/nodes", lbID)).Return(lbnsIn, 200, nil)
	lbnsOut, err := lbs.GetLBNodeList(lbID)
	assert.Nil(err, "Error getting lbNode list")
	assert.Equal(*lbnodesIn, *lbnsOut, "GetLBNodeList returned different lbNodes")

	return lbnsOut
}

// // CreateLBNodeMocked test mocked function
// func CreateLBNodeMocked(t *testing.T, lbn *types.LBNode, lbID string) *types.LBNode {

// 	fmt.Println("in createLBMocked")
// 	assert := assert.New(t)

// 	// wire up
// 	cs := &utils.MockConcertoService{}
// 	lbs, err := NewLoadBalancerService(cs)
// 	assert.Nil(err, "Couldn't load loadBalancer service")
// 	assert.NotNil(lbs, "LoadBalancer service not instanced")

// 	// convertMap
// 	mapIn, err := utils.ItemConvertParams(*lbn)
// 	assert.Nil(err, "lbNode test data corrupted")

// 	// to json
// 	lbnIn, err := json.Marshal(lbn)
// 	assert.Nil(err, "lbNode test data corrupted")

// 	// call service
// 	cs.On("Post", fmt.Sprintf("/v1/network/load_balancers/%s/nodes", lbID), mapIn).Return(lbnIn, 200, nil)
// 	lbnOut, err := lbs.CreateLBNode(mapIn, lbID)
// 	assert.Nil(err, "Error getting lbNode")
// 	assert.Equal(*lbn, *lbnOut, "CreateLBNode returned different lbNodes")

// 	return lbnOut
// }

// // DeleteLBNodeMocked test mocked function
// func DeleteLBNodeMocked(t *testing.T, lbn *types.LBNode, lbID string) {

// 	assert := assert.New(t)

// 	// wire up
// 	cs := &utils.MockConcertoService{}
// 	lbs, err := NewLoadBalancerService(cs)
// 	assert.Nil(err, "Couldn't load loadBalancer service")
// 	assert.NotNil(lbs, "LoadBalancer service not instanced")

// 	// to json
// 	lbnIn, err := json.Marshal(lbn)
// 	assert.Nil(err, "lbNode test data corrupted")

// 	// call service
// 	cs.On("Delete", fmt.Sprintf("/v1/network/load_balancers/%s/nodes/%s", lbID, lbn.Id)).Return(lbnIn, 200, nil)
// 	err = lbs.DeleteLBNode(lbID, lbn.Id)
// 	assert.Nil(err, "Error deleting lbNode")
// }
