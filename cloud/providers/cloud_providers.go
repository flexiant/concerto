package providers

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

type CloudProvider struct {
	Id                  string   `json:"id"`
	Name                string   `json:"name"`
	RequiredCredentials []string `json:"required_credentials"`
	ProvidedServices    []string `json:"provided_services"`
}

func cmdList(c *cli.Context) {
	var cloudProviders []CloudProvider

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/cloud/cloud_providers")
	utils.CheckError(err)

	err = json.Unmarshal(data, &cloudProviders)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tREQUIRED CREDENTIALS\tPROVIDED SERVICES\r")

	for _, cl := range cloudProviders {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", cl.Id, cl.Name, cl.RequiredCredentials, cl.ProvidedServices)
	}

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available cloud providers.",
			Action: cmdList,
		},
	}
}
