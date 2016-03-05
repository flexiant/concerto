package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpCloudProvider prepares common resources to send request to Concerto API
func WireUpCloudProvider(c *cli.Context) (cs *api.CloudProviderService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	cs, err = api.NewCloudProviderService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up cloudProvider service", err)
	}

	return cs, f
}

// CloudProviderList subcommand function
func CloudProviderList(c *cli.Context) {
	debugCmdFuncInfo(c)
	cloudProviderSvc, formatter := WireUpCloudProvider(c)

	cloudProviders, err := cloudProviderSvc.GetCloudProviderList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive cloudProvider data", err)
	}
	if err = formatter.PrintList(cloudProviders); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}
