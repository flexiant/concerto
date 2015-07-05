package settings_reports

import (
	"encoding/json"
	"fmt"
	// log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
	"time"
)

type Report struct {
	Id             string     `json:"id"`
	Year           int        `json:"year"`
	Month          time.Month `json:"month"`
	Start_time     time.Time  `json:"start_time"`
	End_time       time.Time  `json:"end_time"`
	Server_seconds float32    `json:"server_seconds"`
	Closed         bool       `json:"closed"`
	Li             []Lines    `json:"lines"`
}

type Lines struct {
	Id                string    `json:"_id"`
	Commissioned_at   time.Time `json:"commissioned_at"`
	Decommissioned_at time.Time `json:"decommissioned_at"`
	Instance_id       string    `json:"instance_id"`
	Instance_name     string    `json:"instance_name"`
	Instance_fqdn     string    `json:"instance_fqdn"`
	Consumption       time.Time `json:"consumption"`
}

func cmdList(c *cli.Context) {
	var reports []Report

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/admin/reports")
	utils.CheckError(err)

	err = json.Unmarshal(data, &reports)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "REPORT ID\tYEAR\tMONTH\tSTART TIME\tEND TIME\tSERVER SECONDS\tCLOSED\r")

	for _, report := range reports {
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%g\t%t\n", report.Id, report.Year, report.Month, report.Start_time, report.End_time, report.Server_seconds, report.Closed)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	var report Report

	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/admin/reports/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &report)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)

	fmt.Fprintln(w, "REPORT ID\tYEAR\tMONTH\tSTART TIME\tEND TIME\tSERVER SECONDS\tCLOSED\r")
	fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%g\t%t\n", report.Id, report.Year, report.Month, report.Start_time, report.End_time, report.Server_seconds, report.Closed)
	fmt.Fprintln(w, "LINES:\r")
	fmt.Fprintln(w, "ID\tCOMMISSIONED AT\tDECOMMISSIONED AT\tINSTANCE ID\tINSTANCE NAME\tINSTANCE FQDN\tCONSUMPTION\r")

	for _, l := range report.Li {
		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%g\t%t\n", l.Id, l.Commissioned_at, l.Decommissioned_at, l.Instance_id, l.Instance_name, l.Instance_fqdn, l.Consumption)
	}
	w.Flush()

}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Returns information about the reports related to the account group.",
			Action: cmdList,
		},
		{
			Name:   "show_report",
			Usage:  "Returns details about a particular report associated to the account group of the authenticated user.",
			Action: cmdShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Report Identifier",
				},
			},
		},
	}
}
