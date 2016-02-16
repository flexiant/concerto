package templates

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available templates",
			Action: cmd.TemplateList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific template",
			Action: cmd.TemplateShow,
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
			Action: cmd.TemplateCreate,
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
					Usage: "A list of space separated service recipes that is run on the servers at start-up",
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
			Action: cmd.TemplateUpdate,
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
			Action: cmd.TemplateDelete,
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
			Action: cmd.TemplateScriptList,
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
			Action: cmd.TemplateScriptShow,
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
			Action: cmd.TemplateScriptCreate,
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
			Action: cmd.TemplateScriptUpdate,
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
					Name:  "id",
					Usage: "Identifier for the template-script that is parameterised by the script characterisation",
				},
				cli.StringFlag{
					Name:  "parameter_values",
					Usage: "A map that assigns a value to each script parameter",
				},
			},
		},
		{
			Name:   "reorder_template_scripts",
			Usage:  "Reorders the scripts of the template and type specified according to the provided order, changing their execution order as corresponds.",
			Action: cmdReorderTemplateScripts,
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
					Name:  "script_ids",
					Usage: "An array that must contain all the ids of scripts of the given template and type in the desired execution order",
				},
			},
		},
		{
			Name:   "delete_template_script",
			Usage:  "Removes a parametrized script from a template",
			Action: cmd.TemplateScriptDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "template_id",
					Usage: "Template Id",
				},
				cli.StringFlag{
					Name:  "id",
					Usage: "Identifier for the template-script that is parameterised by the script characterisation",
				},
			},
		},
		{
			Name:  "list_template_servers",
			Usage: "Returns information about the servers that use a specific template. ",
			// Action: cmd.TemplateServersList,
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
