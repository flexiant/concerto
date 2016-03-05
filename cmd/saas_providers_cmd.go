package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpSaasProvider prepares common resources to send request to Concerto API
func WireUpSaasProvider(c *cli.Context) (cs *api.SaasProviderService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	cs, err = api.NewSaasProviderService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up saasProvider service", err)
	}

	return cs, f
}

// SaasProviderList subcommand function
func SaasProviderList(c *cli.Context) {
	debugCmdFuncInfo(c)
	saasProviderSvc, formatter := WireUpSaasProvider(c)

	saasProviders, err := saasProviderSvc.GetSaasProviderList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive saasProvider data", err)
	}
	if err = formatter.PrintList(saasProviders); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}
