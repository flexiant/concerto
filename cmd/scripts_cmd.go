package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpScript prepares common resources to send request to Concerto API
func WireUpScript(c *cli.Context) (scs *api.ScriptService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	scs, err = api.NewScriptService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up script service", err)
	}

	return scs, f
}

// ScriptsList subcommand function
func ScriptsList(c *cli.Context) {
	debugCmdFuncInfo(c)
	scriptSvc, formatter := WireUpScript(c)

	scripts, err := scriptSvc.GetScriptList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive script data", err)
	}
	if err = formatter.PrintList(scripts); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ScriptShow subcommand function
func ScriptShow(c *cli.Context) {
	debugCmdFuncInfo(c)
	scriptSvc, formatter := WireUpScript(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	script, err := scriptSvc.GetScript(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive script data", err)
	}
	if err = formatter.PrintItem(*script); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ScriptCreate subcommand function
func ScriptCreate(c *cli.Context) {
	debugCmdFuncInfo(c)
	scriptSvc, formatter := WireUpScript(c)

	checkRequiredFlags(c, []string{"name", "description", "code"}, formatter)
	script, err := scriptSvc.CreateScript(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create script", err)
	}
	if err = formatter.PrintItem(*script); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ScriptUpdate subcommand function
func ScriptUpdate(c *cli.Context) {
	debugCmdFuncInfo(c)
	scriptSvc, formatter := WireUpScript(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	script, err := scriptSvc.UpdateScript(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update script", err)
	}
	if err = formatter.PrintItem(*script); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ScriptDelete subcommand function
func ScriptDelete(c *cli.Context) {
	debugCmdFuncInfo(c)
	scriptSvc, formatter := WireUpScript(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := scriptSvc.DeleteScript(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete script", err)
	}
}
