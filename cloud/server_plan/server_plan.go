package server_plan

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

type ServerPlan struct {
	Id              string  `json:"id"`
	Name            string  `json:"name"`
	Memory          int     `json:"memory"`
	CPUs            float32 `json:"cpus"`
	Storage         int     `json:"storage"`
	LocationId      string  `json:"location_id"`
	CloudProviderId string  `json:"cloud_provider_id"`
}

func cmdList(c *cli.Context) {
	utils.FlagsRequired(c, []string{"cloud_provider_id"})
	var serverPlans []ServerPlan

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/cloud/cloud_providers/%s/server_plans", c.String("cloud_provider_id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &serverPlans)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tMEMORY\tCPUS\tSTORAGE\tLOCATION ID\tCLOUD PROVIDER ID\r")

	for _, sp := range serverPlans {
		fmt.Fprintf(w, "%s\t%s\t%d\t%g\t%d\t%s\t%s\n", sp.Id, sp.Name, sp.Memory, sp.CPUs, sp.Storage, sp.LocationId, sp.CloudProviderId)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var sp ServerPlan

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/cloud/server_plans/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &sp)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tMEMORY\tCPUS\tSTORAGE\tLOCATION ID\tCLOUD PROVIDER ID\r")
	fmt.Fprintf(w, "%s\t%s\t%d\t%g\t%d\t%s\t%s\n", sp.Id, sp.Name, sp.Memory, sp.CPUs, sp.Storage, sp.LocationId, sp.CloudProviderId)

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "This action lists the server plans offered by the cloud provider identified by the given id.",
			Action: cmdList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cloud_provider_id",
					Usage: "Cloud provider id",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "This action shows information about the Server Plan identified by the given id.",
			Action: cmdShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server plan id",
				},
			},
		},
	}
}
