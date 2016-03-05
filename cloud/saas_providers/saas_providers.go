/*
	Software-as-a-Service (SaaS) providers are those entities whose services can be used within the platform.

	The available commands are:
		list	lists all the SaaS providers

	Use "cloud saas_providers --help" on the commandline interface for more information about the available subcommands

	SaaS providers list

	Lists the SaaS providers supported by the platform.

	Usage:

		saas_providers list

*/
package saas_providers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/codegangsta/cli"
// 	"github.com/flexiant/concerto/utils"
// 	"github.com/flexiant/concerto/webservice"
// 	"os"
// 	"text/tabwriter"
// )

// func cmdList(c *cli.Context) {
// 	var saasProviders []SAASProvider

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get("/v1/cloud/saas_providers")
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &saasProviders)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tREQUIRED ACCOUNT DATA\r")

// 	for _, sp := range saasProviders {
// 		fmt.Fprintf(w, "%s\t%s\t%s\n", sp.Id, sp.Name, sp.Required_account_data)
// 	}

// 	w.Flush()
// }
