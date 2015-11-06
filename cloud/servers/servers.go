/*
	Servers represent computational resources or virtual machines the user deploys on a cloud provider.

	The available commands are:
		show
		create
		update
		boot
		reboot
		shutdown
		override
		delete
		list_dns_records
		list_events
		list_operational_Scripts
		execute_script

	Use "cloud servers --help" on the commandline interface for more information about the available subcommands.

	Servers show

	This action returns information about the server identified by the given id.

	Usage:

		servers show (options)

	Options:
		--id <server_id> 		Server Id


	Servers commission

	This action creates a new server and commissions it at the cloud provider.

	Usage:

		servers create (options)

	Options:
		--name <name> 			Logical name of the server
		--workspace_id <workspace_id> 	Identifier of the workspace to which the server will belong
		--template_id <template_id> 			Identifier of the template the server will use
		--server_plan_id <server_plan_id>	Identifier of the server plan in which the server shall be deployed

	Servers update

	Updates an existing server.

	Usage:

		servers update (options)

	Options:
		--id <server_id> 		Server Id
		--name <name> 			Logical name of the server
		--fqdn <fqdn> 	Fully qualified domain name (FQDN) of the server

	Server boot

	This action boots the server with the given id. The server must be in an inactive state.

	Usage:

		servers boot (options)

	Options:
		--id <server_id> 		Server Id

	Server reboot

	This action reboots the server with the given id. The server must be in an operational state.

	Usage:

		servers reboot (options)

	Options:
		--id <server_id> 		Server Id

	Server shutdown

	This action boots the server with the given id. The server must be in a bootstrapping, operational or stalled state.

	Usage:

		servers shutdown (options)

	Options:
		--id <server_id> 		Server Id

	Server override

	This action takes the server with the given id from a stalled state to the operational state, at the user's own risk.

	Usage:

		servers override_server (options)

	Options:
		--id <server_id> 		Server Id

	Server decommission

	This action decommissions the server with the given id. The server must be in a inactive, stalled or commission_stalled state.

	Usage:

		servers delete (options)

	Options:
		--id <server_id> 		Server Id

	List DNS records of a server

	This action returns information on the DNS records associated to the server with the given id.

	Usage:

		servers list_dns_records (options)

	Options:
		--id <server_id> 		Server Id

	List events of a server

	This action returns information about the events related to the server with the given id.

	Usage:

		servers list_events (options)

	Options:
		--id <server_id> 		Server Id

	List operational scripts of a server

	This action returns information about the operational scripts characterisations related to the server with the given id.

	Usage:

		servers list_operational_scripts (options)

	Options:
		--id <server_id> 		Server Id

	Execute script on server

	This action initiates the execution of the script characterisation with the given id on the server with the given id.

	Usage:

		servers execute_script (options)

	Options:
		--server_id <server_id> 		Server Id
		--script_id	<script_id>			Script Id
*/
package servers

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
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

type Dns struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	Type      string `json:"type"`
	IsFQDN    bool   `json:"is_fqdn"`
	Domain_id string `json:"domain_id"`
}

type Event struct {
	Id          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	Level       string    `json:"level"`
	Header      string    `json:"header"`
	Description string    `json:"description"`
}

type ScriptChar struct {
	Id               string   `json:"id"`
	Type             string   `json:"type"`
	Parameter_values struct{} `json:"parameter_values"`
	Template_id      string   `json:"template_id"`
	Script_id        string   `json:"script_id"`
}

func cmdShow(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var server Server

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/cloud/servers/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &server)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFQDN\tSTATE\tPUBLIC IP\tWORKSPACE ID\tTEMPLATE ID\tSERVER PLAN ID\tSSH PROFILE ID\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", server.Id, server.Name, server.Fqdn, server.State, server.Public_ip, server.Workspace_id, server.Template_id, server.Server_plan_id, server.Ssh_profile_id)

	w.Flush()
}

func cmdCommission(c *cli.Context) {
	utils.FlagsRequired(c, []string{"name", "fqdn", "workspace_id", "template_id", "server_plan_id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["name"] = c.String("name")
	v["fqdn"] = c.String("fqdn")
	v["workspace_id"] = c.String("workspace_id")
	v["template_id"] = c.String("template_id")
	v["server_plan_id"] = c.String("server_plan_id")

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post("/v1/cloud/servers/", jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)

	var server Server
	err = json.Unmarshal(res, &server)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFQDN\tSTATE\tPUBLIC IP\tWORKSPACE ID\tTEMPLATE ID\tSERVER PLAN ID\tSSH PROFILE ID\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", server.Id, server.Name, server.Fqdn, server.State, server.Public_ip, server.Workspace_id, server.Template_id, server.Server_plan_id, server.Ssh_profile_id)

	w.Flush()
}

func cmdUpdate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	if c.IsSet("name") {
		v["name"] = c.String("name")
	}
	if c.IsSet("fqdn") {
		v["fqdn"] = c.String("fqdn")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Put(fmt.Sprintf("/v1/cloud/servers/%s", c.String("id")), jsonBytes)

	utils.CheckError(err)

	var server Server
	err = json.Unmarshal(res, &server)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFQDN\tSTATE\tPUBLIC IP\tWORKSPACE ID\tTEMPLATE ID\tSERVER PLAN ID\tSSH PROFILE ID\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", server.Id, server.Name, server.Fqdn, server.State, server.Public_ip, server.Workspace_id, server.Template_id, server.Server_plan_id, server.Ssh_profile_id)

	w.Flush()
}

func cmdBoot(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, res, _ := webservice.Put(fmt.Sprintf("/v1/cloud/servers/%s/boot", c.String("id")), nil)

	utils.CheckError(err)

	var server Server
	err = json.Unmarshal(res, &server)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFQDN\tSTATE\tPUBLIC IP\tWORKSPACE ID\tTEMPLATE ID\tSERVER PLAN ID\tSSH PROFILE ID\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", server.Id, server.Name, server.Fqdn, server.State, server.Public_ip, server.Workspace_id, server.Template_id, server.Server_plan_id, server.Ssh_profile_id)

	w.Flush()
}

func cmdReboot(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, res, _ := webservice.Put(fmt.Sprintf("/v1/cloud/servers/%s/reboot", c.String("id")), nil)

	utils.CheckError(err)

	var server Server
	err = json.Unmarshal(res, &server)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFQDN\tSTATE\tPUBLIC IP\tWORKSPACE ID\tTEMPLATE ID\tSERVER PLAN ID\tSSH PROFILE ID\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", server.Id, server.Name, server.Fqdn, server.State, server.Public_ip, server.Workspace_id, server.Template_id, server.Server_plan_id, server.Ssh_profile_id)

	w.Flush()
}

func cmdShutdown(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, res, _ := webservice.Put(fmt.Sprintf("/v1/cloud/servers/%s/shutdown", c.String("id")), nil)

	utils.CheckError(err)

	var server Server
	err = json.Unmarshal(res, &server)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFQDN\tSTATE\tPUBLIC IP\tWORKSPACE ID\tTEMPLATE ID\tSERVER PLAN ID\tSSH PROFILE ID\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", server.Id, server.Name, server.Fqdn, server.State, server.Public_ip, server.Workspace_id, server.Template_id, server.Server_plan_id, server.Ssh_profile_id)

	w.Flush()
}

func cmdOverride(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, res, _ := webservice.Put(fmt.Sprintf("/v1/cloud/servers/%s/override", c.String("id")), nil)

	utils.CheckError(err)

	var server Server
	err = json.Unmarshal(res, &server)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFQDN\tSTATE\tPUBLIC IP\tWORKSPACE ID\tTEMPLATE ID\tSERVER PLAN ID\tSSH PROFILE ID\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", server.Id, server.Name, server.Fqdn, server.State, server.Public_ip, server.Workspace_id, server.Template_id, server.Server_plan_id, server.Ssh_profile_id)

	w.Flush()
}

func cmdDecommission(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, res, code := webservice.Delete(fmt.Sprintf("/v1/cloud/servers/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(code, res)

	var server Server
	err = json.Unmarshal(res, &server)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFQDN\tSTATE\tPUBLIC IP\tWORKSPACE ID\tTEMPLATE ID\tSERVER PLAN ID\tSSH PROFILE ID\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", server.Id, server.Name, server.Fqdn, server.State, server.Public_ip, server.Workspace_id, server.Template_id, server.Server_plan_id, server.Ssh_profile_id)

	w.Flush()
}

func cmdListDNS(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var dnsList []Dns

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/cloud/servers/%s/records", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &dnsList)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tCONTENT\tTYPE\tIS FQDN\tDOMAIN ID\r")

	for _, dns := range dnsList {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%t\t%s\n", dns.Id, dns.Name, dns.Content, dns.Type, dns.IsFQDN, dns.Domain_id)
	}

	w.Flush()
}

func cmdListEvents(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var events []Event

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/cloud/servers/%s/events", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &events)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTIMESTAMP\tLEVEL\tHEADER\tDESCRIPTION\r")

	for _, event := range events {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", event.Id, event.Timestamp, event.Level, event.Header, event.Description)
	}

	w.Flush()
}

func cmdListScripts(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var scripts []ScriptChar

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/cloud/servers/%s/operational_scripts", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &scripts)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTYPE\tPARAMETER VALUES\tTEMPLATE ID\tSCRIPT ID\r")

	for _, script := range scripts {
		fmt.Fprintf(w, "%s\t%s\t%t\t%s\t%t\n", script.Id, script.Type, script.Parameter_values, script.Template_id, script.Script_id)
	}

	w.Flush()
}

func cmdExecuteScript(c *cli.Context) {
	utils.FlagsRequired(c, []string{"server_id", "script_id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, res, _ := webservice.Put(fmt.Sprintf("/v1/cloud/servers/%s/operational_scripts/%s/execute", c.String("server_id"), c.String("script_id")), nil)
	utils.CheckError(err)

	var event Event
	err = json.Unmarshal(res, &event)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTIMESTAMP\tLEVEL\tHEADER\tDESCRIPTION\r")
	fmt.Fprintf(w, "%s\t%s\t%t\t%s\t%t\n", event.Id, event.Timestamp, event.Level, event.Header, event.Description)

	w.Flush()
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
			Name:   "create",
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
			Name:   "delete",
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
