package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpTemplate prepares common resources to send request to Concerto API
func WireUpTemplate(c *cli.Context) (ts *api.TemplateService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ts, err = api.NewTemplateService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up template service", err)
	}

	return ts, f
}

// TemplateList subcommand function
func TemplateList(c *cli.Context) {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	templates, err := templateSvc.GetTemplateList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive template data", err)
	}
	if err = formatter.PrintList(templates); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// TemplateShow subcommand function
func TemplateShow(c *cli.Context) {
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
}

// TemplateCreate subcommand function
func TemplateCreate(c *cli.Context) {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"name", "generic_image_id"}, formatter)
	template, err := templateSvc.CreateTemplate(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create template", err)
	}
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// TemplateUpdate subcommand function
func TemplateUpdate(c *cli.Context) {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	template, err := templateSvc.UpdateTemplate(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update template", err)
	}
	if err = formatter.PrintItem(*template); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// TemplateDelete subcommand function
func TemplateDelete(c *cli.Context) {
	debugCmdFuncInfo(c)
	templateSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := templateSvc.DeleteTemplate(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete template", err)
	}
}

// =========== Template Scripts =============

// TemplateScriptList subcommand function
func TemplateScriptList(c *cli.Context) {
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
}

// TemplateScriptShow subcommand function
func TemplateScriptShow(c *cli.Context) {
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
}

// TemplateScriptCreate subcommand function
func TemplateScriptCreate(c *cli.Context) {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template_id", "type", "parameter_values", "script_id"}, formatter)

	// parse json parameter values
	params := utils.FlagConvertParams(c)

	// parameter values is raw json.
	parameterValues, err := utils.JSONParam(c.String("parameter_values"))
	if err != nil {
		formatter.PrintFatal("parameter_values must be valid JSON", err)
	}

	(*params)["parameter_values"] = parameterValues

	templateScript, err := templateScriptSvc.CreateTemplateScript(params, c.String("template_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't create templateScript", err)
	}
	if err = formatter.PrintItem(*templateScript); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// TemplateScriptUpdate subcommand function
func TemplateScriptUpdate(c *cli.Context) {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id", "template_id"}, formatter)
	templateScript, err := templateScriptSvc.UpdateTemplateScript(utils.FlagConvertParams(c), c.String("template_id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update templateScript", err)
	}
	if err = formatter.PrintItem(*templateScript); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// TemplateScriptDelete subcommand function
func TemplateScriptDelete(c *cli.Context) {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"id", "template_id"}, formatter)
	err := templateScriptSvc.DeleteTemplateScript(c.String("template_id"), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete templateScript", err)
	}
}

// TemplateScriptReorder subcommand function
func TemplateScriptReorder(c *cli.Context) {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template_id", "type", "script_ids"}, formatter)
	templateScript, err := templateScriptSvc.ReorderTemplateScript(c.String("template_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't reorder templateScript", err)
	}
	if err = formatter.PrintItem(*templateScript); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// =========== Template Servers =============

// TemplateServers subcommand function
func TemplateServersList(c *cli.Context) {
	debugCmdFuncInfo(c)
	templateScriptSvc, formatter := WireUpTemplate(c)

	checkRequiredFlags(c, []string{"template_id"}, formatter)
	templateScripts, err := templateScriptSvc.GetTemplateServersList(c.String("template_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive templateServers data", err)
	}
	if err = formatter.PrintList(templateScripts); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}
