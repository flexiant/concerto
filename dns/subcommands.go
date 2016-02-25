package dns

import (
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/cmd"
)

// SubCommands return CLI subcommands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists the domains of the account group.",
			Action: cmd.DomainList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about a specific domain.",
			Action: cmd.DomainShow,
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
			Action: cmd.DomainCreate,
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
			Action: cmd.DomainUpdate,
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
			Action: cmd.DomainDelete,
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
			Action: cmd.DomainRecordList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Domain Id",
				},
			},
		},
		{
			Name:   "show_domain_record",
			Usage:  "Shows information about a specific DNS record.",
			Action: cmd.DomainRecordShow,
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
			Action: cmd.DomainRecordCreate,
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
			Action: cmd.DomainRecordUpdate,
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
			Action: cmd.DomainRecordDelete,
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
