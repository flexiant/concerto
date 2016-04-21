package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/licensee"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpLicenseeReport prepares common resources to send request to Concerto API
func WireUpLicenseeReport(c *cli.Context) (ns *licensee.LicenseeReportService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ns, err = licensee.NewLicenseeReportService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up report service", err)
	}

	return ns, f
}

// ReportList subcommand function
func LicenseeReportList(c *cli.Context) {
	debugCmdFuncInfo(c)
	reportSvc, formatter := WireUpLicenseeReport(c)

	reports, err := reportSvc.GetLicenseeReportList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive report data", err)
	}
	if err = formatter.PrintList(reports); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ReportShow subcommand function
func LicenseeReportShow(c *cli.Context) {
	debugCmdFuncInfo(c)
	reportSvc, formatter := WireUpLicenseeReport(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	report, err := reportSvc.GetLicenseeReport(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive report data", err)
	}
	if err = formatter.PrintItem(*report); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}
