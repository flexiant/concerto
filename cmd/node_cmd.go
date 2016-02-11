package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpNode prepares common resources to send request to Concerto API
func WireUpNode(c *cli.Context) (ns *api.NodeService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ns, err = api.NewNodeService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up node service", err)
	}

	return ns, f
}

// NodeList subcommand function
func NodeList(c *cli.Context) {
	debugCmdFuncInfo(c)
	nodeSvc, formatter := WireUpNode(c)

	nodes, err := nodeSvc.GetNodeList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive node data", err)
	}
	if err = formatter.PrintList(nodes); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// NodeCreate subcommand function
func NodeCreate(c *cli.Context) {
	debugCmdFuncInfo(c)
	nodeSvc, formatter := WireUpNode(c)

	checkRequiredFlags(c, []string{"cluster", "plan"}, formatter)
	node, err := nodeSvc.CreateNode(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create node", err)
	}
	if err = formatter.PrintItem(*node); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// NodeDelete subcommand function
func NodeDelete(c *cli.Context) {
	debugCmdFuncInfo(c)
	nodeSvc, formatter := WireUpNode(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := nodeSvc.DeleteNode(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete node", err)
	}
}

// NodeStart subcommand function
func NodeStart(c *cli.Context) {
	debugCmdFuncInfo(c)
	nodeSvc, formatter := WireUpNode(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := nodeSvc.StartNode(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't start node", err)
	}
}

// NodeStop subcommand function
func NodeStop(c *cli.Context) {
	debugCmdFuncInfo(c)
	nodeSvc, formatter := WireUpNode(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := nodeSvc.StopNode(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't stop node", err)
	}
}

// NodeRestart subcommand function
func NodeRestart(c *cli.Context) {
	debugCmdFuncInfo(c)
	nodeSvc, formatter := WireUpNode(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := nodeSvc.RestartNode(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't restart node", err)
	}
}
