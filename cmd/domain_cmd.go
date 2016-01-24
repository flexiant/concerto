package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/utils/format"
	"os"
)

// WireUpDomain prepares common resources to send request to Concerto API
func WireUpDomain() (*api.DomainService, format.Formatter) {

	f := format.NewTextFormatter(os.Stdout)

	config, err := utils.GetConcertoConfig()
	if err != nil {
		f.PrintFatal("Couldn't wire up config", err)
	}
	hcs, err := utils.NewHTTPConcertoService(config)
	if err != nil {
		f.PrintFatal("Couldn't wire up concerto service", err)
	}
	ds, err := api.NewDomainService(hcs)
	if err != nil {
		f.PrintFatal("Couldn't wire up domain service", err)
	}

	return ds, f
}

// DomainList subcommand function
func DomainList(c *cli.Context) {
	log.Debugf("DomainList: %+v", c.Args().Tail())

	domainSvc, formatter := WireUpDomain()
	domains, headers, err := domainSvc.GetDomainListForPrinting()
	if err != nil {
		formatter.PrintFatal("Couldn't receive domain data", err)
	}
	formatter.PrintList(domains, headers)
}

// DomainShow subcommand function
func DomainShow(c *cli.Context) {
	log.Debugf("DomainShow: %+v", c.Args().Tail())
	utils.CheckRequiredFlags(c, []string{"id"})

	domainSvc, formatter := WireUpDomain()
	domain, headers, err := domainSvc.GetDomainForPrinting(c.String("id"))
	if err != nil {
		formatter.PrintFatal("Couldn't receive domain data", err)
	}
	formatter.PrintItem(domain, headers)

	//
	// w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	// fmt.Fprintln(w, "ID\tNAME\tTTL\tCONTACT\tMINIMUM\tENABLED\r")
	// fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%d\t%t\n", d.Id, d.Name, d.Ttl, d.Contact, d.Minimum, d.Enabled)
	//
	// w.Flush()
}

// func cmdShow(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})
// 	var d Domain
//
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/dns/domains/%s", c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)
//
// 	err = json.Unmarshal(data, &d)
// 	utils.CheckError(err)
//
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tTTL\tCONTACT\tMINIMUM\tENABLED\r")
// 	fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%d\t%t\n", d.ID, d.Name, d.TTL, d.Contact, d.Minimum, d.Enabled)
//
// 	w.Flush()
// }

// func cmdCreate(c *cli.Context) {
//
// 	d := createDomain(c)
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tTTL\tCONTACT\tMINIMUM\tENABLED\r")
// 	fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%d\t%t\n", d.ID, d.Name, d.TTL, d.Contact, d.Minimum, d.Enabled)
//
// 	w.Flush()
//
// }
//
// func cmdUpdate(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	v := make(map[string]string)
//
// 	if c.IsSet("ttl") {
// 		v["ttl"] = c.String("ttl")
// 	}
// 	if c.IsSet("contact") {
// 		v["contact"] = c.String("contact")
// 	}
// 	if c.IsSet("minimum") {
// 		v["minimum"] = c.String("minimum")
// 	}
//
// 	jsonBytes, err := json.Marshal(v)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Put(fmt.Sprintf("/v1/dns/domains/%s", c.String("id")), jsonBytes)
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)
//
// 	var d Domain
// 	err = json.Unmarshal(res, &d)
// 	utils.CheckError(err)
//
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tTTL\tCONTACT\tMINIMUM\tENABLED\r")
// 	fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%d\t%t\n", d.ID, d.Name, d.TTL, d.Contact, d.Minimum, d.Enabled)
//
// 	w.Flush()
// }
//
// func cmdListDomainRecords(c *cli.Context) {
// 	var domainRecords []DomainRecord
// 	utils.FlagsRequired(c, []string{"domain_id"})
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/dns/domains/%s/records", c.String("domain_id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)
//
// 	err = json.Unmarshal(data, &domainRecords)
// 	utils.CheckError(err)
//
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tTYPE\tNAME\tCONTENT\tTTL\tPRIORITY\tSERVER ID\tDOMAIN ID\r")
//
// 	for _, dr := range domainRecords {
// 		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s\n", dr.ID, dr.Type, dr.Name, dr.Content, dr.TTL, dr.Prio, dr.ServerID, dr.DomainID)
// 	}
// 	w.Flush()
// }
//
// func cmdShowDomainRecords(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"domain_id", "id"})
// 	var dr DomainRecord
//
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/dns/domains/%s/records/%s", c.String("domain_id"), c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)
//
// 	err = json.Unmarshal(data, &dr)
// 	utils.CheckError(err)
//
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tTYPE\tNAME\tCONTENT\tTTL\tPRIORITY\tSERVER ID\tDOMAIN ID\r")
// 	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s\n", dr.ID, dr.Type, dr.Name, dr.Content, dr.TTL, dr.Prio, dr.ServerID, dr.DomainID)
//
// 	w.Flush()
// }
//
// func cmdCreateDomainRecords(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"domain_id", "type", "name"})
// 	if c.String("type") == "A" {
// 		if !c.IsSet("content") && !c.IsSet("server_id") {
// 			log.Warn(fmt.Sprintf("Please use either parameter --content or --server_id"))
// 			fmt.Printf("\n")
// 			cli.ShowCommandHelp(c, c.Command.Name)
// 			os.Exit(2)
// 		}
// 	}
// 	if c.String("type") == "AAAA" {
// 		utils.FlagsRequired(c, []string{"content"})
// 	}
//
// 	if c.String("type") == "CNAME" {
// 		utils.FlagsRequired(c, []string{"content"})
// 	}
//
// 	if c.String("type") == "MX" {
// 		utils.FlagsRequired(c, []string{"content", "prio"})
// 	}
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	v := make(map[string]string)
//
// 	v["type"] = c.String("type")
// 	if c.IsSet("name") {
// 		v["name"] = c.String("name")
// 	}
// 	if c.IsSet("content") {
// 		v["content"] = c.String("content")
// 	}
// 	if c.IsSet("ttl") {
// 		v["ttl"] = c.String("ttl")
// 	}
// 	if c.IsSet("prio") {
// 		v["prio"] = c.String("prio")
// 	}
// 	if c.IsSet("server_id") {
// 		v["server_id"] = c.String("server_id")
// 	}
//
// 	jsonBytes, err := json.Marshal(v)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Post(fmt.Sprintf("/v1/dns/domains/%s/records", c.String("domain_id")), jsonBytes)
// 	if res == nil {
// 		log.Fatal(err)
// 	}
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)
//
// 	var dr DomainRecord
// 	err = json.Unmarshal(res, &dr)
// 	utils.CheckError(err)
//
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tTYPE\tNAME\tCONTENT\tTTL\tPRIORITY\tSERVER ID\tDOMAIN ID\r")
// 	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s\n", dr.ID, dr.Type, dr.Name, dr.Content, dr.TTL, dr.Prio, dr.ServerID, dr.DomainID)
//
// 	w.Flush()
//
// }
//
// func cmdUpdateDomainRecords(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"domain_id", "id"})
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	v := make(map[string]string)
//
// 	if c.IsSet("name") {
// 		v["name"] = c.String("name")
// 	}
// 	if c.IsSet("content") {
// 		v["content"] = c.String("content")
// 	}
// 	if c.IsSet("ttl") {
// 		v["ttl"] = c.String("ttl")
// 	}
// 	if c.IsSet("prio") {
// 		v["prio"] = c.String("prio")
// 	}
// 	if c.IsSet("server_id") {
// 		v["server_id"] = c.String("server_id")
// 	}
// 	jsonBytes, err := json.Marshal(v)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Put(fmt.Sprintf("/v1/dns/domains/%s/records/%s", c.String("domain_id"), c.String("id")), jsonBytes)
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)
//
// 	var dr DomainRecord
// 	err = json.Unmarshal(res, &dr)
// 	utils.CheckError(err)
//
// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tTYPE\tNAME\tCONTENT\tTTL\tPRIORITY\tSERVER ID\tDOMAIN ID\r")
// 	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s\n", dr.ID, dr.Type, dr.Name, dr.Content, dr.TTL, dr.Prio, dr.ServerID, dr.DomainID)
//
// 	w.Flush()
//
// }
//
// func cmdDeleteDomainRecords(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"domain_id", "id"})
//
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)
//
// 	err, mesg, res := webservice.Delete(fmt.Sprintf("/v1/dns/domains/%s/records/%s", c.String("domain_id"), c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, mesg)
// }
