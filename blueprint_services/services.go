package blueprint_services

import (
	"encoding/json"
	"fmt"
	// log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
	// "time"
)

type Service struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Public      bool     `json:"public"`
	License     string   `json:"license"`
	Recipes     []string `json:"recipes"`
}

func cmdList(c *cli.Context) {
	var services []Service

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/blueprint/services")
	utils.CheckError(err)

	err = json.Unmarshal(data, &services)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tPUBLIC\tLICENSE\tRECIPES\r")

	for _, service := range services {
		fmt.Fprintf(w, "%s\t%s\t%s\t%t\t%s\t%s\n", service.Id, service.Name, service.Description, service.Public, service.License, service.Recipes)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var service Service

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/blueprint/services/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &service)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tPUBLIC\tLICENSE\tRECIPES\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%t\t%s\t%s\n", service.Id, service.Name, service.Description, service.Public, service.License, service.Recipes)

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available scripts",
			Action: cmdList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific script",
			Action: cmdShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
			},
		},
	}
}
