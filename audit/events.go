// /*
// 	Events allow the user to track their actions and the state of their servers.

// 	The available commands are:
// 		list	events related to the account group
// 		list_system_events	system-wide events

// 	Use "audit events --help" on the commandline interface for more information about the available subcommands.

// 	Events list

// 	Returns information about the events related to the account group.

// 	Usage:

// 		events list

// 	Events list_system_events

// 	Returns information about system-wide events.

// 	Usage:

// 		events list_system_events

// */
package audit

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/codegangsta/cli"
// 	"github.com/flexiant/concerto/utils"
// 	"github.com/flexiant/concerto/webservice"
// 	"os"
// 	"text/tabwriter"
// 	"time"
// )

// // Event stores an Concerto event item
// type Event struct {
// 	Id          string    `json:"id"`
// 	Timestamp   time.Time `json:"timestamp"`
// 	Level       string    `json:"level"`
// 	Header      string    `json:"header"`
// 	Description string    `json:"description"`
// }

// func cmdListEvents(c *cli.Context) {
// 	var events []Event

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get("/v1/audit/events")
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &events)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tTIMESTAMP\tLEVEL\tHEADER\tDESCRIPTION\r")

// 	for _, event := range events {
// 		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", event.Id, event.Timestamp, event.Level, event.Header, event.Description)
// 	}

// 	w.Flush()
// }

// func cmdListSysEvents(c *cli.Context) {
// 	var events []Event

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get("/v1/audit/system_events")
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &events)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tTIMESTAMP\tLEVEL\tHEADER\tDESCRIPTION\r")

// 	for _, event := range events {
// 		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", event.Id, event.Timestamp, event.Level, event.Header, event.Description)
// 	}

// 	w.Flush()
// }
// a
