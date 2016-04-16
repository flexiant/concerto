/*
	Provides information about server plans available to deploy an application on a location and a cloud provider.

	The available commands are:
		list	list all available server plans

	Use "wizard server_plans --help" on the commandline interface for more information about the available subcommands.

	Servers plans list

	Lists the available server plans.

	Usage:

		server_plans list (options)

	Options:
		--app_id	<app_id>	App identifier
		--location_id	<location_id>	Location identifier
		--cloud_provider_id		<cloud_provider_id>	Cloud provider identifier
*/
package server_plans

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
// 	utils.FlagsRequired(c, []string{"app_id", "location_id", "cloud_provider_id"})
// 	var sps []ServerPlan

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/wizard/server_plans?app_id=%s&location_id=%s&cloud_provider_id=%s", c.String("app_id"), c.String("location_id"), c.String("cloud_provider_id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &sps)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tMEMORY\tCPUS\tSTORAGE\tLOCATION ID\tCLOUD PROVIDER ID\r")

// 	for _, sp := range sps {
// 		fmt.Fprintf(w, "%s\t%s\t%d\t%g\t%d\t%s\t%s\n", sp.Id, sp.Name, sp.Memory, sp.CPUs, sp.Storage, sp.LocationId, sp.CloudProviderId)
// 	}

// 	w.Flush()
// }
