package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/wizard"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpWizServerPlan prepares common resources to send request to Concerto API
func WireUpWizServerPlan(c *cli.Context) (ds *wizard.WizServerPlanService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = wizard.NewWizServerPlanService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up serverPlan service", err)
	}

	return ds, f
}

// WizServerPlanList subcommand function
func WizServerPlanList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	serverPlanSvc, formatter := WireUpWizServerPlan(c)

	checkRequiredFlags(c, []string{"app_id", "location_id", "cloud_provider_id"}, formatter)

	serverPlans, err := serverPlanSvc.GetWizServerPlanList(c.String("app_id"), c.String("location_id"), c.String("cloud_provider_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive serverPlan data", err)
	}
	if err = formatter.PrintList(serverPlans); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
