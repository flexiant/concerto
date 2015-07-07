package saas_providers

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

type SAASProvider struct {
	Id                    string   `json:"id"`
	Name                  string   `json:"name"`
	Required_account_data []string `json:"required_account_data"`
}

func cmdList(c *cli.Context) {
	var saasProviders []SAASProvider

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/cloud/saas_providers")
	utils.CheckError(err)

	err = json.Unmarshal(data, &saasProviders)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tREQUIRED ACCOUNT DATA\r")

	for _, sp := range saasProviders {
		fmt.Fprintf(w, "%s\t%s\t%s\n", sp.Id, sp.Name, sp.Required_account_data)
	}

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the SaaS providers supported by the platform.",
			Action: cmdList,
		},
	}
}
