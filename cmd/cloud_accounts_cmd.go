package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/settings"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpCloudAccount prepares common resources to send request to Concerto API
func WireUpCloudAccount(c *cli.Context) (ds *settings.CloudAccountService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = settings.NewCloudAccountService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up cloudAccount service", err)
	}

	return ds, f
}

// CloudAccountList subcommand function
func CloudAccountList(c *cli.Context) {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpCloudAccount(c)

	cloudAccounts, err := cloudAccountSvc.GetCloudAccountList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudAccount data", err)
	}
	if err = formatter.PrintList(cloudAccounts); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// CloudAccountCreate subcommand function
func CloudAccountCreate(c *cli.Context) {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpCloudAccount(c)

	checkRequiredFlags(c, []string{"cloud_provider_id", "credentials"}, formatter)

	//cloudAccount, err := cloudAccountSvc.CreateCloudAccount(utils.FlagConvertParams(c))

	// parse json parameter values
	params, err := utils.FlagConvertParamsJSON(c, []string{"credentials"})
	if err != nil {
		formatter.PrintFatal("Error parsing parameters", err)
	}

	cloudAccount, err := cloudAccountSvc.CreateCloudAccount(params)

	if err != nil {
		formatter.PrintFatal("Couldn't create cloudAccount", err)
	}
	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// CloudAccountUpdate subcommand function
func CloudAccountUpdate(c *cli.Context) {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpCloudAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	cloudAccount, err := cloudAccountSvc.UpdateCloudAccount(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update cloudAccount", err)
	}
	if err = formatter.PrintItem(*cloudAccount); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// CloudAccountDelete subcommand function
func CloudAccountDelete(c *cli.Context) {
	debugCmdFuncInfo(c)
	cloudAccountSvc, formatter := WireUpCloudAccount(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := cloudAccountSvc.DeleteCloudAccount(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete cloudAccount", err)
	}
}
