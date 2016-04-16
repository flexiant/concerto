package server_plans

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Server Plans.",
			Action: cmd.WizServerPlanList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "app_id",
					Usage: "Identifier of the App",
				},
				cli.StringFlag{
					Name:  "location_id",
					Usage: "Identifier of the Location",
				},
				cli.StringFlag{
					Name:  "cloud_provider_id",
					Usage: "Identifier of the Cloud Provider",
				},
			},
		},
	}
}
