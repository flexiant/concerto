/*


	Licensee reports

	Reports allow the user to have information about the historical uptime of their servers.
  	** The Licensee will have visibility for all the servers in the system.

	The available commands are:
		list	list of reports for all the servers in the system (licensee only)
		show	details about a particular report of a server (licensee only)

	Use "licensee_reports --help" on the commandline interface for more information about the available subcommands

	List licensee reports

	Returns a list of reports for all the servers in the system. The authenticated user must be the licensee.

	Usage:

		licensee_reports list

	Show licensee report

	Returns details about a particular report associated to any account group of the tenant. The authenticated user must be the licensee.

	Usage:

		licensee_reports show --id <report_id>

*/
package licensee

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
	"time"
)

type LicenseeReport struct {
	Id            string     `json:"id"`
	Year          int        `json:"year"`
	Month         time.Month `json:"month"`
	StartTime     time.Time  `json:"start_time"`
	EndTime       time.Time  `json:"end_time"`
	ServerSeconds float32    `json:"server_seconds"`
	Closed        bool       `json:"closed"`
	Li            []Lines    `json:"lines"`
}

type Lines struct {
	Id                string    `json:"_id"`
	Commissioned_at   time.Time `json:"commissioned_at"`
	Decommissioned_at time.Time `json:"decommissioned_at"`
	Instance_id       string    `json:"instance_id"`
	Instance_name     string    `json:"instance_name"`
	Instance_fqdn     string    `json:"instance_fqdn"`
	Consumption       float32   `json:"consumption"`
}

func cmdList(c *cli.Context) {
	var licenseeReports []LicenseeReport

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/licensee/reports")
	utils.CheckError(err)

	err = json.Unmarshal(data, &licenseeReports)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tYEAR\tMONTH\tSTART TIME\tEND TIME\tSERVER SECONDS\tCLOSED\r")

	for _, lr := range licenseeReports {
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%g\t%t\n", lr.Id, lr.Year, lr.Month, lr.StartTime, lr.EndTime, lr.ServerSeconds, lr.Closed)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	var vals LicenseeReport

	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/licensee/reports/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &vals)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)

	fmt.Fprintln(w, "REPORT ID\tYEAR\tMONTH\tSTART TIME\tEND TIME\tSERVER SECONDS\tCLOSED\r")
	fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%g\t%t\n", vals.Id, vals.Year, vals.Month, vals.StartTime, vals.EndTime, vals.ServerSeconds, vals.Closed)

	fmt.Fprintln(w, "LINES:\r")
	fmt.Fprintln(w, "ID\tCOMMISSIONED AT\tDECOMMISSIONED AT\tINSTANCE ID\tINSTANCE NAME\tINSTANCE FQDN\tCONSUMPTION\r")

	for _, l := range vals.Li {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%g\n", l.Id, l.Commissioned_at, l.Decommissioned_at, l.Instance_id, l.Instance_name, l.Instance_fqdn, l.Consumption)
	}
	w.Flush()

}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Returns a list of reports for all the servers in the system",
			Action: cmdList,
		},
		{
			Name:   "show",
			Usage:  "Returns details about a particular report associated to any account group of the tenant.",
			Action: cmdShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Report id",
				},
			},
		},
	}
}
