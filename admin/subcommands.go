package admin

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Returns information about the reports related to all the account groups of the tenant. The authenticated user must be an admin.",
			Action: cmd.AdminReportList,
		},
		{
			Name:   "show",
			Usage:  "Returns details about a particular report associated to any account group of the tenant. The authenticated user must be an admin.",
			Action: cmdShow,
			// Action: cmd.AdminReportShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Report Identifier",
				},
			},
		},
	}
}
