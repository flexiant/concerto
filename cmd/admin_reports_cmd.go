package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/admin"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpReport prepares common resources to send request to Concerto API
func WireUpReport(c *cli.Context) (ns *admin.ReportService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ns, err = admin.NewReportService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up report service", err)
	}

	return ns, f
}

// ReportList subcommand function
func AdminReportList(c *cli.Context) {
	debugCmdFuncInfo(c)
	reportSvc, formatter := WireUpReport(c)

	reports, err := reportSvc.GetAdminReportList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive report data", err)
	}
	if err = formatter.PrintList(reports); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// ReportShow subcommand function
func AdminReportShow(c *cli.Context) {
	debugCmdFuncInfo(c)
	reportSvc, formatter := WireUpReport(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	report, err := reportSvc.GetAdminReport(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive report data", err)
	}
	if err = formatter.PrintList(report); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}
