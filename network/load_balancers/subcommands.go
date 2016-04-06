package load_balancers

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available load balancers",
			Action: cmd.LoadBalancerList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific load balancer",
			Action: cmd.LoadBalancerShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Load balancer Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new load balancer.",
			Action: cmd.LoadBalancerCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the load balancer",
				},
				cli.StringFlag{
					Name:  "fqdn",
					Usage: "Fully qualified domain name of the load balancer",
				},
				cli.StringFlag{
					Name:  "protocol",
					Usage: "Protocol of balanced traffic, either HTTP or HTTPS",
				},
				cli.StringFlag{
					Name:  "port",
					Usage: "Port where the load balancer listens for traffic ",
				},
				cli.StringFlag{
					Name:  "algorithm",
					Usage: "Algorithm used by the load balancer to balance incoming connections between servers. It can be either roundrobin, static-rr or leastconn.",
				},
				cli.StringFlag{
					Name:  "ssl_certificate",
					Usage: "SSL certificate to use, when protocol is HTTPS (SSL termination). ",
				},
				cli.StringFlag{
					Name:  "ssl_certificate_private_key",
					Usage: "Private key of SSL certificate to use, when protocol is HTTPS (SSL termination).",
				},
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Identifier of the DNS domain to which the FQDN of the load balancer belongs ",
				},
				cli.StringFlag{
					Name:  "cloud_provider_id",
					Usage: "Identifier of the cloud provider (that provides the load_balancer service) which shall deploy the load balancer",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing load balancer",
			Action: cmd.LoadBalancerUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Load balancer Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Logical name of the load balancer",
				},
				cli.StringFlag{
					Name:  "fqdn",
					Usage: "Fully qualified domain name of the load balancer",
				},
				cli.StringFlag{
					Name:  "protocol",
					Usage: "Protocol of balanced traffic, either HTTP or HTTPS",
				},
				cli.StringFlag{
					Name:  "port",
					Usage: "Port where the load balancer listens for traffic ",
				},
				cli.StringFlag{
					Name:  "algorithm",
					Usage: "Algorithm used by the load balancer to balance incoming connections between servers. It can be either roundrobin, static-rr or leastconn.",
				},
				cli.StringFlag{
					Name:  "ssl_certificate",
					Usage: "SSL certificate to use, when protocol is HTTPS (SSL termination). ",
				},
				cli.StringFlag{
					Name:  "ssl_certificate_private_key",
					Usage: "Private key of SSL certificate to use, when protocol is HTTPS (SSL termination).",
				},
			},
		},
		{
			Name:   "destroy",
			Usage:  "Destroys a load balancer",
			Action: cmd.LoadBalancerDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Load balancer Id",
				},
			},
		},
		{
			Name:   "list_balancer_nodes",
			Usage:  "This action provides information about the nodes of the load balancer identified by the given id.",
			Action: cmd.LBNodeList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "balancer_id",
					Usage: "Load balancer Id",
				},
			},
		},
		{
			Name:   "add_balancer_node",
			Usage:  "This action adds a node to the load balancer identified by the given id.",
			Action: cmd.LBNodeCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "balancer_id",
					Usage: "Load balancer Id",
				},
				cli.StringFlag{
					Name:  "server_id",
					Usage: "Identifier of the node's server",
				},
				cli.StringFlag{
					Name:  "port",
					Usage: "Port where the node listens for requests",
				},
			},
		},
		{
			Name:   "remove_balancer_node",
			Usage:  "This action removes the node identified by the given id from the load balancer identified by the given id. ",
			Action: cmd.LBNodeDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "balancer_id",
					Usage: "Load balancer Id",
				},
				cli.StringFlag{
					Name:  "node_id",
					Usage: "Identifier of the node",
				},
			},
		},
	}
}
