package firewall_profiles

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all existing firewall profiles",
			Action: cmd.FirewallProfileList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about the firewall profile identified by the given id.",
			Action: cmd.FirewallProfileShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Firewall profile Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a a firewall profile with the given parameters.",
			Action: cmd.FirewallProfileCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the firewall profile",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the firewall profile",
				},
				cli.StringFlag{
					Name:  "rules",
					Usage: "Set of rules of the firewall profile",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates the firewall profile identified by the given id with the given parameters.",
			Action: cmd.FirewallProfileUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Firewall profile Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the firewall profile",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the firewall profile",
				},
				cli.StringFlag{
					Name:  "rules",
					Usage: "Set of rules of the firewall profile",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Destroy a firewall profile",
			Action: cmd.FirewallProfileDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Firewall profile Id",
				},
			},
		},
	}
}
