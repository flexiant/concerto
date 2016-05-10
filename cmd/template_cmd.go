package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/blueprint"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpTemplate prepares common resources to send request to Concerto API
func WireUpTemplate(c *cli.Context) (ts *blueprint.TemplateService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ts, err = blueprint.NewTemplateService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up template service", err)
	}

	return ts, f
}

// TemplateList subcommand function
func TemplateList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	templates, err := templateSvc.GetTemplateList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive template data", err)
	}
	if err = formatter.PrintList(templates); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateShow subcommand function
func TemplateShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	template, err := templateSvc.GetTemplate(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive template data", err)
	}
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateCreate subcommand function
func TemplateCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"name", "generic_image_id"}, formatter)

	// parse json parameter values
	params, err := utils.FlagConvertParamsJSON(c, []string{"service_list", "configuration_attributes"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}

	template, err := templateSvc.CreateTemplate(params)
	if err != nil {
		formatter.PrintFatal("Couldn't create template", err)
	}
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateUpdate subcommand function
func TemplateUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)

	// parse json parameter values
	params, err := utils.FlagConvertParamsJSON(c, []string{"service_list", "configuration_attributes"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}

	template, err := templateSvc.UpdateTemplate(params, c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update template", err)
	}
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateDelete subcommand function
func TemplateDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := templateSvc.DeleteTemplate(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete template", err)
	}
	return nil
}

// =========== Template Scripts =============

// TemplateScriptList subcommand function
func TemplateScriptList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template_id", "type"}, formatter)
	templateScripts, err := templateScriptSvc.GetTemplateScriptList(c.String("template_id"), c.String("type"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive templateScript data", err)
	}
	if err = formatter.PrintList(*templateScripts); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateScriptShow subcommand function
func TemplateScriptShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id", "template_id"}, formatter)
	templateScript, err := templateScriptSvc.GetTemplateScript(c.String("template_id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive templateScript data", err)
	}
	if err = formatter.PrintItem(*templateScript); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateScriptCreate subcommand function
func TemplateScriptCreate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template_id", "type", "script_id"}, formatter)

	// parse json parameter values
	params, err := utils.FlagConvertParamsJSON(c, []string{"parameter_values"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}

	templateScript, err := templateScriptSvc.CreateTemplateScript(params, c.String("template_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't create templateScript", err)
	}
	if err = formatter.PrintItem(*templateScript); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateScriptUpdate subcommand function
func TemplateScriptUpdate(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	// TODO si necessary: type script_id parameter_values ?
	checkRequiredFlags(c, []string{"id", "template_id"}, formatter)

	// parse json parameter values
	params, err := utils.FlagConvertParamsJSON(c, []string{"parameter_values"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}

	templateScript, err := templateScriptSvc.UpdateTemplateScript(params, c.String("template_id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update templateScript", err)
	}
	if err = formatter.PrintItem(*templateScript); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// TemplateScriptDelete subcommand function
func TemplateScriptDelete(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id", "template_id"}, formatter)
	err := templateScriptSvc.DeleteTemplateScript(c.String("template_id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete templateScript", err)
	}
	return nil
}

// TemplateScriptReorder subcommand function
func TemplateScriptReorder(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template_id", "type", "script_ids"}, formatter)
	params, err := utils.FlagConvertParamsJSON(c, []string{"script_ids"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}

	templateScript, err := templateScriptSvc.ReorderTemplateScript(params, c.String("template_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't reorder templateScript", err)
	}
	if err = formatter.PrintList(*templateScript); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// =========== Template Servers =============

// TemplateServersList subcommand function
func TemplateServersList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template_id"}, formatter)
	templateServers, err := templateSvc.GetTemplateServerList(c.String("template_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive template servers data", err)
	}
	if err = formatter.PrintList(*templateServers); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
