package testing

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/fleet"
	"github.com/flexiant/concerto/ship"
	"io/ioutil"
	"os"
	"strings"
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
	// {
	// 	Name:  "container",
	// 	Usage: "Manages Containers in a Ship",
	// 	Flags: []cli.Flag{
	// 		cli.StringFlag{
	// 			Name:  "ship",
	// 			Usage: "Ship Name",
	// 		},
	// 	},
	// 	Action: container.CmbHijack,
	// },
	// {
	// 	Name:  "cluster",
	// 	Usage: "Manages Cluster in a Fleet",
	// 	Flags: []cli.Flag{
	// 		cli.StringFlag{
	// 			Name:  "fleet",
	// 			Usage: "Fleet Name",
	// 		},
	// 	},
	// 	Action: cluster.CmbHijack,
	// },
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

func testing_setup(flags []string, t *testing.T) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	app := cli.NewApp()
	app.Commands = ClientCommands
	app.Run(flags)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	if strings.HasPrefix(string(out), "Incorrect Usage.") {
		t.Error("Incorrect Usage of Flags.")
		return ""
	}

	return string(out)
}

func TestBlob(t *testing.T) {
	println("Testing started")
}

func Test_Fleet_Create(t *testing.T) {
	out := testing_setup([]string{"", "fleet", "create", "--ft", fleetName, "--domain_id=" + domain_id}, t)
	if out != "" {
		fleetId = out
	}
}

func Test_Fleet_List(t *testing.T) {
	testing_setup([]string{"", "fleet", "list"}, t)
}

func Test_Ship_Create_Gru(t *testing.T) {
	out := testing_setup([]string{"", "ship", "create", "--fleet=" + fleetName, "--plan=" + plan}, t)

	if out != "" {
		gruId = out
	}
}

func Test_Ship_Create_Minion(t *testing.T) {
	out := testing_setup([]string{"", "ship", "create", "--fleet=" + fleetName, "--plan=" + plan}, t)
	if out != "" {
		minionId = out
	}
}

func Test_Ship_List(t *testing.T) {
	testing_setup([]string{"", "ship", "list"}, t)
}

func Test_Fleet_Attach_Net(t *testing.T) {
	testing_setup([]string{"", "fleet", "attach_net", "--id=" + fleetId}, t)
}

func Test_Fleet_Start(t *testing.T) {
	testing_setup([]string{"", "fleet", "start", "--id=" + fleetId}, t)
}

func Test_Fleet_Stop(t *testing.T) {
	testing_setup([]string{"", "fleet", "stop", "--id=" + fleetId}, t)
}

func Test_Ship_Start(t *testing.T) {
	testing_setup([]string{"", "ship", "start", "--id=" + gruId}, t)
}

func Test_Ship_Restart(t *testing.T) {
	testing_setup([]string{"", "ship", "restart", "--id=" + gruId}, t)
}

func Test_Ship_Stop(t *testing.T) {
	testing_setup([]string{"", "ship", "stop", "--id=" + gruId}, t)
}

func Test_Ship_Delete(t *testing.T) {
	testing_setup([]string{"", "ship", "delete", "--id=" + minionId}, t)
}

func Test_Fleet_Empty(t *testing.T) {
	testing_setup([]string{"", "fleet", "empty", "--id=" + fleetId}, t)
}

func Test_Fleet_Delete(t *testing.T) {
	testing_setup([]string{"", "fleet", "delete", "--id=" + fleetId}, t)

}

// func Test_Container() {
// testing_setup([]string{"", "container", "--ship"},t)
// }

// func Test_Cluster() {
// testing_setup([]string{"", "cluster", "--fleet"},t)
// }
