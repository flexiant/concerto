package cloud

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// SSHProfileService manages sshProfile operations
type SSHProfileService struct {
	concertoService utils.ConcertoService
}

// NewSSHProfileService returns a Concerto sshProfile service
func NewSSHProfileService(concertoService utils.ConcertoService) (*SSHProfileService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &SSHProfileService{
		concertoService: concertoService,
	}, nil
}

// GetSSHProfileList returns the list of sshProfiles as an array of SSHProfile
func (dm *SSHProfileService) GetSSHProfileList() (sshProfiles []types.SSHProfile, err error) {
	log.Debug("GetSSHProfileList")

	data, status, err := dm.concertoService.Get("/v1/cloud/ssh_profiles")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &sshProfiles); err != nil {
		return nil, err
	}

	return sshProfiles, nil
}

// GetSSHProfile returns a sshProfile by its ID
func (dm *SSHProfileService) GetSSHProfile(ID string) (sshProfile *types.SSHProfile, err error) {
	log.Debug("GetSSHProfile")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/v1/cloud/ssh_profiles/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &sshProfile); err != nil {
		return nil, err
	}

	return sshProfile, nil
}

// CreateSSHProfile creates a sshProfile
func (dm *SSHProfileService) CreateSSHProfile(sshProfileVector *map[string]interface{}) (sshProfile *types.SSHProfile, err error) {
	log.Debug("CreateSSHProfile")

	data, status, err := dm.concertoService.Post("/v1/cloud/ssh_profiles/", sshProfileVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &sshProfile); err != nil {
		return nil, err
	}

	return sshProfile, nil
}

// UpdateSSHProfile updates a sshProfile by its ID
func (dm *SSHProfileService) UpdateSSHProfile(sshProfileVector *map[string]interface{}, ID string) (sshProfile *types.SSHProfile, err error) {
	log.Debug("UpdateSSHProfile")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/v1/cloud/ssh_profiles/%s", ID), sshProfileVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &sshProfile); err != nil {
		return nil, err
	}

	return sshProfile, nil
}

// DeleteSSHProfile deletes a sshProfile by its ID
func (dm *SSHProfileService) DeleteSSHProfile(ID string) (err error) {
	log.Debug("DeleteSSHProfile")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/v1/cloud/ssh_profiles/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
