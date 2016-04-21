package network

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
)

// FirewallProfileService manages firewallProfile operations
type FirewallProfileService struct {
	concertoService utils.ConcertoService
}

// NewFirewallProfileService returns a Concerto firewallProfile service
func NewFirewallProfileService(concertoService utils.ConcertoService) (*FirewallProfileService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &FirewallProfileService{
		concertoService: concertoService,
	}, nil
}

// GetFirewallProfileList returns the list of firewallProfiles as an array of FirewallProfile
func (dm *FirewallProfileService) GetFirewallProfileList() (firewallProfiles []types.FirewallProfile, err error) {
	log.Debug("GetFirewallProfileList")

	data, status, err := dm.concertoService.Get("/v1/network/firewall_profiles")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &firewallProfiles); err != nil {
		return nil, err
	}

	return firewallProfiles, nil
}

// GetFirewallProfile returns a firewallProfile by its ID
func (dm *FirewallProfileService) GetFirewallProfile(ID string) (firewallProfile *types.FirewallProfile, err error) {
	log.Debug("GetFirewallProfile")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/v1/network/firewall_profiles/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &firewallProfile); err != nil {
		return nil, err
	}

	return firewallProfile, nil
}

// CreateFirewallProfile creates a firewallProfile
func (dm *FirewallProfileService) CreateFirewallProfile(firewallProfileVector *map[string]interface{}) (firewallProfile *types.FirewallProfile, err error) {
	log.Debug("CreateFirewallProfile")

	data, status, err := dm.concertoService.Post("/v1/network/firewall_profiles/", firewallProfileVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &firewallProfile); err != nil {
		return nil, err
	}

	return firewallProfile, nil
}

// UpdateFirewallProfile updates a firewallProfile by its ID
func (dm *FirewallProfileService) UpdateFirewallProfile(firewallProfileVector *map[string]interface{}, ID string) (firewallProfile *types.FirewallProfile, err error) {
	log.Debug("UpdateFirewallProfile")

	data, status, err := dm.concertoService.Put(fmt.Sprintf("/v1/network/firewall_profiles/%s", ID), firewallProfileVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &firewallProfile); err != nil {
		return nil, err
	}

	return firewallProfile, nil
}

// DeleteFirewallProfile deletes a firewallProfile by its ID
func (dm *FirewallProfileService) DeleteFirewallProfile(ID string) (err error) {
	log.Debug("DeleteFirewallProfile")

	data, status, err := dm.concertoService.Delete(fmt.Sprintf("/v1/network/firewall_profiles/%s", ID))
	if err != nil {
		return err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return err
	}

	return nil
}
