package cloud_providers

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Cloud Providers",
			Action: cmd.WizCloudProviderList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "app_id",
					Usage: "Identifier of the App",
				},
				cli.StringFlag{
					Name:  "location_id",
					Usage: "Identifier of the Location",
				},
			},
		},
	}
}
