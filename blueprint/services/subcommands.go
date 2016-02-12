package services

import (
	"github.com/codegangsta/cli"
	// "github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available services",
			Action: cmdList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific service",
			Action: cmdShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Service Id",
				},
			},
		},
	}
}
