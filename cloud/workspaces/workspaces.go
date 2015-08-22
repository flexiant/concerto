/*
	A workspace in Concerto is an organizational unit for servers. Workspaces aggregate servers with the same firewall profile, ssh profile and DNS domain settings.

	The available commands are:
		list
		show
		create
		update
		delete
		list_workspace_servers


	Use "cloud workspaces --help" on the commandline interface for more information about the available subcommands.

	Workspaces list

	Lists all available workspaces.

	Usage:

		workspaces list

	Workspaces show

	Shows information about a specific workspace.

	Usage:

		workspaces show (options)

	Options:
		--id <workspace_id> 		Workspace Id


	Workspaces create

	Creates a new workspace to be used in the workspaces.

	Usage:

		workspaces create (options)

	Options:
		--id <workspace_id> 		Workspace Id
		--name <name> 			Logical name of the workspace
		--domain_id <domain_id> 	Identifier of the DNS domain to which the workspace ascribes its servers
		--ssh_profile_id <ssh_profile_id> 			Identifier of the ssh profile which the workspace ascribes its servers
		--firewall_profile_id <firewall_profile_id>	Identifier of the firewall profile to which the workspace ascribes its servers
	Workspaces update

	Updates an existing workspace.

	Usage:

		workspaces update (options)

	Options:
		--id <workspace_id> 		Workspace Id
		--name <name> 			Logical name of the workspace
		--domain_id <domain_id> 	Identifier of the DNS domain to which the workspace ascribes its servers
		--ssh_profile_id <ssh_profile_id> 			Identifier of the ssh profile which the workspace ascribes its servers
		--firewall_profile_id <firewall_profile_id>	Identifier of the firewall profile to which the workspace ascribes its servers

	Workspaces delete

	Deletes a workspace.

	Usage:

		workspaces delete (options)

	Options:
		--id <workspace_id> 		Workspace Id

	List workspace servers

	This action returns information about the servers belonging to the workspace identified by the given id.

	Usage:

		workspaces list_workspace_servers (options)

	Options:
		--workspace_id <workspace_id> 		Workspace Id
*/
package workspaces

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

type Workspace struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Default             bool   `json:"default"`
	Domain_id           string `json:"domain_id"`
	Ssh_profile_id      string `json:"ssh_profile_id"`
	Firewall_profile_id string `json:"firewall_profile_id"`
}

type WorkspaceServer struct {
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

func cmdList(c *cli.Context) {
	var workspaces []Workspace

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/cloud/workspaces")
	utils.CheckError(err)

	err = json.Unmarshal(data, &workspaces)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDEFAULT\tDOMAIN ID\tSSH PROFILE ID\tFIREWALL PROFILE ID\r")

	for _, workspace := range workspaces {
		fmt.Fprintf(w, "%s\t%s\t%t\t%s\t%s\t%s\n", workspace.Id, workspace.Name, workspace.Default, workspace.Domain_id, workspace.Ssh_profile_id, workspace.Firewall_profile_id)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var workspace Workspace

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/cloud/workspaces/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &workspace)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDEFAULT\tDOMAIN ID\tSSH PROFILE ID\tFIREWALL PROFILE ID\r")
	fmt.Fprintf(w, "%s\t%s\t%t\t%s\t%s\t%s\n", workspace.Id, workspace.Name, workspace.Default, workspace.Domain_id, workspace.Ssh_profile_id, workspace.Firewall_profile_id)

	w.Flush()
}

func cmdCreate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"name", "domain_id", "ssh_profile_id", "firewall_profile_id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["name"] = c.String("name")
	v["domain_id"] = c.String("domain_id")
	v["ssh_profile_id"] = c.String("ssh_profile_id")
	v["firewall_profile_id"] = c.String("firewall_profile_id")

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post("/v1/cloud/workspaces", jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)

	var workspace Workspace
	err = json.Unmarshal(res, &workspace)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDEFAULT\tDOMAIN ID\tSSH PROFILE ID\tFIREWALL PROFILE ID\r")
	fmt.Fprintf(w, "%s\t%s\t%t\t%s\t%s\t%s\n", workspace.Id, workspace.Name, workspace.Default, workspace.Domain_id, workspace.Ssh_profile_id, workspace.Firewall_profile_id)

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
	if c.IsSet("domain_id") {
		v["domain_id"] = c.String("domain_id")
	}
	if c.IsSet("ssh_profile_id") {
		v["ssh_profile_id"] = c.String("ssh_profile_id")
	}
	if c.IsSet("firewall_profile_id") {
		v["firewall_profile_id"] = c.String("firewall_profile_id")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Put(fmt.Sprintf("/v1/cloud/workspaces/%s", c.String("id")), jsonBytes)

	utils.CheckError(err)
	fmt.Println(res)

	var workspace Workspace
	err = json.Unmarshal(res, &workspace)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDEFAULT\tDOMAIN ID\tSSH PROFILE ID\tFIREWALL PROFILE ID\r")
	fmt.Fprintf(w, "%s\t%s\t%t\t%s\t%s\t%s\n", workspace.Id, workspace.Name, workspace.Default, workspace.Domain_id, workspace.Ssh_profile_id, workspace.Firewall_profile_id)

	w.Flush()
}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Delete(fmt.Sprintf("/v1/cloud/workspaces/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

	fmt.Println(res)
}

func cmdListWorkspaceServers(c *cli.Context) {
	var workspaceServers []WorkspaceServer
	utils.FlagsRequired(c, []string{"workspace_id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/cloud/workspaces/%s/servers", c.String("workspace_id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &workspaceServers)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFQDN\tSTATE\tPUBLIC IP\tWORKSPACE ID\tTEMPLATE ID\tSERVER PLAN ID\tSSH PROFILE ID\r")

	for _, workspaceServer := range workspaceServers {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", workspaceServer.Id, workspaceServer.Name, workspaceServer.Fqdn, workspaceServer.State, workspaceServer.Public_ip, workspaceServer.Workspace_id, workspaceServer.Template_id, workspaceServer.Server_plan_id, workspaceServer.Ssh_profile_id)
	}
	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available workspaces",
			Action: cmdList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific workspace",
			Action: cmdShow,
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
			Action: cmdCreate,
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
			Action: cmdUpdate,
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
			Action: cmdDelete,
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
			Action: cmdListWorkspaceServers,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "workspace_id",
					Usage: "Workspace Id",
				},
			},
		},
	}
}
