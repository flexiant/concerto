package cloud_accounts

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the cloud accounts of the account group.",
			Action: cmd.CloudAccountList,
		},
		{
			Name:   "create",
			Usage:  "Creates a new cloud account.",
			Action: cmd.CloudAccountCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cloud_provider_id",
					Usage: "Identifier of the cloud provider",
				},
				cli.StringFlag{
					Name:  "credentials",
					Usage: "A mapping assigning a value to each of the required yes credentials of the cloud provider (JSON String)",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing cloud account.",
			Action: cmd.CloudAccountUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Account Id",
				},
				cli.StringFlag{
					Name:  "credentials",
					Usage: "A mapping assigning a value to each of the required yes credentials of the cloud provider (JSON String)",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a cloud account",
			Action: cmd.CloudAccountDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Account Id",
				},
			},
		},
	}
}
