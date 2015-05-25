package main

import (
	"crypto/x509"
	"encoding/pem"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/admin"
	"github.com/flexiant/concerto/audit"
	"github.com/flexiant/concerto/blueprint_scripts"
	"github.com/flexiant/concerto/blueprint_services"
	"github.com/flexiant/concerto/blueprint_templates"
	"github.com/flexiant/concerto/cloud_generic_images"
	"github.com/flexiant/concerto/cloud_providers"
	"github.com/flexiant/concerto/cloud_saas_providers"
	"github.com/flexiant/concerto/cloud_server_plan"
	"github.com/flexiant/concerto/cloud_servers"
	"github.com/flexiant/concerto/cloud_ssh_profiles"
	"github.com/flexiant/concerto/cloud_workspaces"
	"github.com/flexiant/concerto/cluster"
	"github.com/flexiant/concerto/container"
	"github.com/flexiant/concerto/converge"
	"github.com/flexiant/concerto/dispatcher"
	"github.com/flexiant/concerto/firewall"
	"github.com/flexiant/concerto/fleet"
	"github.com/flexiant/concerto/ship"
	"github.com/flexiant/concerto/utils"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const VERSION = "0.1.0"

func initLogging(lvl log.Level) {
	log.SetOutput(os.Stderr)
	log.SetLevel(lvl)
}

var ServerCommands = []cli.Command{
	{
		Name:  "firewall",
		Usage: "Manages Firewall Policies within a Host",
		Subcommands: append(
			firewall.SubCommands(),
		),
	},
	{
		Name:  "scripts",
		Usage: "Manages Execution Scripts within a Host",
		Subcommands: append(
			dispatcher.SubCommands(),
		),
	},
	{
		Name:   "converge",
		Usage:  "Converges Host to original Blueprint",
		Action: converge.CmbConverge,
	},
}

var ClientCommands = []cli.Command{
	{
		Name:  "ship",
		Usage: "Manages Ships",
		Subcommands: append(
			ship.SubCommands(),
		),
	},
	{
		Name:  "fleet",
		Usage: "Manages a Fleet",
		Subcommands: append(
			fleet.SubCommands(),
		),
	},
	{
		Name:  "container",
		Usage: "Manages Containers in a Ship",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "ship",
				Usage: "Ship Name",
			},
		},
		Action: container.CmbHijack,
	},
	{
		Name:  "cluster",
		Usage: "Manages Cluster in a Fleet",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "fleet",
				Usage: "Fleet Name",
			},
		},
		Action: cluster.CmbHijack,
	},
	{
		Name:  "reports",
		Usage: "Provides historical uptime of servers",
		Subcommands: append(
			admin.SubCommands(),
		),
	},
	{
		Name:  "events",
		Usage: "Events allow the user to track their actions and the state of their servers",
		Subcommands: append(
			audit.SubCommands(),
		),
	},
	{
		Name:  "scripts",
		Usage: "Allow the user to manage the scripts they want to run on the servers",
		Subcommands: append(
			blueprint_scripts.SubCommands(),
		),
	},
	{
		Name:  "services",
		Usage: "Provides information on services",
		Subcommands: append(
			blueprint_services.SubCommands(),
		),
	},
	{
		Name:  "templates",
		Usage: "Provides information on templates",
		Subcommands: append(
			blueprint_templates.SubCommands(),
		),
	},
	{
		Name:  "workspaces",
		Usage: "Provides information on workspaces",
		Subcommands: append(
			cloud_workspaces.SubCommands(),
		),
	},
	{
		Name:  "servers",
		Usage: "Provides information on servers",
		Subcommands: append(
			cloud_servers.SubCommands(),
		),
	},
	{
		Name:  "generic_image",
		Usage: "Provides information on generic images",
		Subcommands: append(
			cloud_generic_images.SubCommands(),
		),
	},
	{
		Name:  "ssh_profile",
		Usage: "Provides information on SSH profiles",
		Subcommands: append(
			cloud_ssh_profiles.SubCommands(),
		),
	},
	{
		Name:  "cloud_providers",
		Usage: "Provides information on cloud providers",
		Subcommands: append(
			cloud_providers.SubCommands(),
		),
	},
	{
		Name:  "cloud_server_plan",
		Usage: "Provides information on server plans",
		Subcommands: append(
			cloud_server_plan.SubCommands(),
		),
	},
	{
		Name:  "saas_providers",
		Usage: "Provides information about SAAS providers",
		Subcommands: append(
			cloud_saas_providers.SubCommands(),
		),
	},
}

func cmdNotFound(c *cli.Context, command string) {
	log.Fatalf(
		"%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name,
		command,
		c.App.Name,
		c.App.Name,
	)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isUserCertificate(filename string) bool {
	if utils.Exists(filename) {
		data, err := ioutil.ReadFile(filename)
		checkError(err)
		block, _ := pem.Decode(data)

		cert, err := x509.ParseCertificate(block.Bytes)
		checkError(err)

		if len(cert.Subject.OrganizationalUnit) > 0 {
			if cert.Subject.OrganizationalUnit[0] == "Users" {
				return true

			}
		}
	}
	return false
}

func main() {

	for _, f := range os.Args {
		if f == "-D" || f == "--debug" || f == "-debug" {
			os.Setenv("DEBUG", "1")
			initLogging(log.DebugLevel)
		}
	}

	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Concerto Contributors"
	app.Email = "https://github.com/flexiant/concerto"

	app.CommandNotFound = cmdNotFound
	app.Usage = "Manages comunication between Host and Concerto Platform"
	app.Version = VERSION

	if isUserCertificate(filepath.Join(utils.GetConcertoDir(), "ssl", "cert.crt")) {
		app.Commands = ClientCommands
	} else {
		app.Commands = ServerCommands
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug mode",
		},
		cli.StringFlag{
			EnvVar: "CONCERTO_CA_CERT",
			Name:   "ca-cert",
			Usage:  "CA to verify remotes against",
			Value:  filepath.Join(utils.GetConcertoDir(), "ssl", "ca_cert.pem"),
		},
		cli.StringFlag{
			EnvVar: "CONCERTO_CLIENT_CERT",
			Name:   "client-cert",
			Usage:  "Client cert to use for Concerto",
			Value:  filepath.Join(utils.GetConcertoDir(), "ssl", "cert.crt"),
		},
		cli.StringFlag{
			EnvVar: "CONCERTO_CLIENT_KEY",
			Name:   "client-key",
			Usage:  "Private key used in client Concerto auth",
			Value:  filepath.Join(utils.GetConcertoDir(), "ssl", "/private/cert.key"),
		},
		cli.StringFlag{
			EnvVar: "CONCERTO_CONFIG",
			Name:   "concerto-config",
			Usage:  "Concerto Config File",
			Value:  filepath.Join(utils.GetConcertoDir(), "client.xml"),
		},
		cli.StringFlag{
			EnvVar: "CONCERTO_ENDPOINT",
			Name:   "concerto-endpoint",
			Usage:  "Concerto Endpoint",
			Value:  os.Getenv("CONCERTO_ENDPOINT"),
		},
	}

	app.Run(os.Args)
}
