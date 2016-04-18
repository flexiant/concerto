package api

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// NodeService manages node operations
type NodeService struct {
	concertoService utils.ConcertoService
}

// NewNodeService returns a Concerto node service
func NewNodeService(concertoService utils.ConcertoService) (*NodeService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &NodeService{
		concertoService: concertoService,
	}, nil
}

// GetNodeList returns the list of nodes as an array of Node
func (cl *NodeService) GetNodeList() (nodes []types.Node, err error) {
	log.Debug("GetNodeList")

	data, status, err := cl.concertoService.Get("/v1/kaas/ships")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}

// CreateNode creates a node
func (cl *NodeService) CreateNode(nodeVector *map[string]interface{}) (node *types.Node, err error) {
	log.Debug("CreateNode")

	data, status, err := cl.concertoService.Post("/v1/kaas/ships", nodeVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &node); err != nil {
		return nil, err
	}

	return node, nil
}

// DeleteNode deletes a node by its ID
func (cl *NodeService) DeleteNode(ID string) (err error) {
	log.Debug("DeleteNode")

	data, status, err := cl.concertoService.Delete(fmt.Sprintf("/v1/kaas/ships/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// StartNode starts a node by its ID
func (cl *NodeService) StartNode(nodeVector *map[string]interface{}, ID string) (err error) {
	log.Debug("StartNode")

	data, status, err := cl.concertoService.Put(fmt.Sprintf("/v1/kaas/ships/%s/start", ID), nodeVector)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// StopNode stops a node by its ID
func (cl *NodeService) StopNode(nodeVector *map[string]interface{}, ID string) (err error) {
	log.Debug("StopNode")

	data, status, err := cl.concertoService.Put(fmt.Sprintf("/v1/kaas/ships/%s/stop", ID), nodeVector)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// RestartNode empties a node by its ID
func (cl *NodeService) RestartNode(nodeVector *map[string]interface{}, ID string) (err error) {
	log.Debug("RestartNode")

	data, status, err := cl.concertoService.Put(fmt.Sprintf("/v1/kaas/ships/%s/restart", ID), nodeVector)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
