/*

	Reports allow the user to have information about the historical uptime of their servers.

	The available commands are:
		list	reports related to the account group
		show	details about a particular report associated to the account group of the authenticated user

	Use "reports --help" on the commandline interface for more information about the available subcommands.

	Reports list

	Returns information about the reports related to the account group.

	Usage:

		reports list

	Reports show

	Returns details about a particular report associated to the account group of the authenticated user.

	Usage:

		reports show --id <report_id>
*/
package reports

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/codegangsta/cli"
// 	"github.com/flexiant/concerto/utils"
// 	"github.com/flexiant/concerto/webservice"
// 	"os"
// 	"text/tabwriter"
// 	"time"
// )
// func cmdList(c *cli.Context) {
// 	var reports []Report

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get("/v1/settings/reports")
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &reports)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "REPORT ID\tYEAR\tMONTH\tSTART TIME\tEND TIME\tSERVER SECONDS\tCLOSED\r")

// 	for _, report := range reports {
// 		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%g\t%t\n", report.Id, report.Year, report.Month, report.Start_time, report.End_time, report.Server_seconds, report.Closed)
// 	}

// 	w.Flush()
// }

// func cmdShow(c *cli.Context) {
// 	var report Report

// 	utils.FlagsRequired(c, []string{"id"})

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/settings/reports/%s", c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &report)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)

// 	fmt.Fprintln(w, "REPORT ID\tYEAR\tMONTH\tSTART TIME\tEND TIME\tSERVER SECONDS\tCLOSED\r")
// 	fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%g\t%t\n", report.Id, report.Year, report.Month, report.Start_time, report.End_time, report.Server_seconds, report.Closed)
// 	fmt.Fprintln(w, "LINES:\r")
// 	fmt.Fprintln(w, "ID\tCOMMISSIONED AT\tDECOMMISSIONED AT\tINSTANCE ID\tINSTANCE NAME\tINSTANCE FQDN\tCONSUMPTION\r")

// 	for _, l := range report.Li {
// 		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%s\t%g\n", l.Id, l.Commissioned_at, l.Decommissioned_at, l.Instance_id, l.Instance_name, l.Instance_fqdn, l.Consumption)
// 	}
// 	w.Flush()

// }
