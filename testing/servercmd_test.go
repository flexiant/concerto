package testing

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/dispatcher"
	"github.com/flexiant/concerto/firewall"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var fleetName = "test-" + fmt.Sprint(time.Now().Unix())
var fleetId string
var gruId string
var minionId string
var plan = "55266f9411305a957d000127"
var domain_id = "55266f8611305a957d000018"

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
}

func TestBlob(t *testing.T) {
	println("Testing started")
}

func Test_Firewall_Apply(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "firewall", "apply"})
}

func Test_Firewall_Flush(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "firewall", "flush"})
}

func Test_Firewall_List(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "firewall", "list"})
}

func Test_Scripts_Boot(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "scripts", "boot"})
}

func Test_Scripts_Operational(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "scripts", "operational"})
}

func Test_Scripts_Shutdown(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "scripts", "shutdown"})
}
