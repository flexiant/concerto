package server_plan

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "This action lists the server plans offered by the cloud provider identified by the given id.",
			Action: cmd.ServerPlanList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cloud_provider_id",
					Usage: "Cloud provider id",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "This action shows information about the Server Plan identified by the given id.",
			Action: cmd.ServerPlanShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Server plan id",
				},
			},
		},
	}
}
