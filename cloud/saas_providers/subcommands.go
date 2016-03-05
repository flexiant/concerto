package saas_providers

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the SaaS providers supported by the platform.",
			Action: cmd.SaasProviderList,
		},
	}
}
