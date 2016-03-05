/*
	Server Plans describe the computational resources and the location where servers are deployed.

	The available commands are:
		list	lists all the server plans belonging to a cloud provider
		show	details about a particular server plan

	Use "cloud server_plans --help" on the commandline interface for more information about the available subcommands.

	Server plans list

	This action lists the server plans offered by the cloud provider identified by the given id.

	Usage:

		server_plans list

	Server plans show

	This action shows information about the Server Plan identified by the given id.

	Usage:

		server_plans show --id <server_plan_id>
*/
package server_plan

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/codegangsta/cli"
// 	"github.com/flexiant/concerto/api/types"
// 	"github.com/flexiant/concerto/utils"
// 	"github.com/flexiant/concerto/webservice"
// 	"os"
// 	"text/tabwriter"
// )

// func cmdList(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"cloud_provider_id"})
// 	var serverPlans []types.ServerPlan

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/cloud/cloud_providers/%s/server_plans", c.String("cloud_provider_id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &serverPlans)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tMEMORY\tCPUS\tSTORAGE\tLOCATION ID\tCLOUD PROVIDER ID\r")

// 	for _, sp := range serverPlans {
// 		fmt.Fprintf(w, "%s\t%s\t%d\t%g\t%d\t%s\t%s\n", sp.Id, sp.Name, sp.Memory, sp.CPUs, sp.Storage, sp.LocationId, sp.CloudProviderId)
// 	}

// 	w.Flush()
// }

// func cmdShow(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})
// 	var sp types.ServerPlan

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/cloud/server_plans/%s", c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &sp)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tMEMORY\tCPUS\tSTORAGE\tLOCATION ID\tCLOUD PROVIDER ID\r")
// 	fmt.Fprintf(w, "%s\t%s\t%d\t%g\t%d\t%s\t%s\n", sp.Id, sp.Name, sp.Memory, sp.CPUs, sp.Storage, sp.LocationId, sp.CloudProviderId)

// 	w.Flush()
// }
