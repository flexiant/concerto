/*
A template bundles the operating system to be run by a cloud server and the services and scripts to be applied to it, thus defining a blueprint for cloud server configuration management.

	The available commands are:
		list
		show
		create
		update
		delete
		list_template_scripts
		show_template_script
		create_template_script
		update_template_script
		reorder_template_scripts
		delete_template_script
		list_template_servers

	Use "blueprint templates --help" on the commandline interface for more information about the available subcommands.

	Templates list

	Lists all available templates.

	Usage:

		templates list

	Templates show

	Shows information about a specific template.

	Usage:

		templates show (options)

	Options:
		--id <template_id> 		Template Id


	Templates create

	Creates a new template to be used in the templates.

	Usage:

		templates create (options)

	Options:
		--id <template_id> 		Template Id
		--name <name> 			Name of the template
		--generic_image_id <generic_image_id> 	Identifier of the OS image that the template builds on
		--service_list <service_list> 			A list of service recipes that is run on the servers at start-up
		--configuration_attributes <configuration_attributes>	The attributes used to configure the services in the service_list

	Templates update

	Updates an existing template.

	Usage:

		templates update (options)

	Options:
		--id <template_id> 		Template Id
		--name <name> 			Name of the template
		--generic_image_id <generic_image_id> 	Identifier of the OS image that the template builds on
		--service_list <service_list> 			A list of service recipes that is run on the servers at start-up
		--configuration_attributes <configuration_attributes>	The attributes used to configure the services in the service_list

	Templates delete

	Deletes a template.

	Usage:

		templates delete (options)

	Options:
		--id <template_id> 		Template id



	List template scripts

	Shows the script characterisations, i.e., the parameterised scripts, of a template that either are run automatically during either boot, migration, or shutdown, or can be run in operational state.

	Usage:

		templates list_template_scripts (options)

	Options:
		--template_id <template_id>	Template id
		--type <type> Must be "operational", "boot", "migration", or "shutdown"


	Templates scripts show

	Shows information about a specific template.

	Usage:

		templates show_template_script (options)

	Options:
		--template_id <template_id> 		Template Id
		--script_id	<id>	Script Id

	Templates script create
	Creates a new script characterisation for a template and appends it to the list of script characterisations of the same type.

	Usage:

		templates create_template_script (options)

	Options:
		--template_id <template_id> 		Template Id
		--type <type> 			Must be "operational", "boot", "migration", or "shutdown"
		--script_id <script_id> 	Identifier for the script that is parameterised by the script characterisation
		--parameter_values <parameter_values> 			A map that assigns a value to each script parameter


	Templates script update

	Updates an existing script characterisation for a template.

	Usage:

		templates update_template_script (options)

	Options:
		--template_id <template_id> 		Template Id
		--script_id <script_id> 	Identifier for the script that is parameterised by the script characterisation
		--parameter_values <parameter_values> 			A map that assigns a value to each script parameter

	Template scripts reorder

	Reorders the scripts of the template and type specified according to the provided order, changing their execution order as corresponds.

	Usage:

		templates reorder_template_scripts (options)

	Options:
		--template_id <template_id> 		Template Id
		--type <type> 			Must be "operational", "boot", "migration", or "shutdown"
		--script_ids <[script_id1, script_id2, ...]> 	An array that must contain all the ids of scripts of the given template and type in the desired execution order


	Templates delete

	Removes a parametrized script from a template.

	Usage:

		templates delete (options)

	Options:
		--template_id <template_id> 		Template Id
		--script_id <script_id> 	Identifier for the script that is parameterised by the script characterisation

*/
package templates

// TODO DELETE THIS

// import (
// 	"encoding/json"
// 	"fmt"
// 	//	log "github.com/Sirupsen/logrus"
// 	"github.com/codegangsta/cli"
// 	"github.com/flexiant/concerto/api/types"
// 	"github.com/flexiant/concerto/utils"
// 	"github.com/flexiant/concerto/webservice"
// 	"os"
// 	"text/tabwriter"
// )

// func cmdListTemplateScripts(c *cli.Context) {
// 	var templateScripts []types.TemplateScript
// 	utils.FlagsRequired(c, []string{"template_id", "type"})
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/blueprint/templates/%s/scripts?type=%s", c.String("template_id"), c.String("type")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)
//
// 	err = json.Unmarshal(data, &templateScripts)
// 	utils.CheckError(err)
//
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tTYPE\tEXECUTION ORDER\tTEMPLATE ID\tSCRIPT ID\tPARAMETER VALUES\r")
//
// 	for _, templateScript := range templateScripts {
// 		fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%s\t%s\n", templateScript.ID, templateScript.Type, templateScript.ExecutionOrder, templateScript.TemplateID, templateScript.ScriptID, templateScript.ParameterValues)
// 	}
//
// 	w.Flush()
// }

// func cmdShowTemplateScript(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id", "template_id"})
// 	var templateScript types.TemplateScript
//
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s", c.String("template_id"), c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)
//
// 	err = json.Unmarshal(data, &templateScript)
// 	utils.CheckError(err)
//
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tTYPE\tEXECUTION ORDER\tTEMPLATE ID\tSCRIPT ID\tPARAMETER VALUES\r")
// 	fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%s\t%s\n", templateScript.ID, templateScript.Type, templateScript.ExecutionOrder, templateScript.TemplateID, templateScript.ScriptID, templateScript.ParameterValues)
//
// 	w.Flush()
// }

// func cmdCreateTemplateScript(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"template_id", "type", "parameter_values"})
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	v := make(map[string]interface{})
// 	var params types.TemplateScriptCredentials
//
// 	err = json.Unmarshal([]byte(c.String("parameter_values")), &params)
// 	v["script_id"] = c.String("script_id")
// 	v["type"] = c.String("type")
// 	v["parameter_values"] = params
//
// 	fmt.Printf("************************************script_id: %s\n", c.String("script_id"))
// 	jsonBytes, err := json.Marshal(v)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Post(fmt.Sprintf("/v1/blueprint/templates/%s/scripts", c.String("template_id")), jsonBytes)
// 	if res == nil {
// 		log.Fatal(err)
// 	}
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)
//
// 	var templateScript types.TemplateScript
// 	err = json.Unmarshal(res, &templateScript)
// 	utils.CheckError(err)
//
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tTYPE\tEXECUTION ORDER\tTEMPLATE ID\tSCRIPT ID\tPARAMETER VALUES\r")
// 	fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%s\t%s\n", templateScript.ID, templateScript.Type, templateScript.ExecutionOrder, templateScript.TemplateID, templateScript.ScriptID, templateScript.ParameterValues)
//
// 	w.Flush()
// }
//
// func cmdUpdateTemplateScript(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id", "template_id"})
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	v := make(map[string]interface{})
//
// 	if c.IsSet("parameter_values") {
// 		var params types.TemplateScriptCredentials
// 		err = json.Unmarshal([]byte(c.String("credentials")), &params)
// 		v["parameter_values"] = params
// 	}
//
// 	jsonBytes, err := json.Marshal(v)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Put(fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s", c.String("template_id"), c.String("id")), jsonBytes)
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)
//
// 	var templateScript types.TemplateScript
// 	err = json.Unmarshal(res, &templateScript)
// 	utils.CheckError(err)
//
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tTYPE\tEXECUTION ORDER\tTEMPLATE ID\tSCRIPT ID\tPARAMETER VALUES\r")
// 	fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%s\t%s\n", templateScript.ID, templateScript.Type, templateScript.ExecutionOrder, templateScript.TemplateID, templateScript.ScriptID, templateScript.ParameterValues)
//
// 	w.Flush()
// }
//
// func cmdDeleteTemplateScript(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id", "template_id"})
//
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	err, mesg, res := webservice.Delete(fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s", c.String("template_id"), c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, mesg)
// }
//
// func cmdReorderTemplateScripts(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"template_id", "type", "script_ids"})
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	v := make(map[string]interface{})
// 	v["type"] = c.String("type")
// 	v["script_ids"] = c.GlobalStringSlice("script_ids")
//
// 	jsonBytes, err := json.Marshal(v)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Put(fmt.Sprintf("/v1/blueprint/templates/%s/scripts/reorder", c.String("template_id")), jsonBytes)
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)
//
// 	var templateScripts []types.TemplateScript
// 	err = json.Unmarshal(res, &templateScripts)
// 	utils.CheckError(err)
//
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tTYPE\tEXECUTION ORDER\tTEMPLATE ID\tSCRIPT ID\tPARAMETER VALUES\r")
// 	for _, templateScript := range templateScripts {
// 		fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%s\t%s\n", templateScript.ID, templateScript.Type, templateScript.ExecutionOrder, templateScript.TemplateID, templateScript.ScriptID, templateScript.ParameterValues)
// 	}
// 	w.Flush()
// }
//
// func cmdListTemplateServers(c *cli.Context) {
// 	var templateServers []types.TemplateServer
//
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/blueprint/templates/%s/servers", c.String("template_id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)
//
// 	err = json.Unmarshal(data, &templateServers)
// 	utils.CheckError(err)
//
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tFQDN\tSTATE\tPUBLIC IP\tWORKSPACE ID\tTEMPLATE ID\tSERVER PLAN ID\tSSH PROFILE ID\r")
//
// 	for _, templateServer := range templateServers {
// 		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", templateServer.ID, templateServer.Name, templateServer.Fqdn, templateServer.State, templateServer.PublicIP, templateServer.WorkspaceID, templateServer.TemplateID, templateServer.ServerPlanID, templateServer.SSHProfileID)
// 	}
//
// 	w.Flush()
// }
