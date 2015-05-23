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

// type Fleet struct {
// 	Ships []Ship `json:"ships"`
// }

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
	utils.FlagsRequired(c, []string{"fleet", "plan"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["fleet_name"] = c.String("fleet")
	v["plan"] = c.String("plan")

	json, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, code := webservice.Post("/v1/kaas/ships", json)
	if res == "" {
		log.Fatal(err)
	}
	utils.CheckError(err)
	utils.CheckReturnCode(code)
	fmt.Println(res)
}

func cmdStart(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, res := webservice.Put(fmt.Sprintf("/v1/kaas/ships/%s/start", c.String("id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(res)
	fmt.Println(res)
}

func cmdStop(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, res := webservice.Put(fmt.Sprintf("/v1/kaas/ships/%s/stop", c.String("id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(res)
	fmt.Println(res)
}

func cmdRestart(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, res := webservice.Put(fmt.Sprintf("/v1/kaas/ships/%s/restart", c.String("id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(res)
	fmt.Println(res)
}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, res := webservice.Delete(fmt.Sprintf("/v1/kaas/ships/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)
	fmt.Println(res)
}

func cmdList(c *cli.Context) {
	var ships []Ship

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/kaas/ships")
	utils.CheckError(err)

	err = json.Unmarshal(data, &ships)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "FLEET\tMASTER\tID\tNAME\tFQDN\tIP\tSTATE")

	for _, ship := range ships {
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
					Name:  "plan",
					Usage: "Server Plan to Use to Create Host",
				},
			},
		},
		{
			Name:   "list",
			Usage:  "Lists all available Ships",
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
			Usage:  "Starts a given Ship",
			Action: cmdStart,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Ship Id",
				},
			},
		},
		{
			Name:   "stop",
			Usage:  "Stops a given Ship",
			Action: cmdStop,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Ship Id",
				},
			},
		},
		{
			Name:   "restart",
			Usage:  "Restarts a given Ship",
			Action: cmdRestart,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Ship Id",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a given Ship",
			Action: cmdDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Ship Id",
				},
			},
		},
	}
}
