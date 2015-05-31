package network_load_balancers

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

type LoadBalancer struct {
	Id                          string `json:"id"`
	Name                        string `json:"name"`
	Fqdn                        string `json:"fqdn"`
	Protocol                    string `json:"protocol"`
	Port                        int    `json:"port"`
	Algorithm                   string `json:"algorithm"`
	SslCertificate              string `json:"ssl_certificate"`
	Ssl_certificate_private_key string `json:"ssl_certificate_private_key"`
	Domain_id                   string `json:"domain_id"`
	Cloud_provider_id           string `json:"cloud_provider_id"`
	Traffic_in                  int    `json:"traffic_in"`
	Traffic_out                 int    `json:"traffic_out"`
}

type Node struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	PublicIp string `json:"public_ip"`
	State    string `json:"state"`
	ServerId string `json:"server_id"`
	Port     int    `json:"port"`
}

func cmdList(c *cli.Context) {
	var loadBalancers []LoadBalancer

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/network/load_balancers")
	utils.CheckError(err)

	err = json.Unmarshal(data, &loadBalancers)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFQDN\tPROTOCOL\tPORT\tALGORITHM\tSSL CERTIFICATE\tSSL CERTIFICATE PRIVATE KEY\tDOMAIN ID\tCLOUD PROVIDER ID\tTRAFFIC IN\tTRAFFIC OUT\r")

	for _, lb := range loadBalancers {
		fmt.Fprintf(w, "%s\t%s\t%t\t%s\t%s\t%s\n", lb.Id, lb.Name, lb.Fqdn, lb.Protocol, lb.Port, lb.Algorithm, lb.SslCertificate, lb.Ssl_certificate_private_key, lb.Domain_id, lb.Cloud_provider_id, lb.Traffic_in, lb.Traffic_out)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var lb LoadBalancer

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/network/load_balancers/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &lb)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tFQDN\tPROTOCOL\tPORT\tALGORITHM\tSSL CERTIFICATE\tSSL CERTIFICATE PRIVATE KEY\tDOMAIN ID\tCLOUD PROVIDER ID\tTRAFFIC IN\tTRAFFIC OUT\r")
	fmt.Fprintf(w, "%s\t%s\t%t\t%s\t%s\t%s\n", lb.Id, lb.Name, lb.Fqdn, lb.Protocol, lb.Port, lb.Algorithm, lb.SslCertificate, lb.Ssl_certificate_private_key, lb.Domain_id, lb.Cloud_provider_id, lb.Traffic_in, lb.Traffic_out)

	w.Flush()
}

func cmdCreate(c *cli.Context) {
	if c.String("protocol") == "HTTPS" {
		utils.FlagsRequired(c, []string{"name", "fqdn", "protocol", "domain_id", "cloud_provider_id", "ssl_certificate", "ssl_certificate_private_key"})
	} else {
		utils.FlagsRequired(c, []string{"name", "fqdn", "protocol", "domain_id", "cloud_provider_id"})
	}
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["name"] = c.String("name")
	v["fqdn"] = c.String("fqdn")
	v["protocol"] = c.String("protocol")
	v["domain_id"] = c.String("domain_id")
	v["cloud_provider_id"] = c.String("cloud_provider_id")
	if c.IsSet("ssl_certificate") {
		v["ssl_certificate"] = c.String("ssl_certificate")
	}
	if c.IsSet("ssl_certificate_private_key") {
		v["ssl_certificate_private_key"] = c.String("ssl_certificate_private_key")
	}
	if c.IsSet("port") {
		v["port"] = c.String("port")
	}
	if c.IsSet("algorithm") {
		v["algorithm"] = c.String("algorithm")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post("/v1/network/load_balancers", jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)

}

func cmdUpdate(c *cli.Context) {
	if c.String("protocol") == "HTTPS" {
		utils.FlagsRequired(c, []string{"id", "name", "fqdn", "protocol", "ssl_certificate", "ssl_certificate_private_key"})
	} else {
		utils.FlagsRequired(c, []string{"id", "name", "fqdn", "protocol"})
	}
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["name"] = c.String("name")
	v["fqdn"] = c.String("fqdn")
	v["protocol"] = c.String("protocol")
	if c.IsSet("ssl_certificate") {
		v["ssl_certificate"] = c.String("ssl_certificate")
	}
	if c.IsSet("ssl_certificate_private_key") {
		v["ssl_certificate_private_key"] = c.String("ssl_certificate_private_key")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Put(fmt.Sprintf("/v1/network/load_balancers/%s", c.String("id")), bytes.NewReader(jsonBytes))

	utils.CheckError(err)
	fmt.Println(res)
}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Delete(fmt.Sprintf("/v1/network/load_balancers/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

	fmt.Println(res)
}

func cmdListNodes(c *cli.Context) {
	var nodes []Node

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/network/load_balancers/%s/nodes", c.String("balancer_id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &nodes)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tPUBLIC IP\tSTATE\tSERVER ID\tPORT\r")

	for _, n := range nodes {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%d\n", n.Id, n.Name, n.PublicIp, n.State, n.ServerId, n.Port)
	}

	w.Flush()
}

func cmdAddNode(c *cli.Context) {

	utils.FlagsRequired(c, []string{"balancer_id", "server_id", "port"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["server_id"] = c.String("server_id")
	v["port"] = c.String("port")

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post(fmt.Sprintf("/v1/network/load_balancers/%s/nodes", c.String("balancer_id")), jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)

}

func cmdDelNode(c *cli.Context) {
	utils.FlagsRequired(c, []string{"balancer_id", "node_id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Delete(fmt.Sprintf("/v1/network/load_balancers/%s/nodes/%s", c.String("balancer_id"), c.String("node_id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

	fmt.Println(res)
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available load balancers",
			Action: cmdList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific load balancer",
			Action: cmdShow,
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
			Action: cmdCreate,
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
			Action: cmdUpdate,
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
			Action: cmdDelete,
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
			Action: cmdListNodes,
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
			Action: cmdAddNode,
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
			Action: cmdDelNode,
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
