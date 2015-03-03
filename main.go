package main

import (
	"os"
	"path"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/converge"
	"github.com/flexiant/concerto/dispatcher"
	"github.com/flexiant/concerto/firewall"
	"github.com/flexiant/concerto/utils"
)

const VERSION = "0.1.0"

func initLogging(lvl log.Level) {
	log.SetOutput(os.Stderr)
	log.SetLevel(lvl)
}

var Commands = []cli.Command{
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

func cmdNotFound(c *cli.Context, command string) {
	log.Fatalf(
		"%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name,
		command,
		c.App.Name,
		c.App.Name,
	)
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
	app.Commands = Commands
	app.CommandNotFound = cmdNotFound
	app.Usage = "Manages comunication between Host and Concerto Platform"
	app.Version = VERSION

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
	}

	app.Run(os.Args)
}
