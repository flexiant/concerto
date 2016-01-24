package api

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/flexiant/concerto/utils"
	"strconv"
)

// Domain represents a domain entry
type Domain struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	TTL     int    `json:"ttl"`
	Contact string `json:"contact"`
	Minimum int    `json:"minimum"`
	Enabled bool   `json:"enabled"`
}

// DomainRecord represents a domain record entry
type DomainRecord struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	TTL      int    `json:"ttl"`
	Prio     int    `json:"prio"`
	ServerID string `json:"server_id"`
	DomainID string `json:"domain_id"`
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

// GetDomainListforPrinting returns the list of domains as an array of Map
// This function is intended for printing
func (dm *DomainService) GetDomainListforPrinting() (domains [][]string, headers []string, err error) {
	log.Debug("GetDomainListforPrinting")

	dl, err := dm.GetDomainList()
	if err != nil {
		return nil, nil, err
	}

	// domains to array. keep same order as header!
	for _, d := range dl {
		domain := []string{d.ID, d.Name, strconv.Itoa(d.TTL), d.Contact,
			strconv.Itoa(d.Minimum), strconv.FormatBool(d.Enabled)}
		domains = append(domains, domain)
	}
	headerDomain := []string{"ID", "Name", "TTL", "Contact", "Minimum", "Enabled"}

	return domains, headerDomain, nil
}

// GetDomainList returns the list of domains as an array of Domain
// This function is intended for processing of domains
func (dm *DomainService) GetDomainList() (domains []Domain, err error) {
	log.Debug("GetDomainList")

	data, status, err := dm.concertoService.Get("/v1/dns/domains")
	if err != nil {
		return nil, err
	}

	if err = utils.CheckStandardStatus(status, data); err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &domains)
	utils.CheckError(err)

	return domains, nil
}
