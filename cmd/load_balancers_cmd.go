package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
)

// WireUpLoadBalancer prepares common resources to send request to Concerto API
func WireUpLoadBalancer(c *cli.Context) (ds *api.LoadBalancerService, f format.Formatter) {

	f = format.GetFormatter()

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err = api.NewLoadBalancerService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up loadBalancer service", err)
	}

	return ds, f
}

// LoadBalancerList subcommand function
func LoadBalancerList(c *cli.Context) {
	debugCmdFuncInfo(c)
	loadBalancerSvc, formatter := WireUpLoadBalancer(c)

	loadBalancers, err := loadBalancerSvc.GetLoadBalancerList()
	if err != nil {
		formatter.PrintFatal("Couldn't receive loadBalancer data", err)
	}
	if err = formatter.PrintList(loadBalancers); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// LoadBalancerShow subcommand function
func LoadBalancerShow(c *cli.Context) {
	debugCmdFuncInfo(c)
	loadBalancerSvc, formatter := WireUpLoadBalancer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	loadBalancer, err := loadBalancerSvc.GetLoadBalancer(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive loadBalancer data", err)
	}
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// LoadBalancerCreate subcommand function
func LoadBalancerCreate(c *cli.Context) {
	debugCmdFuncInfo(c)
	loadBalancerSvc, formatter := WireUpLoadBalancer(c)

	checkRequiredFlags(c, []string{"protocol"}, formatter)
	switch c.String("protocol") {
	case "http":
		checkRequiredFlags(c, []string{"name", "fqdn", "protocol", "domain_id", "cloud_provider_id"}, formatter)
	case "https":
		checkRequiredFlags(c, []string{"name", "fqdn", "protocol", "domain_id", "cloud_provider_id", "ssl_certificate", "ssl_certificate_private_key"}, formatter)
	}

	loadBalancer, err := loadBalancerSvc.CreateLoadBalancer(utils.FlagConvertParams(c))
	if err != nil {
		formatter.PrintFatal("Couldn't create loadBalancer", err)
	}
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// LoadBalancerUpdate subcommand function
func LoadBalancerUpdate(c *cli.Context) {
	debugCmdFuncInfo(c)
	loadBalancerSvc, formatter := WireUpLoadBalancer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	loadBalancer, err := loadBalancerSvc.UpdateLoadBalancer(utils.FlagConvertParams(c), c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't update loadBalancer", err)
	}
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// LoadBalancerDelete subcommand function
func LoadBalancerDelete(c *cli.Context) {
	debugCmdFuncInfo(c)
	loadBalancerSvc, formatter := WireUpLoadBalancer(c)

	checkRequiredFlags(c, []string{"id"}, formatter)
	err := loadBalancerSvc.DeleteLoadBalancer(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete loadBalancer", err)
	}
}

// LBNodeList subcommand function
func LBNodeList(c *cli.Context) {
	debugCmdFuncInfo(c)
	loadBalancerSvc, formatter := WireUpLoadBalancer(c)

	checkRequiredFlags(c, []string{"balancer_id"}, formatter)
	loadBalancerRecords, err := loadBalancerSvc.GetLBNodeList(c.String("balancer_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't list loadBalancer nodes", err)
	}
	if err = formatter.PrintList(*loadBalancerRecords); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// LBNodeCreate subcommand function
func LBNodeCreate(c *cli.Context) {
	debugCmdFuncInfo(c)
	loadBalancerSvc, formatter := WireUpLoadBalancer(c)

	checkRequiredFlags(c, []string{"balancer_id", "server_id", "port"}, formatter)
	loadBalancer, err := loadBalancerSvc.CreateLBNode(utils.FlagConvertParams(c), c.String("balancer_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't create loadBalancer node", err)
	}
	if err = formatter.PrintItem(*loadBalancer); err != nil {
		formatter.PrintFatal("Couldn't print/format result", err)
	}
}

// LBNodeDelete subcommand function
func LBNodeDelete(c *cli.Context) {
	debugCmdFuncInfo(c)
	loadBalancerSvc, formatter := WireUpLoadBalancer(c)

	checkRequiredFlags(c, []string{"balancer_id", "node_id"}, formatter)
	err := loadBalancerSvc.DeleteLBNode(c.String("balancer_id"), c.String("node_id"))
	if err != nil {
		formatter.PrintFatal("Couldn't delete loadBalancer node", err)
	}
}
