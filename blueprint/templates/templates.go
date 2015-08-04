/*

  A template bundles the operating system to be run by a cloud server and the services and scripts to be applied to it, thus defining a blueprint for cloud server configuration management.

*/
package templates

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

type Template struct {
	Id                      string          `json:"id"`
	Name                    string          `json:"name"`
	GenericImgId            string          `json:"generic_image_id"`
	ServiceList             []string        `json:"service_list"`
	ConfigurationAttributes json.RawMessage `json:"configuration_attributes"`
}

type TemplateScript struct {
	Id               string          `json:"id"`
	Type             string          `json:"type"`
	Template_Id      string          `json:"template_id"`
	Script_Id        string          `json:"script_id"`
	Parameter_Values json.RawMessage `json:"parameter_values"`
}

type TemplateServer struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Fqdn           string `json:"fqdn"`
	State          string `json:"state"`
	Public_ip      string `json:"public_ip"`
	Workspace_id   string `json:"workspace_id"`
	Template_id    string `json:"template_id"`
	Server_plan_id string `json:"server_plan_id"`
	Ssh_profile_id string `json:"ssh_profile_id"`
}

func cmdList(c *cli.Context) {
	var templates []Template

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/blueprint/templates")
	utils.CheckError(err)

	err = json.Unmarshal(data, &templates)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tGENERIC IMAGE ID\r")

	for _, template := range templates {
		fmt.Fprintf(w, "%s\t%s\t%s\n", template.Id, template.Name, template.GenericImgId)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var template Template

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/blueprint/templates/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &template)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tGENERIC IMAGE ID\tSERVICE LIST\tCONFIGURATION ATTRIBUTES\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", template.Id, template.Name, template.GenericImgId, template.ServiceList, template.ConfigurationAttributes)

	w.Flush()
}

func cmdCreate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"name", "generic_image_id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["name"] = c.String("name")
	v["generic_image_id"] = c.String("generic_image_id")
	if c.IsSet("service_list") {
		v["service_list"] = c.String("service_list")
	}
	if c.IsSet("configuration_attributes") {
		v["configuration_attributes"] = c.String("configuration_attributes")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post("/v1/blueprint/templates", jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)

	var template Template
	err = json.Unmarshal(res, &template)
	utils.CheckError(err)
	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tGENERIC IMAGE ID\tSERVICE LIST\tCONFIGURATION ATTRIBUTES\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", template.Id, template.Name, template.GenericImgId, template.ServiceList, template.ConfigurationAttributes)

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
	if c.IsSet("service_list") {
		v["service_list"] = c.String("service_list")
	}
	if c.IsSet("configuration_attributes") {
		v["configuration_attributes"] = c.String("configuration_attributes")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Put(fmt.Sprintf("/v1/blueprint/templates/%s", c.String("id")), jsonBytes)

	utils.CheckError(err)
	var template Template
	err = json.Unmarshal(res, &template)
	utils.CheckError(err)
	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tGENERIC IMAGE ID\tSERVICE LIST\tCONFIGURATION ATTRIBUTES\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", template.Id, template.Name, template.GenericImgId, template.ServiceList, template.ConfigurationAttributes)

	w.Flush()

}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Delete(fmt.Sprintf("/v1/blueprint/templates/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

	fmt.Println(res)
}

func cmdListTemplateScripts(c *cli.Context) {
	var templateScripts []TemplateScript
	utils.FlagsRequired(c, []string{"template_id", "type"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/blueprint/templates/%s/scripts?type=%s", c.String("template_id"), c.String("type")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &templateScripts)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTYPE\tTEMPLATE ID\tSCRIPT ID\tPARAMETER VALUES\r")

	for _, templateScript := range templateScripts {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", templateScript.Id, templateScript.Type, templateScript.Template_Id, templateScript.Script_Id, templateScript.Parameter_Values)
	}

	w.Flush()
}

func cmdShowTemplateScript(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id", "template_id"})
	var templateScript TemplateScript

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s", c.String("template_id"), c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &templateScript)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTYPE\tTEMPLATE ID\tSCRIPT ID\tPARAMETER VALUES\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", templateScript.Id, templateScript.Type, templateScript.Template_Id, templateScript.Script_Id, templateScript.Parameter_Values)

	w.Flush()
}

func cmdCreateTemplateScript(c *cli.Context) {
	utils.FlagsRequired(c, []string{"template_id", "type", "parameter_values"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["template_id"] = c.String("template_id")
	v["type"] = c.String("type")
	v["parameter_values"] = c.String("parameter_values")

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post(fmt.Sprintf("/v1/blueprint/templates/%s/scripts", c.String("template_id")), jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)

	var templateScript TemplateScript
	err = json.Unmarshal(res, &templateScript)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTYPE\tTEMPLATE ID\tSCRIPT ID\tPARAMETER VALUES\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", templateScript.Id, templateScript.Type, templateScript.Template_Id, templateScript.Script_Id, templateScript.Parameter_Values)

	w.Flush()
}

func cmdUpdateTemplateScript(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id", "template_id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	if c.IsSet("parameter_values") {
		v["parameter_values"] = c.String("parameter_values")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Put(fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s", c.String("template_id"), c.String("id")), jsonBytes)
	utils.CheckError(err)
	fmt.Println(res)

	var templateScript TemplateScript
	err = json.Unmarshal(res, &templateScript)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTYPE\tTEMPLATE ID\tSCRIPT ID\tPARAMETER VALUES\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", templateScript.Id, templateScript.Type, templateScript.Template_Id, templateScript.Script_Id, templateScript.Parameter_Values)

	w.Flush()
}

func cmdDeleteTemplateScript(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id", "template_id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Delete(fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s", c.String("template_id"), c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

	fmt.Println(res)
}

func cmdPlaceTemplateScript(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id", "template_id", "position"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)
	v["position"] = c.String("position")

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, _, res := webservice.Put(fmt.Sprintf("/v1/blueprint/templates/%s/scripts/%s/place", c.String("template_id"), c.String("id")), jsonBytes)
	utils.CheckError(err)
	fmt.Println(res)
}

func cmdListTemplateServers(c *cli.Context) {
	var templateServers []TemplateServer

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/blueprint/templates/%s/servers", c.String("template_id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &templateServers)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFQDN\tSTATE\tPUBLIC IP\tWORKSPACE ID\tTEMPLATE ID\tSERVER PLAN ID\tSSH PROFILE ID\r")

	for _, templateServer := range templateServers {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", templateServer.Id, templateServer.Name, templateServer.Fqdn, templateServer.State, templateServer.Public_ip, templateServer.Workspace_id, templateServer.Template_id, templateServer.Server_plan_id, templateServer.Ssh_profile_id)
	}

	w.Flush()
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available templates",
			Action: cmdList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific template",
			Action: cmdShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Template Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new template.",
			Action: cmdCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the template",
				},
				cli.StringFlag{
					Name:  "generic_image_id",
					Usage: "Identifier of the OS image that the template builds on",
				},
				cli.StringFlag{
					Name:  "service_list",
					Usage: "A list of service recipes that is run on the servers at start-up",
				},
				cli.StringFlag{
					Name:  "configuration_attributes",
					Usage: "The attributes used to configure the services in the service_list",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing template",
			Action: cmdUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the template",
				},
				cli.StringFlag{
					Name:  "service_list",
					Usage: "A list of service recipes that is run on the servers at start-up",
				},
				cli.StringFlag{
					Name:  "configuration_attributes",
					Usage: "The attributes used to configure the services in the service_list",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a template",
			Action: cmdDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Template Id",
				},
			},
		},
		{
			Name:   "list_template_scripts",
			Usage:  "Shows the script characterisations of a template",
			Action: cmdListTemplateScripts,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template_id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "type",
					Usage: "Must be \"operational\", \"boot\", \"migration\", or \"shutdown\"",
				},
			},
		},
		{
			Name:   "show_template_script",
			Usage:  "Shows information about a specific script characterisation",
			Action: cmdShowTemplateScript,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template_id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "id",
					Usage: "Script Id",
				},
			},
		},
		{
			Name:   "create_template_script",
			Usage:  "Creates a new script characterisation for a template and appends it to the list of script characterisations of the same type.",
			Action: cmdCreateTemplateScript,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template_id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "type",
					Usage: "Must be \"operational\", \"boot\", \"migration\", or \"shutdown\"",
				},
				cli.StringFlag{
					Name:  "script_id",
					Usage: "Identifier for the script that is parameterised by the script characterisation",
				},
				cli.StringFlag{
					Name:  "parameter_values",
					Usage: "A map that assigns a value to each script parameter",
				},
			},
		},
		{
			Name:   "update_template_script",
			Usage:  "Updates an existing script characterisation for a template.",
			Action: cmdUpdateTemplateScript,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template_id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "script_id",
					Usage: "Identifier for the script that is parameterised by the script characterisation",
				},
				cli.StringFlag{
					Name:  "parameter_values",
					Usage: "A map that assigns a value to each script parameter",
				},
			},
		},
		{
			Name:   "place_template_script",
			Usage:  "Changes the execution order of the scripts of a template: swaps positions of script_id with the script at \"position\"",
			Action: cmdPlaceTemplateScript,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template_id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "script_id",
					Usage: "Identifier for the script that is parameterised by the script characterisation",
				},
				cli.StringFlag{
					Name:  "position",
					Usage: "The target place for the script characterisation",
				},
			},
		},
		{
			Name:   "delete_template_script",
			Usage:  "Removes a parametrized script from a template",
			Action: cmdDeleteTemplateScript,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template_id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "script_id",
					Usage: "Identifier for the script that is parameterised by the script characterisation",
				},
			},
		},
		{
			Name:   "list_template_servers",
			Usage:  "Returns information about the servers that use a specific template. ",
			Action: cmdListTemplateServers,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template_id",
					Usage: "Template Id",
				},
			},
		},
	}
}
