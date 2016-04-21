package network

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// LoadBalancerService manages loadBalancer operations
type LoadBalancerService struct {
	concertoService utils.ConcertoService
}

// NewLoadBalancerService returns a Concerto loadBalancer service
func NewLoadBalancerService(concertoService utils.ConcertoService) (*LoadBalancerService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &LoadBalancerService{
		concertoService: concertoService,
	}, nil
}

// GetLoadBalancerList returns the list of loadBalancers as an array of LoadBalancer
func (lb *LoadBalancerService) GetLoadBalancerList() (loadBalancers []types.LoadBalancer, err error) {
	log.Debug("GetLoadBalancerList")

	data, status, err := lb.concertoService.Get("/v1/network/load_balancers")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &loadBalancers); err != nil {
		return nil, err
	}

	return loadBalancers, nil
}

// GetLoadBalancer returns a loadBalancer by its ID
func (lb *LoadBalancerService) GetLoadBalancer(ID string) (loadBalancer *types.LoadBalancer, err error) {
	log.Debug("GetLoadBalancer")

	data, status, err := lb.concertoService.Get(fmt.Sprintf("/v1/network/load_balancers/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// CreateLoadBalancer creates a loadBalancer
func (lb *LoadBalancerService) CreateLoadBalancer(loadBalancerVector *map[string]interface{}) (loadBalancer *types.LoadBalancer, err error) {
	log.Debug("CreateLoadBalancer")

	data, status, err := lb.concertoService.Post("/v1/network/load_balancers/", loadBalancerVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// UpdateLoadBalancer updates a loadBalancer by its ID
func (lb *LoadBalancerService) UpdateLoadBalancer(loadBalancerVector *map[string]interface{}, ID string) (loadBalancer *types.LoadBalancer, err error) {
	log.Debug("UpdateLoadBalancer")

	data, status, err := lb.concertoService.Put(fmt.Sprintf("/v1/network/load_balancers/%s", ID), loadBalancerVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &loadBalancer); err != nil {
		return nil, err
	}

	return loadBalancer, nil
}

// DeleteLoadBalancer deletes a loadBalancer by its ID
func (lb *LoadBalancerService) DeleteLoadBalancer(ID string) (err error) {
	log.Debug("DeleteLoadBalancer")

	data, status, err := lb.concertoService.Delete(fmt.Sprintf("/v1/network/load_balancers/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// GetLBNodeList returns a list of lBNode by loadBalancer ID
func (lb *LoadBalancerService) GetLBNodeList(loadBalancerID string) (lBNode *[]types.LBNode, err error) {
	log.Debug("ListLBNodes")

	data, status, err := lb.concertoService.Get(fmt.Sprintf("/v1/network/load_balancers/%s/nodes", loadBalancerID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &lBNode); err != nil {
		return nil, err
	}

	return lBNode, nil
}

// CreateLBNode returns a list of lBNode
func (lb *LoadBalancerService) CreateLBNode(lBNodeVector *map[string]interface{}, lbID string) (lBNode *types.LBNode, err error) {
	log.Debug("CreateLBNode")

	data, status, err := lb.concertoService.Post(fmt.Sprintf("/v1/network/load_balancers/%s/nodes", lbID), lBNodeVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &lBNode); err != nil {
		return nil, err
	}

	return lBNode, nil
}

// DeleteLBNode deletes a loadBalancer node
func (lb *LoadBalancerService) DeleteLBNode(lbID string, ID string) (err error) {
	log.Debug("DeleteLBNode")

	data, status, err := lb.concertoService.Delete(fmt.Sprintf("/v1/network/load_balancers/%s/nodes/%s", lbID, ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
