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

func get_state_fleet(id string, t *testing.T) string {
	values := strings.Fields(testing_setup([]string{"", "fleet", "list"}, t))
	// println("running again\n\n")
	for i := 18; i <= (len(values) - 9); i = i + 9 {

		// println(values[i+3])
		// println(values[i+2])
		if strings.TrimSpace(values[i+2]) == strings.TrimSpace(id) {
			// println("return " + values[i+3])
			return strings.TrimSpace(values[i+3])
		}
	}

	return ""
}

func get_state_ship(id string, t *testing.T) string {
	out := testing_setup([]string{"", "ship", "list"}, t)
	values := strings.Fields(out)

	for _, row := range values {
		values := strings.Split(row, "\t")
		if values[2] == id {
			return values[6]
		}
	}
	return ""
}

func TestBlob(t *testing.T) {
	println("Testing started")
}

func Test_Fleet_Create(t *testing.T) {
	out := testing_setup([]string{"", "fleet", "create", "--fleet", fleetName, "--domain_id=" + domain_id}, t)
	id := ""
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

func Test_Fleet_Stop(t *testing.T) {
	// println("Stopping fleet id " + fleetId)
	state := get_state_fleet(fleetId, t)
	for state != "operational" {
		time.Sleep(30 * time.Second)
		state = get_state_fleet(fleetId, t)
	}
	time.Sleep(120 * time.Second)
	testing_setup([]string{"", "fleet", "stop", "--id=" + fleetId}, t)
}

func Test_Fleet_Start(t *testing.T) {
	state := get_state_fleet(fleetId, t)
	if state == "operational" {
		t.Error("The previous fleet stop operation did not succeed; cannot continue testing")
	}
	for state != "inactive" {
		time.Sleep(30 * time.Second)
		state = get_state_fleet(fleetId, t)
	}
	time.Sleep(60 * time.Second)
	testing_setup([]string{"", "fleet", "start", "--id=" + fleetId}, t)
}

// func Test_Fleet_Attach_Net(t *testing.T) {
// 	state := get_state_ship(gruId, t)
// 	for state != "operational" {
// 		time.Sleep(30 * time.Second)
// 		state = get_state_fleet(fleetId, t)
// 	}
// 	testing_setup([]string{"", "fleet", "attach_net", "--id=" + fleetId}, t)
// }

func Test_Ship_Stop(t *testing.T) {
	state := get_state_ship(gruId, t)
	for state != "operational" {
		time.Sleep(30 * time.Second)
		state = get_state_fleet(fleetId, t)
	}
	testing_setup([]string{"", "ship", "stop", "--id=" + gruId}, t)
}

func Test_Ship_Start(t *testing.T) {
	state := get_state_ship(gruId, t)
	for state != "inactive" {
		time.Sleep(30 * time.Second)
		state = get_state_fleet(fleetId, t)
	}

	testing_setup([]string{"", "ship", "start", "--id=" + gruId}, t)
}

func Test_Ship_Restart(t *testing.T) {
	state := get_state_ship(minionId, t)
	for state != "operational" {
		time.Sleep(30 * time.Second)
		state = get_state_fleet(fleetId, t)
	}
	testing_setup([]string{"", "ship", "restart", "--id=" + minionId}, t)
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
