package api

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// ClusterService manages cluster operations
type ClusterService struct {
	concertoService utils.ConcertoService
}

// NewClusterService returns a Concerto cluster service
func NewClusterService(concertoService utils.ConcertoService) (*ClusterService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &ClusterService{
		concertoService: concertoService,
	}, nil
}

// GetClusterList returns the list of clusters as an array of Cluster
func (cl *ClusterService) GetClusterList() (clusters []types.Cluster, err error) {
	log.Debug("GetClusterList")

	data, status, err := cl.concertoService.Get("/v1/kaas/fleets")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {

	}

	if err = json.Unmarshal(data, &clusters); err != nil {
		return nil, err
	}

	return clusters, nil
}

// CreateCluster creates a cluster
func (cl *ClusterService) CreateCluster(clusterVector *map[string]interface{}) (cluster *types.Cluster, err error) {
	log.Debug("CreateCluster")

	data, status, err := cl.concertoService.Post("/v1/kaas/fleets", clusterVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &cluster); err != nil {
		return nil, err
	}

	return cluster, nil
}

// DeleteCluster deletes a cluster by its ID
func (cl *ClusterService) DeleteCluster(ID string) (err error) {
	log.Debug("DeleteCluster")

	data, status, err := cl.concertoService.Delete(fmt.Sprintf("/v1/kaas/fleets/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// StartCluster starts a cluster by its ID
func (cl *ClusterService) StartCluster(clusterVector *map[string]interface{}, ID string) (err error) {
	log.Debug("StartCluster")

	data, status, err := cl.concertoService.Put(fmt.Sprintf("/v1/kaas/fleets/%s/start", ID), clusterVector)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// StopCluster stops a cluster by its ID
func (cl *ClusterService) StopCluster(clusterVector *map[string]interface{}, ID string) (err error) {
	log.Debug("StopCluster")

	data, status, err := cl.concertoService.Put(fmt.Sprintf("/v1/kaas/fleets/%s/stop", ID), clusterVector)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}

// EmptyCluster empties a cluster by its ID
func (cl *ClusterService) EmptyCluster(clusterVector *map[string]interface{}, ID string) (err error) {
	log.Debug("EmptyCluster")

	data, status, err := cl.concertoService.Put(fmt.Sprintf("/v1/kaas/fleets/%s/empty", ID), clusterVector)
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
