package apps

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Apps.",
			Action: cmd.AppList,
		},
		{
			Name:   "deploy",
			Usage:  "Deploys the App with the given id as a server on the cloud.",
			Action: cmd.AppDeploy,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "App Id",
				},
				cli.StringFlag{
					Name:  "location_id",
					Usage: "Identifier of the Location on which the App will be deployed",
				},
				cli.StringFlag{
					Name:  "cloud_provider_id",
					Usage: "Identifier of the Cloud Provider on which the App will be deployed",
				},
				cli.StringFlag{
					Name:  "server_plan_id",
					Usage: "Identifier of the Server Plan on which the App will be deployed",
				},
				cli.StringFlag{
					Name:  "hostname",
					Usage: "A hostname for the cloud server to deploy",
				},
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Identifier of the Domain under which the App will be deployed",
				},
			},
		},
	}
}
