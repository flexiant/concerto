package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpWorkspace prepares common resources to send request to Concerto API
func WireUpWorkspace(c *cli.Context) (ds *api.WorkspaceService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = api.NewWorkspaceService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up workspace service", err)
	}

	return ds, f
}

// WorkspaceList subcommand function
func WorkspaceList(c *cli.Context) {
	debugCmdFuncInfo(c)
	workspaceSvc, formatter := WireUpWorkspace(c)

	workspaces, err := workspaceSvc.GetWorkspaceList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive workspace data", err)
	}
	if err = formatter.PrintList(workspaces); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// WorkspaceShow subcommand function
func WorkspaceShow(c *cli.Context) {
	debugCmdFuncInfo(c)
	workspaceSvc, formatter := WireUpWorkspace(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	workspace, err := workspaceSvc.GetWorkspace(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive workspace data", err)
	}
	if err = formatter.PrintItem(*workspace); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// WorkspaceCreate subcommand function
func WorkspaceCreate(c *cli.Context) {
	debugCmdFuncInfo(c)
	workspaceSvc, formatter := WireUpWorkspace(c)

	checkRequiredFlags(c, []string{"name", "domain_id", "ssh_profile_id", "firewall_profile_id"}, formatter)
	workspace, err := workspaceSvc.CreateWorkspace(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create workspace", err)
	}
	if err = formatter.PrintItem(*workspace); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// WorkspaceUpdate subcommand function
func WorkspaceUpdate(c *cli.Context) {
	debugCmdFuncInfo(c)
	workspaceSvc, formatter := WireUpWorkspace(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	workspace, err := workspaceSvc.UpdateWorkspace(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update workspace", err)
	}
	if err = formatter.PrintItem(*workspace); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// WorkspaceDelete subcommand function
func WorkspaceDelete(c *cli.Context) {
	debugCmdFuncInfo(c)
	workspaceSvc, formatter := WireUpWorkspace(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := workspaceSvc.DeleteWorkspace(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete workspace", err)
	}
}

// WorkspaceServerList subcommand function
func WorkspaceServerList(c *cli.Context) {
	debugCmdFuncInfo(c)
	workspaceSvc, formatter := WireUpWorkspace(c)

	checkRequiredFlags(c, []string{"workspace_id"}, formatter)
	workspaceServers, err := workspaceSvc.GetWorkspaceServerList(c.String("workspace_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list workspace records", err)
	}
	if err = formatter.PrintList(*workspaceServers); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}
