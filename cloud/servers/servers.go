package servers

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	// "time"

	// log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
)

func cmdListDNS(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var dnsList []types.Dns

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, data, res := webservice.Get(fmt.Sprintf("/v1/cloud/servers/%s/records", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res, data)

	err = json.Unmarshal(data, &dnsList)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tCONTENT\tTYPE\tIS FQDN\tDOMAIN ID\r")

	for _, dns := range dnsList {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%t\t%s\n", dns.Id, dns.Name, dns.Content, dns.Type, dns.IsFQDN, dns.Domain_id)
	}

	w.Flush()
}

func cmdListEvents(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var events []types.Event

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, data, res := webservice.Get(fmt.Sprintf("/v1/cloud/servers/%s/events", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res, data)

	err = json.Unmarshal(data, &events)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTIMESTAMP\tLEVEL\tHEADER\tDESCRIPTION\r")

	for _, event := range events {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", event.Id, event.Timestamp, event.Level, event.Header, event.Description)
	}

	w.Flush()
}

func cmdListScripts(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var scripts []types.ScriptChar

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, data, res := webservice.Get(fmt.Sprintf("/v1/cloud/servers/%s/operational_scripts", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res, data)

	err = json.Unmarshal(data, &scripts)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTYPE\tPARAMETER VALUES\tTEMPLATE ID\tSCRIPT ID\r")

	for _, script := range scripts {
		fmt.Fprintf(w, "%s\t%s\t%t\t%s\t%t\n", script.Id, script.Type, script.Parameter_values, script.Template_id, script.Script_id)
	}

	w.Flush()
}

func cmdExecuteScript(c *cli.Context) {
	utils.FlagsRequired(c, []string{"server_id", "script_id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, res, code := webservice.Put(fmt.Sprintf("/v1/cloud/servers/%s/operational_scripts/%s/execute", c.String("server_id"), c.String("script_id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(code, res)

	var event types.Event
	err = json.Unmarshal(res, &event)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTIMESTAMP\tLEVEL\tHEADER\tDESCRIPTION\r")
	fmt.Fprintf(w, "%s\t%s\t%t\t%s\t%t\n", event.Id, event.Timestamp, event.Level, event.Header, event.Description)

	w.Flush()
}
