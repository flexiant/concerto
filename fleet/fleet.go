package fleet

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

type Fleet struct {
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	State             string   `json:"state"`
	MasterCount       int      `json:"master_count"`
	SlaveCount        int      `json:"slave_count"`
	WorkspaceId       string   `json:"workspace_id"`
	FirewallProfileId string   `json:"firewall_profile_id"`
	MasterTemplateId  string   `json:"master_template_id"`
	SlaveTemplateId   string   `json:"slave_template_id"`
	Masters           []string `json:"masters"`
}

func cmdCreate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"fleet"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["name"] = c.String("fleet")
	if c.IsSet("domain_id") {
		v["domain_id"] = c.String("domain_id")
	}

	json, err := json.Marshal(v)
	utils.CheckError(err)

	err, _, code := webservice.Post("/v1/kaas/fleets", json)
	utils.CheckError(err)
	utils.CheckReturnCode(code)

}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Delete(fmt.Sprintf("/v1/kaas/fleets/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

}

func cmdStart(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Put(fmt.Sprintf("/v1/kaas/fleets/%s/start", c.String("id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(res)

}

func cmdStop(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Put(fmt.Sprintf("/v1/kaas/fleets/%s/stop", c.String("id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(res)

}

func cmdEmpty(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Put(fmt.Sprintf("/v1/kaas/fleets/%s/empty", c.String("id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(res)

}

func cmdAttachNet(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Put(fmt.Sprintf("/v1/kaas/fleets/%s/attach_network", c.String("id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(res)

}

func cmdList(c *cli.Context) {
	var fleets []Fleet

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/kaas/fleets")
	utils.CheckError(err)

	err = json.Unmarshal(data, &fleets)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "FLEET NAME\tID\tSTATE\tMASTER COUNT\tSLAVE COUNT\tWORKSPACE ID\tFIREWALL PROFILE ID\tMASTER TEMPLATE ID\tSLAVE TEMPLATE ID")

	for _, fleet := range fleets {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%d\t%s\t%s\t%s\t%s\n", fleet.Name, fleet.Id, fleet.State, fleet.MasterCount, fleet.SlaveCount, fleet.WorkspaceId, fleet.FirewallProfileId, fleet.MasterTemplateId, fleet.SlaveTemplateId)
	}

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available Fleets",
			Action: cmdList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "fleet",
					Usage: "Fleet Name to Attach Ship",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of Host",
				},
				cli.StringFlag{
					Name:  "fqdn",
					Usage: "Full Qualify Domain Name of Host",
				},
				cli.StringFlag{
					Name:  "plan",
					Usage: "Server Plan to Use to Create Host",
				},
			},
		},
		{
			Name:   "start",
			Usage:  "Starts a given Fleet",
			Action: cmdStart,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Fleet Id",
				},
			},
		},
		{
			Name:   "stop",
			Usage:  "Stops a given Fleet",
			Action: cmdStop,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Fleet Id",
				},
			},
		},
		{
			Name:   "empty",
			Usage:  "Empties a given Fleet",
			Action: cmdEmpty,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Fleet Id",
				},
			},
		},
		{
			Name:   "attach_net",
			Usage:  "Attaches network to a given Fleet",
			Action: cmdAttachNet,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Fleet Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a Fleet",
			Action: cmdCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "fleet",
					Usage: "Fleet Name",
				},
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Domain Id",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a given Fleet",
			Action: cmdDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Fleet Id",
				},
			},
		},
	}
}
