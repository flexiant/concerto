package cloud_servers

import (
	// "bytes"
	// "encoding/json"
	// "fmt"
	// log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	// "github.com/flexiant/concerto/utils"
	// "github.com/flexiant/concerto/webservice"
	// "os"
	// "text/tabwriter"
	"time"
)

type Server struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Fqdn           string `json:"fqdn"`
	State          string `json:"state"`
	Public_ip      string `json:"public_ip"`
	Workspace_id   string `json:"workspace_id"`
	Template_id    string `json:"template_id"`
	Server_plan_id string `json:"server_plan_id"`
	Ssh_profile_id string `json:"ssh_profile_id"`
}

type Event struct {
	Id          string    `json:"id"`
	Timestamp   time.Time `json:"name"`
	Level       string    `json:"fqdn"`
	Header      string    `json:"state"`
	Description string    `json:"public_ip"`
}

type ScriptChar struct {
	Id               string   `json:"id"`
	Type             string   `json:"type"`
	Parameter_values struct{} `json:"parameter_values"`
	Template_id      string   `json:"template_id"`
	Script_id        string   `json:"script_id"`
}

func cmdShow(c *cli.Context) {

}

func cmdCommission(c *cli.Context) {

}

func cmdUpdate(c *cli.Context) {

}

func cmdBoot(c *cli.Context) {

}

func cmdReboot(c *cli.Context) {

}

func cmdShutdown(c *cli.Context) {

}

func cmdOverride(c *cli.Context) {

}

func cmdDecommission(c *cli.Context) {

}

func cmdListDNS(c *cli.Context) {

}

func cmdListEvents(c *cli.Context) {

}

func cmdListScripts(c *cli.Context) {

}

func cmdExecuteScript(c *cli.Context) {

}
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "show",
			Usage:  "Shows information about the server identified by the given id.",
			Action: cmdShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "comission_server",
			Usage:  "Creates a new server.",
			Action: cmdCommission,
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
			Action: cmdUpdate,
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
			Action: cmdBoot,
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
			Action: cmdReboot,
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
			Action: cmdShutdown,
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
			Action: cmdOverride,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server Id",
				},
			},
		},
		{
			Name:   "decommission_server",
			Usage:  "This action decommissions the server with the given id. The server must be in a inactive, stalled or commission_stalled state.",
			Action: cmdDecommission,
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
