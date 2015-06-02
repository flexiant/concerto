package dns

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
	err, res, _ := webservice.Put(fmt.Sprintf("/v1/dns/domains/%s", c.String("id")), bytes.NewReader(jsonBytes))

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
	utils.FlagsRequired(c, []string{"domain_id", "type"})
	if c.String("type") == "A" {
		//FIXME add content of server_id
		utils.FlagsRequired(c, []string{"name"})
	}
	if c.String("type") == "AAAA" {
		utils.FlagsRequired(c, []string{"name", "content"})
	}

	if c.String("type") == "CNAME" {
		utils.FlagsRequired(c, []string{"name", "content"})
	}

	if c.String("type") == "MX" {
		utils.FlagsRequired(c, []string{"name", "content", "prio"})
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
	err, res, _ := webservice.Put(fmt.Sprintf("/v1/dns/domains/%s/records/%s", c.String("domain_id"), c.String("id")), bytes.NewReader(jsonBytes))

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
			},
		},
	}
}
