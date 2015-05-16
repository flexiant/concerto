package testing

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/fleet"
	"github.com/flexiant/concerto/ship"
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

var ClientCommands = []cli.Command{
	{
		Name:  "ship",
		Usage: "Manages Container Ships in Host",
		Subcommands: append(
			ship.SubCommands(),
		),
	},
	{
		Name:  "fleet",
		Usage: "Manages Container Fleets in Host",
		Subcommands: append(
			fleet.SubCommands(),
		),
	},
}

/*
The logical order of the tests:
1. Creates a Fleet
2. List the fleet
3. Create Gru
4. Create Minion
5. List the ships
6. Attach net to fleet
7. Start Fleet
8. Stop Fleet
9. Start Ship
10. Restart Ship
11. Stop
12. Delete ship
13. Empty fleet
14. Delete Fleet
*/

func TestBlob(t *testing.T) {
	println("Testing started")
}

func Test_Fleet_Create(t *testing.T) {

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "fleet", "create", "--fleet", fleetName, "--domain_id=" + domain_id})

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	fleetId = string(out)

}

func Test_Fleet_List(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "fleet", "list"})
}

func Test_Ship_Create_Gru(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "ship", "create", "--fleet=" + fleetName, "--plan=" + plan})

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	gruId = string(out)
}

func Test_Ship_Create_Minion(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "ship", "create", "--fleet=" + fleetName, "--plan=" + plan})

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	minionId = string(out)
}

func Test_Ship_List(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "ship", "list"})
}

func Test_Fleet_Attach_Net(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "fleet", "attach_net", "--id=" + fleetId})
}

func Test_Fleet_Start(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "fleet", "start", "--id=" + fleetId})
}

func Test_Fleet_Stop(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "fleet", "stop", "--id=" + fleetId})
}

func Test_Ship_Start(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "ship", "start", "--id=" + gruId})
}

func Test_Ship_Restart(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "ship", "restart", "--id=" + gruId})
}

func Test_Ship_Stop(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "ship", "stop", "--id=" + gruId})
}

func Test_Ship_Delete(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "ship", "delete", "--id=" + minionId})
}

func Test_Fleet_Empty(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "fleet", "empty", "--id=" + fleetId})
}

func Test_Fleet_Delete(t *testing.T) {
	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run([]string{"", "fleet", "delete", "--id=" + fleetId})
}
