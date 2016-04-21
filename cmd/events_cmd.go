package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/audit"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpEvent prepares common resources to send request to Concerto API
func WireUpEvent(c *cli.Context) (ns *audit.EventService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ns, err = audit.NewEventService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up event service", err)
	}

	return ns, f
}

// EventList subcommand function
func EventList(c *cli.Context) {
	debugCmdFuncInfo(c)
	eventSvc, formatter := WireUpEvent(c)

	events, err := eventSvc.GetEventList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive event data", err)
	}
	if err = formatter.PrintList(events); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// SysEventList subcommand function
func SysEventList(c *cli.Context) {
	debugCmdFuncInfo(c)
	eventSvc, formatter := WireUpEvent(c)

	events, err := eventSvc.GetSysEventList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive system event data", err)
	}
	if err = formatter.PrintList(events); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}
