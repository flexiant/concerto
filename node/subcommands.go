package node

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "create",
			Usage:  "Creates a Node",
			Action: cmdCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cluster",
					Usage: "Cluster Name to Attach Node",
				},
				cli.StringFlag{
					Name:  "plan",
					Usage: "Server Plan to Use to Create Node",
				},
			},
		},
		{
			Name:   "list",
			Usage:  "Lists all available Nodes",
			Action: cmd.NodeList,
		},
		{
			Name:   "start",
			Usage:  "Starts a given Node",
			Action: cmd.NodeStart,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Node Id",
				},
			},
		},
		{
			Name:   "stop",
			Usage:  "Stops a given Node",
			Action: cmd.NodeStop,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Node Id",
				},
			},
		},
		{
			Name:   "restart",
			Usage:  "Restarts a given Node",
			Action: cmd.NodeRestart,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Node Id",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a given Node",
			Action: cmd.NodeDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Node Id",
				},
			},
		},
		{
			Name:   "docker",
			Usage:  "Docker command line wrapper",
			Action: cmdDockerHijack,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "node",
					Usage: "Node Name",
				},
			},
		},
	}
}
