package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
"github.com/flexiant/concerto/cluster"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/admin"
	"github.com/flexiant/concerto/audit"
	"github.com/flexiant/concerto/blueprint/scripts"
	"github.com/flexiant/concerto/blueprint/services"
	"github.com/flexiant/concerto/blueprint/templates"
	"github.com/flexiant/concerto/cloud/generic_images"
	"github.com/flexiant/concerto/cloud/providers"
	"github.com/flexiant/concerto/cloud/saas_providers"
	"github.com/flexiant/concerto/cloud/server_plan"
	"github.com/flexiant/concerto/cloud/servers"
	"github.com/flexiant/concerto/cloud/ssh_profiles"
	"github.com/flexiant/concerto/cloud/workspaces"
	"github.com/flexiant/concerto/converge"
	"github.com/flexiant/concerto/dispatcher"
	"github.com/flexiant/concerto/dns"
	"github.com/flexiant/concerto/firewall"
	"github.com/flexiant/concerto/licensee"
	"github.com/flexiant/concerto/network/firewall_profiles"
	"github.com/flexiant/concerto/network/load_balancers"
	"github.com/flexiant/concerto/node"
	"github.com/flexiant/concerto/settings/cloud_accounts"
	"github.com/flexiant/concerto/settings/reports"
	"github.com/flexiant/concerto/settings/saas_accounts"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/wizard/apps"
	"github.com/flexiant/concerto/wizard/cloud_providers"
	"github.com/flexiant/concerto/wizard/locations"
	"github.com/flexiant/concerto/wizard/server_plans"
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

var BlueprintCommands = []cli.Command{
	{
		Name:  "scripts",
		Usage: "Allow the user to manage the scripts they want to run on the servers",
		Subcommands: append(
			scripts.SubCommands(),
		),
	},
	{
		Name:  "services",
		Usage: "Provides information on services",
		Subcommands: append(
			services.SubCommands(),
		),
	},
	{
		Name:  "templates",
		Usage: "Provides information on templates",
		Subcommands: append(
			templates.SubCommands(),
		),
	},
}

var CloudCommands = []cli.Command{
	{
		Name:  "workspaces",
		Usage: "Provides information on workspaces",
		Subcommands: append(
			workspaces.SubCommands(),
		),
	},
	{
		Name:  "servers",
		Usage: "Provides information on servers",
		Subcommands: append(
			servers.SubCommands(),
		),
	},
	{
		Name:  "generic_images",
		Usage: "Provides information on generic images",
		Subcommands: append(
			generic_images.SubCommands(),
		),
	},
	{
		Name:  "ssh_profiles",
		Usage: "Provides information on SSH profiles",
		Subcommands: append(
			ssh_profiles.SubCommands(),
		),
	},
	{
		Name:  "cloud_providers",
		Usage: "Provides information on cloud providers",
		Subcommands: append(
			providers.SubCommands(),
		),
	},
	{
		Name:  "server_plans",
		Usage: "Provides information on server plans",
		Subcommands: append(
			server_plan.SubCommands(),
		),
	},
	{
		Name:  "saas_providers",
		Usage: "Provides information about SAAS providers",
		Subcommands: append(
			saas_providers.SubCommands(),
		),
	},
}

var NetCommands = []cli.Command{
	{
		Name:  "firewall_profiles",
		Usage: "Provides information about firewall profiles",
		Subcommands: append(
			firewall_profiles.SubCommands(),
		),
	},
	{
		Name:  "load_balancers",
		Usage: "Provides information about load balancers",
		Subcommands: append(
			load_balancers.SubCommands(),
		),
	},
}

var SettingsCommands = []cli.Command{
	{
		Name:  "cloud_accounts",
		Usage: "Provides information about cloud accounts",
		Subcommands: append(
			cloud_accounts.SubCommands(),
		),
	},
	{
		Name:  "reports",
		Usage: "Provides information about reports",
		Subcommands: append(
			reports.SubCommands(),
		),
	},
	{
		Name:  "saas_accounts",
		Usage: "Provides information about SaaS accounts",
		Subcommands: append(
			saas_accounts.SubCommands(),
		),
	},
}

var WizardCommands = []cli.Command{
	{
		Name:  "apps",
		Usage: "Provides information about apps",
		Subcommands: append(
			apps.SubCommands(),
		),
	},
	{
		Name:  "cloud_providers",
		Usage: "Provides information about cloud providers",
		Subcommands: append(
			cloud_providers.SubCommands(),
		),
	},
	{
		Name:  "locations",
		Usage: "Provides information about locations",
		Subcommands: append(
			locations.SubCommands(),
		),
	},
	{
		Name:  "server_plans",
		Usage: "Provides information about server plans",
		Subcommands: append(
			server_plans.SubCommands(),
		),
	},
}

var ClientCommands = []cli.Command{
	{
		Name:      "nodes",
		ShortName: "no",
		Usage:     "Manages Docker Nodes",
		Subcommands: append(
			node.SubCommands(),
		),
	},
	{
		Name:      "cluster",
		ShortName: "clu",
		Usage:     "Manages a Kubernetes Cluster",
		Subcommands: append(
			cluster.SubCommands(),
		),
	},
	// {
	// 	Name:      "cluster",
	// 	ShortName: "clu",
	// 	Usage:     "Manages Kubernetes Cluster",
	// 	Flags: []cli.Flag{
	// 		cli.StringFlag{
	// 			Name:  "cluster",
	// 			Usage: "Cluster Name",
	// 		},
	// 	},
	// 	Action: cluster.CmbHijack,
	// },
	{
		Name:      "reports",
		ShortName: "rep",
		Usage:     "Provides historical uptime of servers",
		Subcommands: append(
			admin.SubCommands(),
		),
	},
	{
		Name:      "events",
		ShortName: "ev",
		Usage:     "Events allow the user to track their actions and the state of their servers",
		Subcommands: append(
			audit.SubCommands(),
		),
	},

	{
		Name:      "blueprint",
		ShortName: "bl",
		Usage:     "Manages blueprint commands for scripts, services and templates",
		Subcommands: append(
			BlueprintCommands,
		),
	},
	{
		Name:      "cloud",
		ShortName: "clo",
		Usage:     "Manages cloud related commands for workspaces, servers, generic images, ssh profiles, cloud providers, server plans and Saas providers",
		Subcommands: append(
			CloudCommands,
		),
	},
	{
		Name:      "dns_domains",
		ShortName: "dns",
		Usage:     "Provides information about DNS records",
		Subcommands: append(
			dns.SubCommands(),
		),
	},
	{
		Name:      "licensee_reports",
		ShortName: "lic",
		Usage:     "Provides information about licensee reports",
		Subcommands: append(
			licensee.SubCommands(),
		),
	},
	{
		Name:      "network",
		ShortName: "net",
		Usage:     "Manages network related commands for firewall profiles and load balancers",
		Subcommands: append(
			NetCommands,
		),
	},
	{
		Name:      "settings",
		ShortName: "set",
		Usage:     "Provides settings for cloud and Saas accounts as well as reports",
		Subcommands: append(
			SettingsCommands,
		),
	},
	{
		Name:      "wizard",
		ShortName: "wiz",
		Usage:     "Manages wizard related commands for apps, locations, cloud providers, server plans",
		Subcommands: append(
			WizardCommands,
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
