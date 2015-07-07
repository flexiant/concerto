package generic_images

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

type GenericImage struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func cmdList(c *cli.Context) {
	var images []GenericImage

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/cloud/generic_images")
	utils.CheckError(err)

	err = json.Unmarshal(data, &images)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\r")

	for _, image := range images {
		fmt.Fprintf(w, "%s\t%s\n", image.Id, image.Name)
	}

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "This action lists the available generic images.",
			Action: cmdList,
		},
	}
}
