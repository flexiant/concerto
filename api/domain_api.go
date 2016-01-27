package api

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/utils"
)

// Domain represents a domain entry
type Domain struct {
	ID      string `json:"id" header:"ID"`
	Name    string `json:"name" header:"NAME"`
	TTL     int    `json:"ttl" header:"TTL"`
	Contact string `json:"contact" header:"CONTACT"`
	Minimum int    `json:"minimum" header:"MINIMUM"`
	Enabled bool   `json:"enabled" header:"ENABLED"`
}

// DomainRecord represents a domain record entry
type DomainRecord struct {
	ID       string `json:"id" header:"ID"`
	Type     string `json:"type" header:"TYPE"`
	Name     string `json:"name" header:"NAME"`
	Content  string `json:"content" header:"CONTENT"`
	TTL      int    `json:"ttl" header:"TTL"`
	Prio     int    `json:"prio" header:"PRIO"`
	ServerID string `json:"server_id" header:"SERVER ID"`
	DomainID string `json:"domain_id" header:"DOMAIN ID"`
}

// DomainService manages domain operations
type DomainService struct {
	concertoService utils.ConcertoService
}

// NewDomainService returns a Concerto domain service
func NewDomainService(concertoService utils.ConcertoService) (*DomainService, error) {
	if concertoService == nil {
		return nil, fmt.Errorf("Must initialize ConcertoService before using it")
	}

	return &DomainService{
		concertoService: concertoService,
	}, nil
}

// GetDomainList returns the list of domains as an array of Domain
func (dm *DomainService) GetDomainList() (domains []Domain, err error) {
	log.Debug("GetDomainList")

	data, status, err := dm.concertoService.Get("/v1/dns/domains")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {

	}

	if err = json.Unmarshal(data, &domains); err != nil {
		return nil, err
	}

	return domains, nil
}

// GetDomain returns a domain by its ID
func (dm *DomainService) GetDomain(ID string) (domain *Domain, err error) {
	log.Debug("GetDomainList")

	data, status, err := dm.concertoService.Get(fmt.Sprintf("/v1/dns/domains/%s", ID))
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &domain); err != nil {
		return nil, err
	}

	return domain, nil
}

// CreateDomain returns a domain by its ID
func (dm *DomainService) CreateDomain(domainVector *map[string]string) (domain *Domain, err error) {
	log.Debug("CreateDomain")

	data, status, err := dm.concertoService.Post("/v1/dns/domains/", domainVector)
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &domain); err != nil {
		return nil, err
	}

	return domain, nil
}
