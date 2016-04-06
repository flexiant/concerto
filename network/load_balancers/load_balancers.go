package load_balancers

// import (
// 	"encoding/json"
// 	"fmt"
// 	log "github.com/Sirupsen/logrus"
// 	"github.com/codegangsta/cli"
// 	"github.com/flexiant/concerto/api/types"
// 	"github.com/flexiant/concerto/utils"
// 	"github.com/flexiant/concerto/webservice"
// 	"os"
// 	"strings"
// 	"text/tabwriter"
// )

// func cmdList(c *cli.Context) {
// 	var loadBalancers []LoadBalancer

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get("/v1/network/load_balancers")
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &loadBalancers)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tFQDN\tPROTOCOL\tPORT\tALGORITHM\tSSL CERTIFICATE\tSSL CERTIFICATE PRIVATE KEY\tDOMAIN ID\tCLOUD PROVIDER ID\tTRAFFIC IN\tTRAFFIC OUT\r")

// 	for _, lb := range loadBalancers {
// 		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%s\t%s\t%s\t%s\t%s\t%d\t%d\n", lb.Id, lb.Name, lb.Fqdn, lb.Protocol, lb.Port, lb.Algorithm, lb.SslCertificate, lb.Ssl_certificate_private_key, lb.Domain_id, lb.Cloud_provider_id, lb.Traffic_in, lb.Traffic_out)
// 	}

// 	w.Flush()
// }

// func cmdShow(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})
// 	var lb LoadBalancer

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/network/load_balancers/%s", c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &lb)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tFQDN\tPROTOCOL\tPORT\tALGORITHM\tSSL CERTIFICATE\tSSL CERTIFICATE PRIVATE KEY\tDOMAIN ID\tCLOUD PROVIDER ID\tTRAFFIC IN\tTRAFFIC OUT\r")
// 	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%s\t%s\t%s\t%s\t%s\t%d\t%d\n", lb.Id, lb.Name, lb.Fqdn, lb.Protocol, lb.Port, lb.Algorithm, lb.SslCertificate, lb.Ssl_certificate_private_key, lb.Domain_id, lb.Cloud_provider_id, lb.Traffic_in, lb.Traffic_out)

// 	w.Flush()
// }

// func cmdCreate(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"protocol"})
// 	if c.String("protocol") == "https" {
// 		utils.FlagsRequired(c, []string{"name", "fqdn", "protocol", "domain_id", "cloud_provider_id", "ssl_certificate", "ssl_certificate_private_key"})
// 	} else {
// 		utils.FlagsRequired(c, []string{"name", "fqdn", "protocol", "domain_id", "cloud_provider_id"})
// 	}
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	v := make(map[string]string)

// 	v["name"] = c.String("name")
// 	v["fqdn"] = c.String("fqdn")
// 	v["protocol"] = strings.ToLower(c.String("protocol"))
// 	v["domain_id"] = c.String("domain_id")
// 	v["cloud_provider_id"] = c.String("cloud_provider_id")
// 	if c.IsSet("ssl_certificate") {
// 		v["ssl_certificate"] = c.String("ssl_certificate")
// 	}
// 	if c.IsSet("ssl_certificate_private_key") {
// 		v["ssl_certificate_private_key"] = c.String("ssl_certificate_private_key")
// 	}
// 	if c.IsSet("port") {
// 		v["port"] = c.String("port")
// 	}
// 	if c.IsSet("algorithm") {
// 		v["algorithm"] = c.String("algorithm")
// 	}

// 	jsonBytes, err := json.Marshal(v)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Post("/v1/network/load_balancers", jsonBytes)
// 	if res == nil {
// 		log.Fatal(err)
// 	}
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)

// 	var lb types.LoadBalancer

// 	err = json.Unmarshal(res, &lb)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tFQDN\tPROTOCOL\tPORT\tALGORITHM\tSSL CERTIFICATE\tSSL CERTIFICATE PRIVATE KEY\tDOMAIN ID\tCLOUD PROVIDER ID\tTRAFFIC IN\tTRAFFIC OUT\r")
// 	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%s\t%s\t%s\t%s\t%s\t%d\t%d\n", lb.Id, lb.Name, lb.Fqdn, lb.Protocol, lb.Port, lb.Algorithm, lb.SslCertificate, lb.Ssl_certificate_private_key, lb.Domain_id, lb.Cloud_provider_id, lb.Traffic_in, lb.Traffic_out)

// 	w.Flush()

// }

// func cmdUpdate(c *cli.Context) {
// 	if c.String("protocol") == "HTTPS" {
// 		utils.FlagsRequired(c, []string{"id", "name", "fqdn", "protocol", "ssl_certificate", "ssl_certificate_private_key"})
// 	} else {
// 		utils.FlagsRequired(c, []string{"id", "name", "fqdn", "protocol"})
// 	}
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	v := make(map[string]string)

// 	v["name"] = c.String("name")
// 	v["fqdn"] = c.String("fqdn")
// 	v["protocol"] = strings.ToLower(c.String("protocol"))
// 	if c.IsSet("ssl_certificate") {
// 		v["ssl_certificate"] = c.String("ssl_certificate")
// 	}
// 	if c.IsSet("ssl_certificate_private_key") {
// 		v["ssl_certificate_private_key"] = c.String("ssl_certificate_private_key")
// 	}

// 	jsonBytes, err := json.Marshal(v)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Put(fmt.Sprintf("/v1/network/load_balancers/%s", c.String("id")), jsonBytes)
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)

// 	var lb LoadBalancer

// 	err = json.Unmarshal(res, &lb)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tFQDN\tPROTOCOL\tPORT\tALGORITHM\tSSL CERTIFICATE\tSSL CERTIFICATE PRIVATE KEY\tDOMAIN ID\tCLOUD PROVIDER ID\tTRAFFIC IN\tTRAFFIC OUT\r")
// 	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%s\t%s\t%s\t%s\t%s\t%d\t%d\n", lb.Id, lb.Name, lb.Fqdn, lb.Protocol, lb.Port, lb.Algorithm, lb.SslCertificate, lb.Ssl_certificate_private_key, lb.Domain_id, lb.Cloud_provider_id, lb.Traffic_in, lb.Traffic_out)

// 	w.Flush()

// }

// func cmdDelete(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, mesg, res := webservice.Delete(fmt.Sprintf("/v1/network/load_balancers/%s", c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, mesg)
// }

// func cmdListNodes(c *cli.Context) {
// 	var nodes []Node

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/network/load_balancers/%s/nodes", c.String("balancer_id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &nodes)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tPUBLIC IP\tSTATE\tSERVER ID\tPORT\r")

// 	for _, n := range nodes {
// 		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%d\n", n.Id, n.Name, n.PublicIp, n.State, n.ServerId, n.Port)
// 	}

// 	w.Flush()
// }

// func cmdAddNode(c *cli.Context) {

// 	utils.FlagsRequired(c, []string{"balancer_id", "server_id", "port"})

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	v := make(map[string]string)

// 	v["server_id"] = c.String("server_id")
// 	v["port"] = c.String("port")

// 	jsonBytes, err := json.Marshal(v)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Post(fmt.Sprintf("/v1/network/load_balancers/%s/nodes", c.String("balancer_id")), jsonBytes)
// 	if res == nil {
// 		log.Fatal(err)
// 	}
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)

// 	var n Node

// 	err = json.Unmarshal(res, &n)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tPUBLIC IP\tSTATE\tSERVER ID\tPORT\r")
// 	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%d\n", n.Id, n.Name, n.PublicIp, n.State, n.ServerId, n.Port)

// 	w.Flush()

// }

// func cmdDelNode(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"balancer_id", "node_id"})

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, mesg, res := webservice.Delete(fmt.Sprintf("/v1/network/load_balancers/%s/nodes/%s", c.String("balancer_id"), c.String("node_id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, mesg)
// }
