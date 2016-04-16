package reports

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Returns information about the reports related to the account group.",
			Action: cmd.SettingsReportList,
		},
		{
			Name:   "show",
			Usage:  "Returns details about a particular report associated to the account group of the authenticated user.",
			Action: cmd.SettingsReportShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Report Identifier",
				},
			},
		},
	}
}
