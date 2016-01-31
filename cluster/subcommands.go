package cluster

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available Clusters",
			Action: cmd.ClusterList,
		},
		{
			Name:   "start",
			Usage:  "Starts a given Cluster",
			Action: cmd.ClusterStart,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "stop",
			Usage:  "Stops a given Cluster",
			Action: cmd.ClusterStop,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "empty",
			Usage:  "Empties a given Cluster",
			Action: cmd.ClusterEmpty,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a Cluster",
			Action: cmd.ClusterCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Cluster Name",
				},
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Domain Id",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a given Cluster",
			Action: cmd.ClusterDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "kubectl",
			Usage:  "Kubectl command line wrapper",
			Action: cmdKubectlHijack,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cluster",
					Usage: "Cluster Name",
				},
			},
		},
	}
}
