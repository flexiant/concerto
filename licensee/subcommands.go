package licensee

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Returns a list of reports for all the servers in the system",
			Action: cmd.LicenseeReportList,
		},
		{
			Name:   "show",
			Usage:  "Returns details about a particular report associated to any account group of the tenant.",
			Action: cmd.LicenseeReportShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Report id",
				},
			},
		},
	}
}
