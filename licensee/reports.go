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
		fmt.Fprintf(w, "%s\t%s\t%d\t%g\t%d\t%s\t%s\n", lr.Id, lr.Year, lr.Month, lr.StartTime, lr.EndTime, lr.ServerSeconds, lr.Closed)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	// utils.FlagsRequired(c, []string{"id"})
	// var sp ServerPlan

	// webservice, err := webservice.NewWebService()
	// utils.CheckError(err)

	// data, err := webservice.Get(fmt.Sprintf("/v1/cloud/server_plans/%s", c.String("id")))
	// utils.CheckError(err)

	// err = json.Unmarshal(data, &sp)
	// utils.CheckError(err)

	// w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	// fmt.Fprintln(w, "ID\tNAME\tMEMORY\tCPUS\tSTORAGE\tLOCATION ID\tCLOUD PROVIDER ID\r")
	// fmt.Fprintf(w, "%s\t%s\t%d\t%g\t%d\t%s\t%s\n", sp.Id, sp.Name, sp.Memory, sp.CPUs, sp.Storage, sp.LocationId, sp.CloudProviderId)

	// w.Flush()
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
