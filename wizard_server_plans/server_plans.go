package wizard_server_plans

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
	var sps []ServerPlan

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/wizard/server_plans")
	utils.CheckError(err)

	err = json.Unmarshal(data, &sps)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tMEMORY\tCPUS\tSTORAGE\tLOCATION ID\tCLOUD PROVIDER ID\r")

	for _, sp := range sps {
		fmt.Fprintf(w, "%s\t%s\n", sp.Id, sp.Name, sp.Memory, sp.CPUs, sp.Storage, sp.LocationId, sp.CloudProviderId)
	}

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Server Plans.",
			Action: cmdList,
		},
	}
}
