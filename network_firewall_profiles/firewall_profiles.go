package network_firewall_profiles

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

type FirewallProfile struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Default     bool   `json:"default"`
	Rules       []Rule `json:"rules"`
}

type Rule struct {
	Protocol string `json:"protocol"`
	MinPort  int    `json:"min_port"`
	MaxPort  int    `json:"max_port"`
	CidrIp   string `json:"cidr_ip"`
}

func cmdList(c *cli.Context) {
	var firewallProfiles []FirewallProfile

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/network/firewall_profiles")
	utils.CheckError(err)

	err = json.Unmarshal(data, &firewallProfiles)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tDEFAULT\r")

	for _, firewallProfile := range firewallProfiles {
		fmt.Fprintf(w, "%s\t%s\t%s\t%t\n", firewallProfile.Id, firewallProfile.Name, firewallProfile.Description, firewallProfile.Default)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var firewallProfile FirewallProfile

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/network/firewall_profiles/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &firewallProfile)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tDEFAULT\tRULES\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%t\t%s\n", firewallProfile.Id, firewallProfile.Name, firewallProfile.Description, firewallProfile.Default, firewallProfile.Rules)

	w.Flush()
}

func cmdCreate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"name", "description"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["name"] = c.String("name")
	v["description"] = c.String("description")
	if c.IsSet("rules") {
		v["rules"] = c.String("rules")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post("/v1/network/firewall_profiles", jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)

}

func cmdUpdate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	if c.IsSet("name") {
		v["name"] = c.String("name")
	}
	if c.IsSet("description") {
		v["description"] = c.String("description")
	}
	if c.IsSet("rules") {
		v["rules"] = c.String("rules")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Put(fmt.Sprintf("/v1/network/firewall_profiles/%s", c.String("id")), bytes.NewReader(jsonBytes))

	utils.CheckError(err)
	fmt.Println(res)
}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Delete(fmt.Sprintf("/v1/network/firewall_profiles/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

	fmt.Println(res)
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all existing firewall profiles",
			Action: cmdList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about the firewall profile identified by the given id.",
			Action: cmdShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Firewall profile Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a a firewall profile with the given parameters.",
			Action: cmdCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the firewall profile",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the firewall profile",
				},
				cli.StringFlag{
					Name:  "rules",
					Usage: "Set of rules of the firewall profile",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates the firewall profile identified by the given id with the given parameters.",
			Action: cmdUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Firewall profile Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the firewall profile",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the firewall profile",
				},
				cli.StringFlag{
					Name:  "rules",
					Usage: "Set of rules of the firewall profile",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Destroy a firewall profile",
			Action: cmdDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Firewall profile Id",
				},
			},
		},
	}
}
