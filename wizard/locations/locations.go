package locations

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

type Location struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func cmdList(c *cli.Context) {
	var locs []Location

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/wizard/locations")
	utils.CheckError(err)

	err = json.Unmarshal(data, &locs)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\r")

	for _, loc := range locs {
		fmt.Fprintf(w, "%s\t%s\n", loc.Id, loc.Name)
	}

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Locations.",
			Action: cmdList,
		},
	}
}
