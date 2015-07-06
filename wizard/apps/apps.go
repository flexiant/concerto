package wizard_apps

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

// type FlavourRequirement struct {
// 	MinMem      int      `json:"min_men"`
// 	Cpus        int      `json:"cpus"`
// 	StorageVals []string `json:"storage_vals"`
// }

type WizardApp struct {
	Id                   string          `json:"id"`
	Name                 string          `json:"name"`
	Flavour_requirements json.RawMessage `json:"flavour_requirements"`
	Generic_image_id     string          `json:"generic_image_id"`
}

func cmdList(c *cli.Context) {
	var wizzApps []WizardApp

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/wizard/apps")
	utils.CheckError(err)

	err = json.Unmarshal(data, &wizzApps)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFLAVOUR REQUIREMENTS\tGENERIC IMAGE ID\r")

	for _, wa := range wizzApps {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", wa.Id, wa.Name, wa.Flavour_requirements, wa.Generic_image_id)
	}

	w.Flush()
}

func cmdDeploy(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id", "location_id", "cloud_provider_id", "hostname", "domain_id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["location_id"] = c.String("location_id")
	v["domain_id"] = c.String("domain_id")
	v["cloud_provider_id"] = c.String("cloud_provider_id")
	if c.IsSet("server_plan_id") {
		v["server_plan_id"] = c.String("server_plan_id")
	}
	v["hostname"] = c.String("hostname")

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post(fmt.Sprintf("/v1/wizard/apps/%s/deploy", c.String("id")), jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)

}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the available Apps.",
			Action: cmdList,
		},
		{
			Name:   "deploy",
			Usage:  "Deploys the App with the given id as a server on the cloud.",
			Action: cmdDeploy,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "App Id",
				},
				cli.StringFlag{
					Name:  "location_id",
					Usage: "Identifier of the Location on which the App will be deployed",
				},
				cli.StringFlag{
					Name:  "cloud_provider_id",
					Usage: "Identifier of the Cloud Provider on which the App will be deployed",
				},
				cli.StringFlag{
					Name:  "server_plan_id",
					Usage: "Identifier of the Server Plan on which the App will be deployed",
				},
				cli.StringFlag{
					Name:  "hostname",
					Usage: "A hostname for the cloud server to deploy",
				},
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Identifier of the Domain under which the App will be deployed",
				},
			},
		},
	}
}
