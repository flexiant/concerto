package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/settings"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpSettingsReport prepares common resources to send request to Concerto API
func WireUpSettingsReport(c *cli.Context) (ns *settings.SettingsReportService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ns, err = settings.NewSettingsReportService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up report service", err)
	}

	return ns, f
}

// SettingsReportList subcommand function
func SettingsReportList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	reportSvc, formatter := WireUpSettingsReport(c)

	reports, err := reportSvc.GetSettingsReportList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive report data", err)
	}
	if err = formatter.PrintList(reports); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}

// SettingsReportShow subcommand function
func SettingsReportShow(c *cli.Context) error {
	debugCmdFuncInfo(c)
	reportSvc, formatter := WireUpSettingsReport(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	report, err := reportSvc.GetSettingsReport(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive report data", err)
	}
	if err = formatter.PrintList(report); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
