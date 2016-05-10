package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/cloud"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpGenericImage prepares common resources to send request to Concerto API
func WireUpGenericImage(c *cli.Context) (ns *cloud.GenericImageService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ns, err = cloud.NewGenericImageService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up genericImage service", err)
	}

	return ns, f
}

// GenericImageList subcommand function
func GenericImageList(c *cli.Context) error {
	debugCmdFuncInfo(c)
	genericImageSvc, formatter := WireUpGenericImage(c)

	genericImages, err := genericImageSvc.GetGenericImageList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive genericImage data", err)
	}
	if err = formatter.PrintList(genericImages); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
	return nil
}
