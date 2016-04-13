/*
	A cloud account allows the platform to commission and control servers on behalf of the user.
  	A cloud account stores the credentials needed to access a cloud provider.

	The available commands are:
		list
		create
		update
		delete

	Use "settings cloud_accounts --help" on the commandline interface for more information about the available subcommands.

	Cloud accounts list

	Lists the cloud accounts of the account group.

	Usage:

		cloud_accounts list

	Cloud account create

	Creates a new cloud account.

	Usage:

		cloud_accounts create (options)

	Options:
		--cloud_provider_id <cloud_provider_id> 	Identifier of the cloud provider
		--credentials <credentials> 	A mapping assigning a value to each of the required yes credentials of the cloud provider

	Cloud account update

	Updates an existing cloud account.

	Usage:

		cloud_accounts update (options)

	Options:
		--id <cloud_account_id> 		Identifier of the cloud account
		--credentials <credentials> 	A mapping assigning a value to each of the required yes credentials of the cloud provider

	Cloud account delete

	This action deletes a cloud account.

	Usage:

		cloud_accounts delete (options)

	Options:
		--id <cloud_account_id> 		Identifier of the cloud account
*/
package cloud_accounts

// import (
// 	"encoding/json"
// 	"fmt"
// 	log "github.com/Sirupsen/logrus"
// 	"github.com/codegangsta/cli"
// 	"github.com/flexiant/concerto/api/types"
// 	"github.com/flexiant/concerto/utils"
// 	"github.com/flexiant/concerto/webservice"
// 	"os"
// 	"text/tabwriter"
// )

// func cmdList(c *cli.Context) {
// 	var accounts []Account

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get("/v1/settings/cloud_accounts")
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &accounts)
// 	utils.CheckError(err)

// 	err, data, res = webservice.Get("/v1/cloud/cloud_providers")
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	var cloudProviders []types.CloudProvider
// 	err = json.Unmarshal(data, &cloudProviders)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tCLOUD PROVIDER ID\tNAME\r")

// 	for _, ac := range accounts {
// 		acName := ""
// 		for _, cp := range cloudProviders {
// 			if ac.CloudProvId == cp.Id {
// 				acName = cp.Name
// 				break
// 			}
// 		}

// 		fmt.Fprintf(w, "%s\t%s\t%s\n", ac.Id, ac.CloudProvId, acName)
// 	}

// 	w.Flush()
// }

// func cmdCreate(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"cloud_provider_id", "credentials"})
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	credentialsString := []byte(c.String("credentials"))

// 	var jsonCredentials RequiredCredentials
// 	err = json.Unmarshal(credentialsString, &jsonCredentials)

// 	v := make(map[string]interface{})
// 	v["cloud_provider_id"] = c.String("cloud_provider_id")
// 	v["credentials"] = jsonCredentials

// 	jsonBytes, err := json.Marshal(v)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Post("/v1/settings/cloud_accounts", jsonBytes)
// 	if res == nil {
// 		log.Fatal(err)
// 	}
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)

// }

// func cmdUpdate(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	v := make(map[string]interface{})

// 	if c.IsSet("credentials") {
// 		credentialsString := []byte(c.String("credentials"))
// 		var jsonCredentials RequiredCredentials
// 		err = json.Unmarshal(credentialsString, &jsonCredentials)
// 		v["credentials"] = jsonCredentials
// 	}

// 	jsonBytes, err := json.Marshal(v)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Put(fmt.Sprintf("/v1/settings/cloud_accounts/%s", c.String("id")), jsonBytes)

// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)
// }

// func cmdDelete(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, mesg, res := webservice.Delete(fmt.Sprintf("/v1/settings/cloud_accounts/%s", c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, mesg)
// }
