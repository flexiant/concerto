package api

import (
	"fmt"
	"github.com/flexiant/concerto/api/types"
	"gopkg.in/cucumber/gherkin-go.v3"
)

func domainDoesNotExists(domainName string) error {

	domain, err := lookUpDomain(domainName)
	if err != nil {
		return err
	}
	if domain != nil {
		return fmt.Errorf("Domain %s exists.", domainName)
	}
	return nil
}

func domainExists(domainName string) error {

	domain, err := lookUpDomain(domainName)
	if err != nil {
		return err
	}
	if domain == nil {
		return fmt.Errorf("Domain %s doesn't exists.", domainName)
	}
	return nil
}

func lookUpDomain(domainName string) (*types.Domain, error) {
	domains, err := domainService.GetDomainList()
	if err != nil {
		return nil, fmt.Errorf("Couldn't list domains: %s", err.Error())
	}
	for _, domain := range domains {
		if domain.Name == domainName {
			return &domain, nil
		}
	}

	return nil, nil
}

func domainIsDeletedIfExists(domainName string) error {

	domain, err := lookUpDomain(domainName)
	if err != nil {
		return err
	}
	if domain == nil {
		return nil
	}
	err = domainService.DeleteDomain(domain.ID)
	if err != nil {
		return fmt.Errorf("Error deleting domain %s with id %s: %s", domainName, domain.ID, err.Error())
	}

	return nil
}

func domainIsDeleted(domainName string) error {

	domain, err := lookUpDomain(domainName)
	if err != nil {
		return err
	}
	if domain == nil {
		return fmt.Errorf("Domain %s does not exists", domainName)
	}
	err = domainService.DeleteDomain(domain.ID)
	if err != nil {
		return fmt.Errorf("Error deleting domain %s with id %s: %s", domainName, domain.ID, err.Error())
	}

	return nil
}

func createDomain(domainName string, contact string) error {

	v := make(map[string]interface{})
	v["name"] = domainName
	v["contact"] = contact

	domain, err := domainService.CreateDomain(&v)
	if err != nil {
		return fmt.Errorf("Couldn't create domain: %s", err.Error())
	}

	if domain.Name != domainName {
		return fmt.Errorf("Domain created, but we expected name to be %s and returned %s",
			domainName, domain.Name)
	}

	if domain.Contact != contact {
		return fmt.Errorf("Domain created, but we expected contact to be %s and returned %s",
			contact, domain.Contact)
	}

	return nil
}

func updateDomain(domainName string, contact string) error {

	domain, err := lookUpDomain(domainName)
	if err != nil {
		return err
	}

	if domain == nil {
		return fmt.Errorf("Couln't find domain %s: %s", domainName, err.Error())
	}

	v := make(map[string]interface{})
	v["name"] = domainName
	v["contact"] = contact

	domain, err = domainService.UpdateDomain(&v, domain.ID)
	if err != nil {
		return fmt.Errorf("Couldn't update domains: %s", err.Error())
	}

	if domain.Name != domainName {
		return fmt.Errorf("Domain updated, but we expected name to be %s and returned %s",
			domainName, domain.Name)
	}

	if domain.Contact != contact {
		return fmt.Errorf("Domain updated, but we expected contact to be %s and returned %s",
			contact, domain.Contact)
	}

	return nil
}

func domainListShouldInclude(domains *gherkin.DataTable) error {
	var fields []string
	var marks []string
	head := domains.Rows[0].Cells
	for _, cell := range head {
		fields = append(fields, cell.Value)
		marks = append(marks, "?")
	}

	realDomains, err := domainService.GetDomainList()
	if err != nil {
		return fmt.Errorf("Couldn't list domains: %s", err.Error())
	}

	// iterate test table
	for i := 1; i < len(domains.Rows); i++ {
		name := ""
		contact := ""
		found := false
		for n, cell := range domains.Rows[i].Cells {
			switch head[n].Value {
			case "name":
				name = cell.Value
			case "contact":
				contact = cell.Value
			}
		}

		// iterate real domains
		for _, realDomain := range realDomains {
			if realDomain.Name == name && realDomain.Contact == contact {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("Couldn't find domain with name '%s' and contact '%s'", name, contact)
		}

	}

	return nil
}
