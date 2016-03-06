package servers

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists information about all the servers on this account.",
			Action: cmd.ServerList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about the server identified by the given id.",
			Action: cmd.ServerShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new server.",
			Action: cmd.ServerCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the server",
				},
				cli.StringFlag{
					Name:  "fqdn",
					Usage: "Fully qualified domain name (FQDN) of the server",
				},
				cli.StringFlag{
					Name:  "workspace_id",
					Usage: "Identifier of the workspace to which the server shall belong",
				},
				cli.StringFlag{
					Name:  "template_id",
					Usage: "Identifier of the template the server shall use",
				},
				cli.StringFlag{
					Name:  "server_plan_id",
					Usage: "Identifier of the server plan in which the server shall be deployed",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing server",
			Action: cmd.ServerUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the server",
				},
				cli.StringFlag{
					Name:  "fqdn",
					Usage: "Fully qualified domain name (FQDN) of the server",
				},
			},
		},
		{
			Name:   "boot",
			Usage:  "Boots a server with the given id",
			Action: cmd.ServerBoot,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "reboot",
			Usage:  "Reboots a server with the given id",
			Action: cmd.ServerReboot,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "shutdown",
			Usage:  "Shuts down a server with the given id",
			Action: cmd.ServerShutdown,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "override_server",
			Usage:  "This action takes the server with the given id from a stalled state to the operational state, at the user's own risk.",
			Action: cmd.ServerOverride,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "This action decommissions the server with the given id. The server must be in a inactive, stalled or commission_stalled state.",
			Action: cmd.ServerDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "list_dns_records",
			Usage:  "This action returns information on the DNS records associated to the server with the given id.",
			Action: cmdListDNS,
			// Action: cmd.DNSList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "list_events",
			Usage:  "This action returns information about the events related to the server with the given id.",
			Action: cmdListEvents,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "list_operational_scripts",
			Usage:  "This action returns information about the operational scripts characterisations related to the server with the given id.",
			Action: cmdListScripts,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "execute_script",
			Usage:  "This action initiates the execution of the script characterisation with the given id on the server with the given id.",
			Action: cmdExecuteScript,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "server_id",
					Usage: "Server Id",
				},
				cli.StringFlag{
					Name:  "script_id",
					Usage: "Script Id",
				},
			},
		},
	}
}
