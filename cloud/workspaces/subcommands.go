package workspaces

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available workspaces",
			Action: cmd.WorkspaceList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific workspace",
			Action: cmd.WorkspaceShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Workspace Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new workspace.",
			Action: cmd.WorkspaceCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the workspace",
				},
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Identifier of the DNS domain to which the workspace ascribes its servers",
				},
				cli.StringFlag{
					Name:  "ssh_profile_id",
					Usage: "Identifier of the ssh profile which the workspace ascribes its servers",
				},
				cli.StringFlag{
					Name:  "firewall_profile_id",
					Usage: "Identifier of the firewall profile to which the workspace ascribes its servers",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing workspace",
			Action: cmd.WorkspaceUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Workspace Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the workspace ",
				},
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Identifier of the DNS domain to which the workspace ascribes its servers",
				},
				cli.StringFlag{
					Name:  "ssh_profile_id",
					Usage: "Identifier of the ssh profile which the workspace ascribes its server",
				},
				cli.StringFlag{
					Name:  "firewall_profile_id",
					Usage: "Identifier of the firewall profile to which the workspace ascribes its servers",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a workspace",
			Action: cmd.WorkspaceDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Workspace Id",
				},
			},
		},
		{
			Name:   "list_workspace_servers",
			Usage:  "Shows  the servers belonging to the workspace identified by the given id.",
			Action: cmd.WorkspaceServerList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "workspace_id",
					Usage: "Workspace Id",
				},
			},
		},
	}
}
