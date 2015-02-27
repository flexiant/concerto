package firewall

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
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
			Name:   "list",
			Usage:  "Lists all firewall rules associated to host",
			Action: cmdList,
		},
	}
}
