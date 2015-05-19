package audit

import (
	"encoding/json"
	"fmt"
	// log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
	"time"
)

type Event struct {
	Id          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
	Level       string    `json:"level"`
	Header      string    `json:"header"`
	Description string    `json:"description"`
}

func cmdListEvents(c *cli.Context) {
	var events []Event

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/audit/events")
	utils.CheckError(err)

	err = json.Unmarshal(data, &events)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTIMESTAMP\tLEVEL\tHEADER\tDESCRIPTION\r")

	for _, event := range events {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", event.Id, event.Timestamp, event.Level, event.Header, event.Description)
	}

	w.Flush()
}

func cmdListSysEvents(c *cli.Context) {
	var events []Event

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/audit/system_events")
	utils.CheckError(err)

	err = json.Unmarshal(data, &events)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTIMESTAMP\tLEVEL\tHEADER\tDESCRIPTION\r")

	for _, event := range events {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", event.Id, event.Timestamp, event.Level, event.Header, event.Description)
	}

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list_events",
			Usage:  "Returns information about the events related to the account group.",
			Action: cmdListEvents,
		},
		{
			Name:   "list_system_events",
			Usage:  "Returns information about system-wide events.",
			Action: cmdListSysEvents,
		},
	}
}
