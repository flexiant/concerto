package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpApp prepares common resources to send request to Concerto API
func WireUpApp(c *cli.Context) (ds *api.AppService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = api.NewAppService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up app service", err)
	}

	return ds, f
}

// AppList subcommand function
func AppList(c *cli.Context) {
	debugCmdFuncInfo(c)
	appSvc, formatter := WireUpApp(c)

	apps, err := appSvc.GetAppList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive app data", err)
	}
	if err = formatter.PrintList(apps); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// AppDeploy subcommand function
func AppDeploy(c *cli.Context) {
	debugCmdFuncInfo(c)
	appSvc, formatter := WireUpApp(c)

	checkRequiredFlags(c, []string{"id", "location_id", "cloud_provider_id", "hostname", "domain_id"}, formatter)
	app, err := appSvc.DeployApp(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't deploy app", err)
	}
	if err = formatter.PrintItem(*app); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}
