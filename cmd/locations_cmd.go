package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/wizard"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpLocation prepares common resources to send request to Concerto API
func WireUpLocation(c *cli.Context) (ds *wizard.LocationService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = wizard.NewLocationService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up location service", err)
	}

	return ds, f
}

// LocationList subcommand function
func LocationList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	locationSvc, formatter := WireUpLocation(c)

	locations, err := locationSvc.GetLocationList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive location data", err)
	}
	if err = formatter.PrintList(locations); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
