package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/network"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpFirewallProfile prepares common resources to send request to Concerto API
func WireUpFirewallProfile(c *cli.Context) (ds *network.FirewallProfileService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = network.NewFirewallProfileService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up firewallProfile service", err)
	}

	return ds, f
}

// FirewallProfileList subcommand function
func FirewallProfileList(c *cli.Context) {
	debugCmdFuncInfo(c)
	firewallProfileSvc, formatter := WireUpFirewallProfile(c)

	firewallProfiles, err := firewallProfileSvc.GetFirewallProfileList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive firewallProfile data", err)
	}
	if err = formatter.PrintList(firewallProfiles); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// FirewallProfileShow subcommand function
func FirewallProfileShow(c *cli.Context) {
	debugCmdFuncInfo(c)
	firewallProfileSvc, formatter := WireUpFirewallProfile(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	firewallProfile, err := firewallProfileSvc.GetFirewallProfile(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive firewallProfile data", err)
	}
	if err = formatter.PrintItem(*firewallProfile); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// FirewallProfileCreate subcommand function
func FirewallProfileCreate(c *cli.Context) {
	debugCmdFuncInfo(c)
	firewallProfileSvc, formatter := WireUpFirewallProfile(c)

	checkRequiredFlags(c, []string{"name", "description"}, formatter)
	firewallProfile, err := firewallProfileSvc.CreateFirewallProfile(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create firewallProfile", err)
	}
	if err = formatter.PrintItem(*firewallProfile); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// FirewallProfileUpdate subcommand function
func FirewallProfileUpdate(c *cli.Context) {
	debugCmdFuncInfo(c)
	firewallProfileSvc, formatter := WireUpFirewallProfile(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	firewallProfile, err := firewallProfileSvc.UpdateFirewallProfile(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update firewallProfile", err)
	}
	if err = formatter.PrintItem(*firewallProfile); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// FirewallProfileDelete subcommand function
func FirewallProfileDelete(c *cli.Context) {
	debugCmdFuncInfo(c)
	firewallProfileSvc, formatter := WireUpFirewallProfile(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := firewallProfileSvc.DeleteFirewallProfile(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete firewallProfile", err)
	}
}
