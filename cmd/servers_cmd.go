package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpServer prepares common resources to send request to Concerto API
func WireUpServer(c *cli.Context) (ds *api.ServerService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = api.NewServerService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up server service", err)
	}

	return ds, f
}

// ServerList subcommand function
func ServerList(c *cli.Context) {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	servers, err := serverSvc.GetServerList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive server data", err)
	}
	if err = formatter.PrintList(servers); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ServerShow subcommand function
func ServerShow(c *cli.Context) {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.GetServer(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive server data", err)
	}
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ServerCreate subcommand function
func ServerCreate(c *cli.Context) {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"name", "fqdn", "workspace_id", "template_id", "server_plan_id"}, formatter)
	server, err := serverSvc.CreateServer(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create server", err)
	}
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ServerUpdate subcommand function
func ServerUpdate(c *cli.Context) {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.UpdateServer(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update server", err)
	}
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ServerBoot subcommand function
func ServerBoot(c *cli.Context) {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.BootServer(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't boot server", err)
	}
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ServerReboot subcommand function
func ServerReboot(c *cli.Context) {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.RebootServer(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't reboot server", err)
	}
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ServerShutdown subcommand function
func ServerShutdown(c *cli.Context) {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.ShutdownServer(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't shutdown server", err)
	}
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ServerOverride subcommand function
func ServerOverride(c *cli.Context) {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	server, err := serverSvc.OverrideServer(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't override server", err)
	}
	if err = formatter.PrintItem(*server); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ServerDelete subcommand function
func ServerDelete(c *cli.Context) {
	debugCmdFuncInfo(c)
	serverSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := serverSvc.DeleteServer(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete server", err)
	}
}

// ========= DNS ========

// DNSList subcommand function
func DNSList(c *cli.Context) {
	debugCmdFuncInfo(c)
	dnsSvc, formatter := WireUpServer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	dnss, err := dnsSvc.GetDNSList(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive dns data", err)
	}
	if err = formatter.PrintList(dnss); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}
