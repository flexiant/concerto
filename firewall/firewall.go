package firewall

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

const endpoint = "cloud/firewall_profile"

type Policy struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Protocol string `json:"ip_protocol"`
	Cidr     string `json:"cidr_ip"`
	MaxPort  int    `json:"max_port"`
	MinPort  int    `json:"min_port"`
}

func list(policy Policy) error {
	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "CIDR\tPROTOCOL\tMIN\tMAX")

	for _, rule := range policy.Rules {
		fmt.Fprintf(w, "%s\t%s\t%d\t%d\n", rule.Cidr, rule.Protocol, rule.MinPort, rule.MaxPort)
	}
	w.Flush()
	return nil
}

func get() Policy {
	var policy Policy
	webservice, err := webservice.NewWebService()
	if err != nil {
		log.Fatal(err)
	}

	log.Debugf("Current firewall driver %s", driverName())
	data, err := webservice.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &policy)
	if err != nil {
		log.Fatal(err)
	}
	return policy
}

func cmdList(c *cli.Context) {
	list(get())
}

func cmdApply(c *cli.Context) {
	apply(get())
}

func cmdFlush(c *cli.Context) {
	flush()
}

func cmdAdd(c *cli.Context) {
	utils.FlagsRequired(c, []string{"cidr", "minPort", "maxPort", "ipProtocol"})

	exists := false
	newRule := &Rule{
		c.String("ipProtocol"),
		c.String("cidr"),
		c.Int("minPort"),
		c.Int("maxPort"),
	}
	policy := get()

	fmt.Printf("%#v\n\n", policy)

	for _, rule := range policy.Rules {
		if (rule.Cidr == newRule.Cidr) && (rule.MaxPort == newRule.MaxPort) && (rule.MinPort == newRule.MinPort) && (rule.Protocol == newRule.Protocol) {
			exists = true
		}
	}

	fmt.Printf("\n++++\n%#v\n++++\n", exists)

	if exists == false {
		fmt.Printf("%#v", newRule)
		fmt.Printf("We are going to insert firewall")
		policy.Rules = append(policy.Rules, *newRule)
		fmt.Printf("\n\n%#v\n\n", policy)

		webservice, err := webservice.NewWebService()
		utils.CheckError(err)

		json, err := json.Marshal(policy)
		utils.CheckError(err)
		err, res, code := webservice.Put(endpoint, json)
		if res == nil {
			log.Fatal(err)
		}
		utils.CheckError(err)
		utils.CheckReturnCode(code)
	}

}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "apply",
			Usage:  "Applies selected firewall rules in host",
			Action: cmdApply,
		},
		{
			Name:   "flush",
			Usage:  "Flushes all firewall rules from host",
			Action: cmdFlush,
		},
		{
			Name:   "add",
			Usage:  "Adds a firewall rule to host",
			Action: cmdAdd,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cidr",
					Usage: "CIDR",
				},
				cli.IntFlag{
					Name:  "minPort",
					Usage: "Minimum Port",
				},
				cli.IntFlag{
					Name:  "maxPort",
					Usage: "Maximum Port",
				},
				cli.StringFlag{
					Name:  "ipProtocol",
					Usage: "Ip protocol udp or tcp",
				},
			},
		},
		{
			Name:   "list",
			Usage:  "Lists all firewall rules associated to host",
			Action: cmdList,
		},
	}
}
