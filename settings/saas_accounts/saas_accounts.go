/*
	A SaaS account stores the credentials needed to access a SaaS provider.
  	It allows the user to access the services of a SaaS provider through the platform.

	The available commands are:
		list
		create
		update
		delete

	Use "settings saas_accounts --help" on the commandline interface for more information about the available subcommands.

	SaaS accounts list

	Lists the SaaS accounts of the account group.

	Usage:

		saas_accounts list

	SaaS account create

	Creates a new SaaS account.

	Usage:

		saas_accounts create (options)

	Options:
		--saas_provider_id <saas_provider_id> 	Identifier of the saas provider
		--account_data <account_data> 	A mapping assigning a value to each of the required account data of the SaaS provider

	SaaS account update

	Updates an existing SaaS account.

	Usage:

		saas_accounts update (options)

	Options:
		--id <saas_account_id> 		Identifier of the SaaS account
		--saas_provider_id <saas_provider_id> 	Identifier of the saas provider

	SaaS account delete

	This action deletes a SaaS account.

	Usage:

		saas_accounts delete (options)

	Options:
		--id <saas_account_id> 		Identifier of the SaaS account
*/
package saas_accounts

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

type Account struct {
	Id         string `json:"id"`
	SaasProvId string `json:"saas_provider_id"`
}

type SaasRequiredCredentials interface{}

func cmdList(c *cli.Context) {
	var accounts []Account

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/settings/saas_accounts")
	utils.CheckError(err)

	err = json.Unmarshal(data, &accounts)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tSAAS PROVIDER ID\r")

	for _, ac := range accounts {
		fmt.Fprintf(w, "%s\t%s\n", ac.Id, ac.SaasProvId)
	}

	w.Flush()
}

func cmdCreate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"saas_provider_id", "account_data"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	credentialsString := []byte(c.String("account_data"))

	var jsonCredentials SaasRequiredCredentials
	err = json.Unmarshal(credentialsString, &jsonCredentials)

	v := make(map[string]interface{})

	v["saas_provider_id"] = c.String("saas_provider_id")
	v["account_data"] = jsonCredentials

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post("/v1/settings/saas_accounts", jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)

}

func cmdUpdate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	credentialsString := []byte(c.String("account_data"))

	var jsonCredentials SaasRequiredCredentials
	err = json.Unmarshal(credentialsString, &jsonCredentials)

	v := make(map[string]interface{})

	if c.IsSet("account_data") {
		v["account_data"] = jsonCredentials
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Put(fmt.Sprintf("/v1/settings/saas_accounts/%s", c.String("id")), jsonBytes)

	utils.CheckError(err)
	fmt.Println(res)
}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Delete(fmt.Sprintf("/v1/settings/saas_accounts/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

	fmt.Println(res)
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the SaaS accounts of the account group.",
			Action: cmdList,
		},
		{
			Name:   "create",
			Usage:  "Creates a new SaaS account.",
			Action: cmdCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "saas_provider_id",
					Usage: "Identifier of the SaaS provider",
				},
				cli.StringFlag{
					Name:  "account_data",
					Usage: "A mapping assigning a value to each of the required account data of the SaaS provider (JSON String)",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing SaaS account.",
			Action: cmdUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Account Id",
				},
				cli.StringFlag{
					Name:  "account_data",
					Usage: "A mapping assigning a value to each of the required account data of the SaaS provider (JSON String)",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a SaaS account",
			Action: cmdDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Account Id",
				},
			},
		},
	}
}
