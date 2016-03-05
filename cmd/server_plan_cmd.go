package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpServerPlan prepares common resources to send request to Concerto API
func WireUpServerPlan(c *cli.Context) (ds *api.ServerPlanService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = api.NewServerPlanService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up serverPlan service", err)
	}

	return ds, f
}

// ServerPlanList subcommand function
func ServerPlanList(c *cli.Context) {
	debugCmdFuncInfo(c)
	serverPlanSvc, formatter := WireUpServerPlan(c)

	serverPlans, err := serverPlanSvc.GetServerPlanList(c.String("cloud_provider_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}
	if err = formatter.PrintList(serverPlans); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ServerPlanShow subcommand function
func ServerPlanShow(c *cli.Context) {
	debugCmdFuncInfo(c)
	serverPlanSvc, formatter := WireUpServerPlan(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	serverPlan, err := serverPlanSvc.GetServerPlan(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}
	if err = formatter.PrintItem(*serverPlan); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}
