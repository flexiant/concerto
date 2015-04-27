package ship

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

type Fleet struct {
	Ships []Ship `json:"ships"`
}

type Ship struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Fqdn      string `json:"fqdn"`
	PublicIp  string `json:"public_ip"`
	State     string `json:"state"`
	Os        string `json:"os"`
	Plan      string `json:"plan"`
	FleetName string `json:"fleet_name"`
	Master    bool   `json:"is_master"`
}

func cmdCreate(c *cli.Context) {

	parameters := false

	if !c.IsSet("fleet") {
		log.Warn("Please use parameter --fleet")
		parameters = true
	}
	if !c.IsSet("fqdn") {
		log.Warn("Please use parameter --fqdn")
		parameters = true
	}
	if !c.IsSet("name") {
		log.Warn("Please use parameter --name")
		parameters = true
	}
	if !c.IsSet("plan") {
		log.Warn("Please use parameter --plan")
		parameters = true
	}

	if parameters {
		log.Fatal("execute create -h to find out how to use it correctly")
	}

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["fleet_name"] = c.String("fleet")
	v["fqdn"] = c.String("fqdn")
	v["name"] = c.String("name")
	v["plan"] = c.String("plan")

	json, err := json.Marshal(v)
	utils.CheckError(err)

	err = webservice.Post("/v1/kaas/ships", json)
	utils.CheckError(err)
}

func cmdList(c *cli.Context) {
	var ships Fleet

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/kaas/ships")
	utils.CheckError(err)

	err = json.Unmarshal(data, &ships)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "FLEET\tMASTER\tID\tNAME\tFQDN\tIP\tSTATE")

	for _, ship := range ships.Ships {
		if ship.Master {
			fmt.Fprintf(w, "%s\t*\t%s\t%s\t%s\t%s\t%s\n", ship.FleetName, ship.Id, ship.Name, ship.Fqdn, ship.PublicIp, ship.State)
		} else {
			fmt.Fprintf(w, "%s\t\t%s\t%s\t%s\t%s\t%s\n", ship.FleetName, ship.Id, ship.Name, ship.Fqdn, ship.PublicIp, ship.State)
		}

	}

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "create",
			Usage:  "Creates a Ship",
			Action: cmdCreate,
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
			Name:   "list",
			Usage:  "Lists all available Ships",
			Action: cmdList,
		},
	}
}
