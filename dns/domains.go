/*
  	Dns Domains

	The domains of the domain-name system (DNS) and their records are used by our name servers for name lookup.
  	Each workspace belongs to a domain, and each server has an A or an AAAA record in the domain of its workspace.
  	The user can add additional records - currently supported are: A, AAAA, CNAME, MX, and TXT records.

	The available commands are:
		list
		show
		create
		update
		delete
		list_domain_records
		show_domain_record
		create_domain_record
		update_domain_record
		delete_domain_record

	Use "dns_domains --help" or "dns --help" on the commandline interface for more information about the available subcommands.

	Domains list

	Lists the domains of the account group.

	Usage:

		dns_domains list

	Domains show

	Shows information about a specific domain.

	Usage:

		dns_domains show (options)

	Options:
		--id <domain_id> 		Identifier of the domain

	Domain create

	Creates a new domain

	Usage:

		dns_domains create (options)

	Options:
		--name <name> 			Fully-qualified domain name (FQDN)
		--ttl <value> 	Time to live (TTL) of the Start of Authority (SOA) record
		--contact <email> 			Contact e-mail
		--minimum <value>	The minimum TTL of the SOA record

	Domain update

	Updates an existing domain.

	Usage:

		dns_domains update (options)

	Options:
		--id <domain_id> 		Identifier of the domain
		--ttl <value> 	Time to live (TTL) of the Start of Authority (SOA) record
		--contact <email> 			Contact e-mail
		--minimum <value>	The minimum TTL of the SOA record

	Domain delete

	This action deletes a domain.

	Usage:

		dns_domains delete (options)

	Options:
		--id <domain_id> 		Identifier of the domain

	List domain records

	Lists the DNS records of a domain.

	Usage:

		dns_domains list_domain_records

	Get domain records

	Shows information about a specific DNS record.

	Usage:

		dns_domains get_domain_record (options)

	Options:
		--domain_id <domain_id> 		Identifier of the domain
		--record_id <record_id> 		Identifier of the DNS record

	Domain record create

	Creates a new DNS record

	Usage:

		dns_domains create_domain_record (options)

	Options:
		--type <type> 			Type of record (A, AAAA, CNAME, MX, TXT)
		--name <name> 	Record name
		--content <content> 			Record content
		--ttl <value>	Time to live (TTL)
		--prio <value>	Priority (only MX records)
		--server_id	<server_id>	Identifier of the associated server (only A and AAAA records)

	Domain record update

	Updates an existing DNS record.

	Usage:

		dns_domains update_domain_record (options)

	Options:
		--domain_id <domain_id> 		Identifier of the domain
		--record_id <record_id> 		Identifier of the DNS record
		--name <name> 	Record name
		--content <content> 			Record content
		--ttl <value>	Time to live (TTL)
		--prio <value>	Priority (only MX records)
		--server_id	<server_id>	Identifier of the associated server (only A and AAAA records)

	Domain record delete

	This action deletes a DNS record.

	Usage:

		dns_domains delete_domain_record (options)

	Options:
		--domain_id <domain_id> 		Identifier of the domain
		--record_id <record_id> 		Identifier of the DNS record



*/
package dns

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

type Domain struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Ttl     int    `json:"ttl"`
	Contact string `json:"contact"`
	Minimum int    `json:"minimum"`
	Enabled bool   `json:"enabled"`
}

type DomainRecord struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	Ttl       int    `json:"ttl"`
	Prio      int    `json:"prio"`
	Server_id string `json:"server_id"`
	Domain_id string `json:"domain_id"`
}

func cmdList(c *cli.Context) {
	var domains []Domain

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/dns/domains")
	utils.CheckError(err)

	err = json.Unmarshal(data, &domains)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tTTL\tCONTACT\tMINIMUM\tENABLED\r")

	for _, d := range domains {
		fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%d\t%t\n", d.Id, d.Name, d.Ttl, d.Contact, d.Minimum, d.Enabled)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var d Domain

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/dns/domains/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &d)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tTTL\tCONTACT\tMINIMUM\tENABLED\r")
	fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%d\t%t\n", d.Id, d.Name, d.Ttl, d.Contact, d.Minimum, d.Enabled)

	w.Flush()
}

func cmdCreate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"name", "contact"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["name"] = c.String("name")
	if c.IsSet("ttl") {
		v["ttl"] = c.String("ttl")
	}
	v["contact"] = c.String("contact")

	if c.IsSet("minimum") {
		v["minimum"] = c.String("minimum")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post("/v1/dns/domains", jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)

	var d Domain
	err = json.Unmarshal(res, &d)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tTTL\tCONTACT\tMINIMUM\tENABLED\r")
	fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%d\t%t\n", d.Id, d.Name, d.Ttl, d.Contact, d.Minimum, d.Enabled)

	w.Flush()

}

func cmdUpdate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	if c.IsSet("ttl") {
		v["ttl"] = c.String("ttl")
	}
	if c.IsSet("contact") {
		v["contact"] = c.String("contact")
	}
	if c.IsSet("minimum") {
		v["minimum"] = c.String("minimum")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Put(fmt.Sprintf("/v1/dns/domains/%s", c.String("id")), jsonBytes)

	utils.CheckError(err)

	var d Domain
	err = json.Unmarshal(res, &d)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tTTL\tCONTACT\tMINIMUM\tENABLED\r")
	fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%d\t%t\n", d.Id, d.Name, d.Ttl, d.Contact, d.Minimum, d.Enabled)

	w.Flush()
}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Delete(fmt.Sprintf("/v1/dns/domains/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

	fmt.Println(res)
}

func cmdListDomainRecords(c *cli.Context) {
	var domainRecords []DomainRecord
	utils.FlagsRequired(c, []string{"domain_id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/dns/domains/%s/records", c.String("domain_id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &domainRecords)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTYPE\tNAME\tCONTENT\tTTL\tPRIORITY\tSERVER ID\tDOMAIN ID\r")

	for _, dr := range domainRecords {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s\n", dr.Id, dr.Type, dr.Name, dr.Content, dr.Ttl, dr.Prio, dr.Server_id, dr.Domain_id)
	}
	w.Flush()
}

func cmdShowDomainRecords(c *cli.Context) {
	utils.FlagsRequired(c, []string{"domain_id", "id"})
	var dr DomainRecord

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/dns/domains/%s/records/%s", c.String("domain_id"), c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &dr)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTYPE\tNAME\tCONTENT\tTTL\tPRIORITY\tSERVER ID\tDOMAIN ID\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s\n", dr.Id, dr.Type, dr.Name, dr.Content, dr.Ttl, dr.Prio, dr.Server_id, dr.Domain_id)

	w.Flush()
}

func cmdCreateDomainRecords(c *cli.Context) {
	utils.FlagsRequired(c, []string{"domain_id", "type", "name"})
	if c.String("type") == "A" {
		if !c.IsSet("content") && !c.IsSet("server_id") {
			log.Warn(fmt.Sprintf("Please use either parameter --content or --server_id"))
			fmt.Printf("\n")
			cli.ShowCommandHelp(c, c.Command.Name)
			os.Exit(2)
		}
	}
	if c.String("type") == "AAAA" {
		utils.FlagsRequired(c, []string{"content"})
	}

	if c.String("type") == "CNAME" {
		utils.FlagsRequired(c, []string{"content"})
	}

	if c.String("type") == "MX" {
		utils.FlagsRequired(c, []string{"content", "prio"})
	}
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["type"] = c.String("type")
	if c.IsSet("name") {
		v["name"] = c.String("name")
	}
	if c.IsSet("content") {
		v["content"] = c.String("content")
	}
	if c.IsSet("ttl") {
		v["ttl"] = c.String("ttl")
	}
	if c.IsSet("prio") {
		v["prio"] = c.String("prio")
	}
	if c.IsSet("server_id") {
		v["server_id"] = c.String("server_id")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post(fmt.Sprintf("/v1/dns/domains/%s/records", c.String("domain_id")), jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)

	var dr DomainRecord
	err = json.Unmarshal(res, &dr)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTYPE\tNAME\tCONTENT\tTTL\tPRIORITY\tSERVER ID\tDOMAIN ID\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s\n", dr.Id, dr.Type, dr.Name, dr.Content, dr.Ttl, dr.Prio, dr.Server_id, dr.Domain_id)

	w.Flush()

}

func cmdUpdateDomainRecords(c *cli.Context) {
	utils.FlagsRequired(c, []string{"domain_id", "id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	if c.IsSet("name") {
		v["name"] = c.String("name")
	}
	if c.IsSet("content") {
		v["content"] = c.String("content")
	}
	if c.IsSet("ttl") {
		v["ttl"] = c.String("ttl")
	}
	if c.IsSet("prio") {
		v["prio"] = c.String("prio")
	}
	if c.IsSet("server_id") {
		v["server_id"] = c.String("server_id")
	}
	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Put(fmt.Sprintf("/v1/dns/domains/%s/records/%s", c.String("domain_id"), c.String("id")), jsonBytes)

	utils.CheckError(err)

	var dr DomainRecord
	err = json.Unmarshal(res, &dr)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tTYPE\tNAME\tCONTENT\tTTL\tPRIORITY\tSERVER ID\tDOMAIN ID\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s\n", dr.Id, dr.Type, dr.Name, dr.Content, dr.Ttl, dr.Prio, dr.Server_id, dr.Domain_id)

	w.Flush()

}

func cmdDeleteDomainRecords(c *cli.Context) {
	utils.FlagsRequired(c, []string{"domain_id", "id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Delete(fmt.Sprintf("/v1/dns/domains/%s/records/%s", c.String("domain_id"), c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

	fmt.Println(res)
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the domains of the account group.",
			Action: cmdList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific domain.",
			Action: cmdShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Domain Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new domain.",
			Action: cmdCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Fully-qualified domain name (FQDN)",
				},
				cli.StringFlag{
					Name:  "ttl",
					Usage: "Time to live (TTL) of the Start of Authority (SOA) record",
				},
				cli.StringFlag{
					Name:  "contact",
					Usage: "Contact e-mail",
				},
				cli.StringFlag{
					Name:  "minimum",
					Usage: "The minimum TTL of the SOA record",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing domain",
			Action: cmdUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Domain Id",
				},
				cli.StringFlag{
					Name:  "ttl",
					Usage: "Time to live (TTL) of the Start of Authority (SOA) record",
				},
				cli.StringFlag{
					Name:  "contact",
					Usage: "Contact e-mail",
				},
				cli.StringFlag{
					Name:  "minimum",
					Usage: "The minimum TTL of the SOA record",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a domain",
			Action: cmdDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Domain Id",
				},
			},
		},
		{
			Name:   "list_domain_records",
			Usage:  "Lists the DNS records of a domain.",
			Action: cmdListDomainRecords,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Domain Id",
				},
			},
		},
		{
			Name:   "get_domain_record",
			Usage:  "Shows information about a specific DNS record.",
			Action: cmdShowDomainRecords,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Domain Id",
				},
				cli.StringFlag{
					Name:  "id",
					Usage: "Record Id",
				},
			},
		},
		{
			Name:   "create_domain_record",
			Usage:  "Creates a new DNS record.",
			Action: cmdCreateDomainRecords,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Domain Id",
				},
				cli.StringFlag{
					Name:  "type",
					Usage: "Type of record (A, AAAA, CNAME, MX, TXT)",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Record name",
				},
				cli.StringFlag{
					Name:  "content",
					Usage: "Record content",
				},
				cli.StringFlag{
					Name:  "ttl",
					Usage: "Time to live (TTL)",
				},
				cli.StringFlag{
					Name:  "prio",
					Usage: "Priority (only MX records)",
				},
				cli.StringFlag{
					Name:  "server_id",
					Usage: "Identifier of the associated server (only A and AAAA records)",
				},
			},
		},
		{
			Name:   "update_domain_record",
			Usage:  "Updates an existing DNS record.",
			Action: cmdUpdateDomainRecords,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Domain Id",
				},
				cli.StringFlag{
					Name:  "id",
					Usage: "Record Id",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Record name",
				},
				cli.StringFlag{
					Name:  "content",
					Usage: "Record content",
				},
				cli.StringFlag{
					Name:  "ttl",
					Usage: "Time to live (TTL)",
				},
				cli.StringFlag{
					Name:  "prio",
					Usage: "Priority (only MX records)",
				},
				cli.StringFlag{
					Name:  "server_id",
					Usage: "Identifier of the associated server (only A and AAAA records)",
				},
			},
		},
		{
			Name:   "delete_domain_record",
			Usage:  "Deletes a DNS record",
			Action: cmdDeleteDomainRecords,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Record Id",
				},
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Record Id",
				},
			},
		},
	}
}
