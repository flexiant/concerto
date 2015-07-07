package firewall

import (
	"crypto/md5"
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

type FirewallProfile stuct {
	Profile Policy `json:"firewall_profile"`
}

type Policy struct {
	Rules []Rule `json:"rules"`
	Md5   string `json:"md5"`
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
	policy.Md5 = fmt.Sprintf("%x", md5.Sum(data))
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

func check(policy Policy, rule Rule) bool {
	exists := false
	for _, policyRule := range policy.Rules {
		if (policyRule.Cidr == rule.Cidr) && (policyRule.MaxPort == rule.MaxPort) && (policyRule.MinPort == rule.MinPort) && (policyRule.Protocol == rule.Protocol) {
			exists = true
		}
	}
	return exists
}

func cmdCheck(c *cli.Context) {
	utils.FlagsRequired(c, []string{"cidr", "minPort", "maxPort", "ipProtocol"})

	newRule := &Rule{
		c.String("ipProtocol"),
		c.String("cidr"),
		c.Int("minPort"),
		c.Int("maxPort"),
	}
	policy := get()

	fmt.Printf("%t\n", check(policy, *newRule))
}

func cmdAdd(c *cli.Context) {
	utils.FlagsRequired(c, []string{"cidr", "minPort", "maxPort", "ipProtocol"})

	newRule := &Rule{
		c.String("ipProtocol"),
		c.String("cidr"),
		c.Int("minPort"),
		c.Int("maxPort"),
	}
	policy := get()

	exists := check(policy, *newRule)

	if exists == false {
		fmt.Printf("%#v", newRule)
		fmt.Printf("We are going to insert firewall")
		policy.Rules = append(policy.Rules, *newRule)
		fmt.Printf("\n\n%#v\n\n", policy)

		webservice, err := webservice.NewWebService()
		utils.CheckError(err)


		profile := &FirewallProfile{
			policy
		}
		
		json, err := json.Marshal(profile)
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
			Name:   "check",
			Usage:  "Checks if a firewall rule exists in host",
			Action: cmdCheck,
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
