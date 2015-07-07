package cloud_providers

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

type CloudProviders struct {
	Id                  string   `json:"id"`
	Name                string   `json:"name"`
	RequiredCredentials []string `json:"required_credentials"`
	ProvidedServices    []string `json:"provided_services"`
}

func cmdList(c *cli.Context) {
	utils.FlagsRequired(c, []string{"app_id", "location_id"})
	var cps []CloudProviders

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/wizard/cloud_providers?app_id=%s&location_id=%s", c.String("app_id"), c.String("location_id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &cps)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tREQUIRED CREDENTIALS\tPROVIDED SERVICES\r")

	for _, cp := range cps {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", cp.Id, cp.Name, cp.RequiredCredentials, cp.ProvidedServices)
	}

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Cloud Providers",
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
			},
		},
	}
}
