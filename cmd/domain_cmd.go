package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpDomain prepares common resources to send request to Concerto API
func WireUpDomain(c *cli.Context) (ds *api.DomainService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = api.NewDomainService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up domain service", err)
	}

	return ds, f
}

// DomainList subcommand function
func DomainList(c *cli.Context) {
	debugCmdFuncInfo(c)
	domainSvc, formatter := WireUpDomain(c)

	domains, err := domainSvc.GetDomainList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive domain data", err)
	}
	if err = formatter.PrintList(domains); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// DomainShow subcommand function
func DomainShow(c *cli.Context) {
	debugCmdFuncInfo(c)
	domainSvc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	domain, err := domainSvc.GetDomain(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive domain data", err)
	}
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// DomainCreate subcommand function
func DomainCreate(c *cli.Context) {
	debugCmdFuncInfo(c)
	domainSvc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"name", "contact"}, formatter)
	domain, err := domainSvc.CreateDomain(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create domain", err)
	}
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// DomainUpdate subcommand function
func DomainUpdate(c *cli.Context) {
	debugCmdFuncInfo(c)
	domainSvc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	domain, err := domainSvc.UpdateDomain(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update domain", err)
	}
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// DomainDelete subcommand function
func DomainDelete(c *cli.Context) {
	debugCmdFuncInfo(c)
	domainSvc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := domainSvc.DeleteDomain(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete domain", err)
	}
}

// DomainRecordList subcommand function
func DomainRecordList(c *cli.Context) {
	debugCmdFuncInfo(c)
	domainSvc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"domain_id"}, formatter)
	domainRecords, err := domainSvc.GetDomainRecordList(c.String("domain_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list domain records", err)
	}
	if err = formatter.PrintList(*domainRecords); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// DomainRecordShow subcommand function
func DomainRecordShow(c *cli.Context) {
	debugCmdFuncInfo(c)
	domainSvc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"domain_id", "id"}, formatter)
	domain, err := domainSvc.GetDomainRecord(c.String("domain_id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list domain records", err)
	}
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// DomainRecordCreate subcommand function
func DomainRecordCreate(c *cli.Context) {
	debugCmdFuncInfo(c)
	domainSvc, formatter := WireUpDomain(c)

	checkRequiredFlags(c, []string{"domain_id", "type", "name"}, formatter)

	switch c.String("type") {
	case "A":
		checkRequiredFlagsOr(c, []string{"content", "server_id"}, formatter)
	case "AAAA", "CNAME":
		checkRequiredFlags(c, []string{"content"}, formatter)
	case "MX":
		checkRequiredFlags(c, []string{"content", "prio"}, formatter)
	}

	domain, err := domainSvc.CreateDomainRecord(utils.FlagConvertParams(c), c.String("domain_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't create domain record", err)
	}
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// DomainRecordUpdate subcommand function
func DomainRecordUpdate(c *cli.Context) {
	domainSvc, formatter := WireUpDomain(c)
	debugCmdFuncInfo(c)
	checkRequiredFlags(c, []string{"domain_id", "id"}, formatter)

	domain, err := domainSvc.UpdateDomainRecord(utils.FlagConvertParams(c), c.String("domain_id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update domain", err)
	}
	if err = formatter.PrintItem(*domain); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// DeleteDomainRecord subcommand function
func DeleteDomainRecord(c *cli.Context) {
	domainSvc, formatter := WireUpDomain(c)
	debugCmdFuncInfo(c)
	checkRequiredFlags(c, []string{"domain_id", "id"}, formatter)

	err := domainSvc.DeleteDomainRecord(c.String("domain_id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete domain", err)
	}
}
