package saas_accounts

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the SaaS accounts of the account group.",
			Action: cmd.SaasAccountList,
		},
		{
			Name:   "create",
			Usage:  "Creates a new SaaS account.",
			Action: cmd.SaasAccountCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "saas_provider_id",
					Usage: "Identifier of the SaaS provider",
				},
				cli.StringFlag{
					Name:  "account_data",
					Usage: "A mapping assigning a value to each of the required account data of the SaaS provider (JSON String)",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing SaaS account.",
			Action: cmd.SaasAccountUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Account Id",
				},
				cli.StringFlag{
					Name:  "account_data",
					Usage: "A mapping assigning a value to each of the required account data of the SaaS provider (JSON String)",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a SaaS account",
			Action: cmd.SaasAccountDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Account Id",
				},
			},
		},
	}
}
