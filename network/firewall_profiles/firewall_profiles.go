/*
	A firewall is a device or set of devices designed to permit or deny network transmissions based upon a set of rules and is frequently used to protect networks from unauthorized access while permitting legitimate communications to pass.
  	Firewall profiles represent these sets of rules and allow their application to a set of servers.

	The available commands are:
		list
		show
		create
		update
		delete

	Use "network firewall_profiles --help" on the commandline interface for more information about the available subcommands.

	Firewall Profiles list

	Lists all available firewall profiles.

	Usage:

		firewall_profiles list

	Firewall Profiles show

	Shows information about a specific firewall profile.

	Usage:

		firewall_profiles show (options)

	Options:
		--id <firewall_profile_id> 		firewall profile id


	Firewall Profiles create

	This action creates an firewall profile with the given parameters.

	Usage:

		firewall_profiles create (options)

	Options:
		--name <name> 			Logical name of the firewall profile
		--description <description> 	Description of the firewall profile
		--rules <rules> 	Set of rules of the firewall profile, each rule having the following fields:
								a string protocol, specifying the protocol whose traffic is opened by the rule (TCP or UDP)
								an integer min_port, specifying where the port interval opened by the rule starts
								an integer max_port, specifying where the port interval opened by the rule ends and
								a string cidr_ip, specifying with the CIDR format to which network the rule opens traffic

	Firewall Profiles update

	Updates an existing firewall profile.

	Usage:

		firewall_profiles update (options)

	Options:
		--id <firewall_profile_id> 		firewall profile id
		--name <name> 			Logical name of the firewall profile
		--description <description> 	Description of the firewall profile
		--rules <rules> 	Set of rules of the firewall profile, each rule having the following fields:
								a string protocol, specifying the protocol whose traffic is opened by the rule (TCP or UDP)
								an integer min_port, specifying where the port interval opened by the rule starts
								an integer max_port, specifying where the port interval opened by the rule ends and
								a string cidr_ip, specifying with the CIDR format to which network the rule opens traffic


	Firewall Profiles delete

	Deletes an firewall profile.

	Usage:

		firewall_profiles delete (options)

	Options:
		--id <firewall_profile_id> 		firewall profile id

*/
package firewall_profiles

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

type FirewallProfile struct {
	Id          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Default     bool   `json:"default,omitempty"`
	Rules       []Rule `json:"rules,omitempty"`
}

type Rule struct {
	Protocol string `json:"ip_protocol"`
	MinPort  int    `json:"min_port"`
	MaxPort  int    `json:"max_port"`
	CidrIp   string `json:"source"`
}

func cmdList(c *cli.Context) {
	var firewallProfiles []FirewallProfile

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, data, res := webservice.Get("/v1/network/firewall_profiles")
	utils.CheckError(err)
	utils.CheckReturnCode(res, data)

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

	err, data, res := webservice.Get(fmt.Sprintf("/v1/network/firewall_profiles/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res, data)

	err = json.Unmarshal(data, &firewallProfile)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tDEFAULT\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%t\n", firewallProfile.Id, firewallProfile.Name, firewallProfile.Description, firewallProfile.Default)
	fmt.Fprintln(w, "RULES:\r")
	fmt.Fprintln(w, "\tPROTOCOL\tMIN PORT\tMAX PORT\tSOURCE\r")
	for _, r := range firewallProfile.Rules {
		fmt.Fprintf(w, "\t%s\t%d\t%d\t%s\n", r.Protocol, r.MinPort, r.MaxPort, r.CidrIp)
	}
	w.Flush()
}

func cmdCreate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"name", "description"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	fp := FirewallProfile{
		Name:        c.String("name"),
		Description: c.String("description"),
	}

	if c.IsSet("rules") {
		var rules []Rule
		err = json.Unmarshal([]byte(c.String("rules")), &rules)
		utils.CheckError(err)
		fp.Rules = rules
	}

	jsonBytes, err := json.Marshal(fp)
	utils.CheckError(err)
	err, res, code := webservice.Post("/v1/network/firewall_profiles", jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)
	utils.CheckReturnCode(code, res)

	var firewallProfile FirewallProfile

	err = json.Unmarshal(res, &firewallProfile)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tDEFAULT\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%t\n", firewallProfile.Id, firewallProfile.Name, firewallProfile.Description, firewallProfile.Default)
	fmt.Fprintln(w, "RULES:\r")
	fmt.Fprintln(w, "\tPROTOCOL\tMIN PORT\tMAX PORT\tSOURCE\r")
	for _, r := range firewallProfile.Rules {
		fmt.Fprintf(w, "\t%s\t%d\t%d\t%s\n", r.Protocol, r.MinPort, r.MaxPort, r.CidrIp)
	}
	w.Flush()

}

func cmdUpdate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	fp := FirewallProfile{
		Id: c.String("id"),
	}

	if c.IsSet("name") {
		fp.Name = c.String("name")
	}
	if c.IsSet("description") {
		fp.Description = c.String("description")
	}

	if c.IsSet("rules") {
		var rules []Rule
		err = json.Unmarshal([]byte(c.String("rules")), &rules)
		utils.CheckError(err)
		fp.Rules = rules
	}

	jsonBytes, err := json.Marshal(fp)
	utils.CheckError(err)
	err, res, code := webservice.Put(fmt.Sprintf("/v1/network/firewall_profiles/%s", c.String("id")), jsonBytes)
	utils.CheckError(err)
	utils.CheckReturnCode(code, res)

	var firewallProfile FirewallProfile

	err = json.Unmarshal(res, &firewallProfile)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tDEFAULT\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%t\n", firewallProfile.Id, firewallProfile.Name, firewallProfile.Description, firewallProfile.Default)
	fmt.Fprintln(w, "RULES:\r")
	fmt.Fprintln(w, "\tPROTOCOL\tMIN PORT\tMAX PORT\tSOURCE\r")
	for _, r := range firewallProfile.Rules {
		fmt.Fprintf(w, "\t%s\t%d\t%d\t%s\n", r.Protocol, r.MinPort, r.MaxPort, r.CidrIp)
	}
	w.Flush()

}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, mesg, res := webservice.Delete(fmt.Sprintf("/v1/network/firewall_profiles/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res, mesg)
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
