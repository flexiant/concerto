package dns

import (
	"flag"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"strings"
)

const stagineEndPointSuffix = "staging.concerto.io"

var config utils.Config
var concertoService utils.ConcertoService
var domainService *DomainService

func configureConcerto() {

	// use default config, with no flags set
	fs := flag.NewFlagSet("fsTest", flag.ContinueOnError)
	fs.String("concerto-config", "", "")
	c := cli.NewContext(nil, fs, nil)
	cfg, err := utils.InitializeConcertoConfig(c)

	if err != nil {
		panic(err)
	}

	// if endpoint is not staging
	if !strings.Contains(cfg.APIEndpoint, stagineEndPointSuffix) {
		panic(fmt.Errorf("Concerto is expected to be configured to staging.concerto.io, found %s instead", cfg.APIEndpoint))
	}

	config = *cfg
	// wireup concerto service
	cs, err := utils.NewHTTPConcertoService(cfg)
	if err != nil {
		panic(fmt.Errorf("Couldn't create the HTTP service: %s", err))
	}

	concertoService = cs

	// setup domain service
	ds, err := NewDomainService(cs)
	if err != nil {
		panic(fmt.Errorf("Couldn't create the domain service: %s", err))
	}
	domainService = ds

}

func featureContext(s *godog.Suite) {

	s.BeforeSuite(configureConcerto)

	// domains
	s.Step(`^"([^"]*)" domain doesn't exists$`, domainDoesNotExists)
	s.Step(`^"([^"]*)" domain exists$`, domainExists)
	s.Step(`^"([^"]*)" domain is created with contact "([^"]*)"`, createDomain)
	s.Step(`^"([^"]*)" domain is updated with contact "([^"]*)"`, updateDomain)
	s.Step(`^"([^"]*)" domain is deleted$`, domainIsDeleted)
	s.Step(`^"([^"]*)" domain is deleted if exists$`, domainIsDeletedIfExists)
	s.Step(`^list domains should include:$`, domainListShouldInclude)

	// domain records
	s.Step(`^records for domain "([^"]*)" are created:$`, createDomainRecords)
	s.Step(`^records for domain "([^"]*)" contain:$`, containedInDomainRecords)
	// clusters

	// services

}
