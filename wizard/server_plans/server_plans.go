package server_plans

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
	utils.FlagsRequired(c, []string{"app_id", "location_id", "cloud_provider_id"})
	var sps []ServerPlan

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/wizard/server_plans?app_id=%s&location_id=%s&cloud_provider_id=%s", c.String("app_id"), c.String("location_id"), c.String("cloud_provider_id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &sps)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tMEMORY\tCPUS\tSTORAGE\tLOCATION ID\tCLOUD PROVIDER ID\r")

	for _, sp := range sps {
		fmt.Fprintf(w, "%s\t%s\t%d\t%g\t%d\t%s\t%s\n", sp.Id, sp.Name, sp.Memory, sp.CPUs, sp.Storage, sp.LocationId, sp.CloudProviderId)
	}

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Server Plans.",
			Action: cmdList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "app_id",
					Usage: "Identifier of the App",
				},
				cli.StringFlag{
					Name:  "location_id",
					Usage: "Identifier of the Location",
				},
				cli.StringFlag{
					Name:  "cloud_provider_id",
					Usage: "Identifier of the Cloud Provider",
				},
			},
		},
	}
}
