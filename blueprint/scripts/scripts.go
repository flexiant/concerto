/*
	Allow the user to manage the scripts they wants to run on the servers.

	The available commands are:
		list
		show
		create
		update
		delete

	Use "blueprint scripts --help" on the commandline interface for more information about the available subcommands.

	Scripts list

	Lists all available scripts.

	Usage:

		scripts list

	Scripts show

	Shows information about a specific script.

	Usage:

		scripts show (options)

	Options:
		--id <script_id> 		Script Id


	Scripts create

	Creates a new script to be used in the templates.

	Usage:

		scripts create (options)

	Options:
		--id <script_id> 		Script Id
		--name <name> 			Name of the script
		--description <description> 	Description of the script's purpose
		--code <code> 			The script's code
		--parameters <parameters>	The names of the script's parameters

	Scripts update

	Updates an existing script.

	Usage:

		scripts update (options)

	Options:
		--id <script_id> 		Script Id
		--name <name> 			Name of the script
		--description <description> 	Description of the script's purpose
		--code <code> 			The script's code
		--parameters <parameters>	The names of the script's parameters

	Scripts delete

	Deletes a script.

	Usage:

		scripts delete (options)

	Options:
		--id <script_id> 		Script Id

*/
package scripts

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

type Script struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Code        string   `json:"code"`
	Parameters  []string `json:"parameters"`
}

func cmdList(c *cli.Context) {
	var scripts []Script

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/blueprint/scripts")
	utils.CheckError(err)

	err = json.Unmarshal(data, &scripts)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tCODE\tPARAMETERS\r")

	for _, script := range scripts {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", script.Id, script.Name, script.Description, script.Code, script.Parameters)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var script Script

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/blueprint/scripts/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &script)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tCODE\tPARAMETERS\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", script.Id, script.Name, script.Description, script.Code, script.Parameters)

	w.Flush()
}

func cmdCreate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"name", "description", "code"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["name"] = c.String("name")
	v["description"] = c.String("description")
	v["code"] = c.String("code")
	if c.IsSet("parameters") {
		v["parameters"] = c.String("parameters")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post("/v1/blueprint/scripts", jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)
	var new_script Script
	err = json.Unmarshal(res, &new_script)
	utils.CheckError(err)
	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tCODE\tPARAMETERS\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", new_script.Id, new_script.Name, new_script.Description, new_script.Code, new_script.Parameters)

	w.Flush()

}

func cmdUpdate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)
	if c.IsSet("name") {
		v["name"] = c.String("name")
	}
	if c.IsSet("description") {
		v["description"] = c.String("description")
	}
	if c.IsSet("code") {
		v["code"] = c.String("code")
	}
	if c.IsSet("parameters") {
		v["parameters"] = c.String("parameters")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Put(fmt.Sprintf("/v1/blueprint/scripts/%s", c.String("id")), jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)
	var new_script Script
	err = json.Unmarshal(res, &new_script)
	utils.CheckError(err)
	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tCODE\tPARAMETERS\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", new_script.Id, new_script.Name, new_script.Description, new_script.Code, new_script.Parameters)

	w.Flush()
}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Delete(fmt.Sprintf("/v1/blueprint/scripts/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

	fmt.Println(res)
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available scripts",
			Action: cmdList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific script",
			Action: cmdShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new script to be used in the templates. ",
			Action: cmdCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the script",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the script's purpose ",
				},
				cli.StringFlag{
					Name:  "code",
					Usage: "The script's code",
				},
				cli.StringFlag{
					Name:  "parameters",
					Usage: "The names of the script's parameters",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing script",
			Action: cmdUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the script",
				},
				cli.StringFlag{
					Name:  "description",
					Usage: "Description of the script's purpose ",
				},
				cli.StringFlag{
					Name:  "code",
					Usage: "The script's code",
				},
				cli.StringFlag{
					Name:  "parameters",
					Usage: "The names of the script's parameters",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a script",
			Action: cmdDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
			},
		},
	}
}
