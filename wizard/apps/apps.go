/*
	Apps are predefined software stacks ready to be deployed on cloud providers.

	The available commands are:
		list	list all available apps
		deploy	deploys the App with the given id as a server on the cloud

	Use "wizard apps --help" on the commandline interface for more information about the available subcommands.

	Apps list

	Lists the available Apps.

	Usage:

		apps list

	Deploy App

	Deploys the App with the given id as a server on the cloud.

	Usage:

		apps deploy (options)

	Options:
		--id <app_id>	App Id
		--location_id	<location_id>	Identifier of the Location on which the App will be deployed
		--cloud_provider_id	<cloud_provider_id>	Identifier of the Cloud Provider on which the App will be deployed
		--server_plan_id	<server_plan_id>	Identifier of the Server Plan on which the App will be deployed (optional)
		--hostname	<hostname>	A hostname for the cloud server to deploy
		--domain_id	<domain_id>	Identifier of the Domain under which the App will be deployed

*/
package apps

// import (
// 	"encoding/json"
// 	"fmt"
// 	log "github.com/Sirupsen/logrus"
// 	"github.com/codegangsta/cli"
// 	"github.com/flexiant/concerto/utils"
// 	"github.com/flexiant/concerto/webservice"
// 	"os"
// 	"text/tabwriter"
// )

// func cmdList(c *cli.Context) {
// 	var wizzApps []WizardApp

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get("/v1/wizard/apps")
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &wizzApps)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tFLAVOUR REQUIREMENTS\tGENERIC IMAGE ID\r")

// 	for _, wa := range wizzApps {
// 		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", wa.Id, wa.Name, wa.Flavour_requirements, wa.Generic_image_id)
// 	}

// 	w.Flush()
// }

// func cmdDeploy(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id", "location_id", "cloud_provider_id", "hostname", "domain_id"})
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	v := make(map[string]string)

// 	v["location_id"] = c.String("location_id")
// 	v["domain_id"] = c.String("domain_id")
// 	v["cloud_provider_id"] = c.String("cloud_provider_id")
// 	if c.IsSet("server_plan_id") {
// 		v["server_plan_id"] = c.String("server_plan_id")
// 	}
// 	v["hostname"] = c.String("hostname")

// 	jsonBytes, err := json.Marshal(v)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Post(fmt.Sprintf("/v1/wizard/apps/%s/deploy", c.String("id")), jsonBytes)
// 	if res == nil {
// 		log.Fatal(err)
// 	}
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)
// }
