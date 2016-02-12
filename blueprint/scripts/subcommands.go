package scripts

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available scripts",
			Action: cmd.ScriptsList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific script",
			Action: cmd.ScriptShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new script to be used in the templates. ",
			Action: cmd.ScriptCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the script",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the script's purpose ",
				},
				cli.StringFlag{
					Name:  "code",
					Usage: "The script's code",
				},
				cli.StringFlag{
					Name:  "parameters",
					Usage: "The names of the script's parameters",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing script",
			Action: cmd.ScriptUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the script",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the script's purpose ",
				},
				cli.StringFlag{
					Name:  "code",
					Usage: "The script's code",
				},
				cli.StringFlag{
					Name:  "parameters",
					Usage: "The names of the script's parameters",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a script",
			Action: cmd.ScriptDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
			},
		},
	}
}
