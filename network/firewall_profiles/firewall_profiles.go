package firewall_profiles

// import (
// 	"encoding/json"
// 	"fmt"
// 	log "github.com/Sirupsen/logrus"
// 	"github.com/codegangsta/cli"
// 	"github.com/flexiant/concerto/utils"
// 	"github.com/flexiant/concerto/webservice"
// 	"os"
// 	"text/tabwriter"
// )

// func cmdList(c *cli.Context) {
// 	var firewallProfiles []FirewallProfile

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get("/v1/network/firewall_profiles")
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &firewallProfiles)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tDEFAULT\r")

// 	for _, firewallProfile := range firewallProfiles {
// 		fmt.Fprintf(w, "%s\t%s\t%s\t%t\n", firewallProfile.Id, firewallProfile.Name, firewallProfile.Description, firewallProfile.Default)
// 	}

// 	w.Flush()
// }

// func cmdShow(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})
// 	var firewallProfile FirewallProfile

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, data, res := webservice.Get(fmt.Sprintf("/v1/network/firewall_profiles/%s", c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, data)

// 	err = json.Unmarshal(data, &firewallProfile)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tDEFAULT\r")
// 	fmt.Fprintf(w, "%s\t%s\t%s\t%t\n", firewallProfile.Id, firewallProfile.Name, firewallProfile.Description, firewallProfile.Default)
// 	fmt.Fprintln(w, "RULES:\r")
// 	fmt.Fprintln(w, "\tPROTOCOL\tMIN PORT\tMAX PORT\tSOURCE\r")
// 	for _, r := range firewallProfile.Rules {
// 		fmt.Fprintf(w, "\t%s\t%d\t%d\t%s\n", r.Protocol, r.MinPort, r.MaxPort, r.CidrIp)
// 	}
// 	w.Flush()
// }

// func cmdCreate(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"name", "description"})
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	fp := FirewallProfile{
// 		Name:        c.String("name"),
// 		Description: c.String("description"),
// 	}

// 	if c.IsSet("rules") {
// 		var rules []Rule
// 		err = json.Unmarshal([]byte(c.String("rules")), &rules)
// 		utils.CheckError(err)
// 		fp.Rules = rules
// 	}

// 	jsonBytes, err := json.Marshal(fp)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Post("/v1/network/firewall_profiles", jsonBytes)
// 	if res == nil {
// 		log.Fatal(err)
// 	}
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)

// 	var firewallProfile FirewallProfile

// 	err = json.Unmarshal(res, &firewallProfile)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tDEFAULT\r")
// 	fmt.Fprintf(w, "%s\t%s\t%s\t%t\n", firewallProfile.Id, firewallProfile.Name, firewallProfile.Description, firewallProfile.Default)
// 	fmt.Fprintln(w, "RULES:\r")
// 	fmt.Fprintln(w, "\tPROTOCOL\tMIN PORT\tMAX PORT\tSOURCE\r")
// 	for _, r := range firewallProfile.Rules {
// 		fmt.Fprintf(w, "\t%s\t%d\t%d\t%s\n", r.Protocol, r.MinPort, r.MaxPort, r.CidrIp)
// 	}
// 	w.Flush()

// }

// func cmdUpdate(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})
// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	fp := FirewallProfile{
// 		Id: c.String("id"),
// 	}

// 	if c.IsSet("name") {
// 		fp.Name = c.String("name")
// 	}
// 	if c.IsSet("description") {
// 		fp.Description = c.String("description")
// 	}

// 	if c.IsSet("rules") {
// 		var rules []Rule
// 		err = json.Unmarshal([]byte(c.String("rules")), &rules)
// 		utils.CheckError(err)
// 		fp.Rules = rules
// 	}

// 	jsonBytes, err := json.Marshal(fp)
// 	utils.CheckError(err)
// 	err, res, code := webservice.Put(fmt.Sprintf("/v1/network/firewall_profiles/%s", c.String("id")), jsonBytes)
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(code, res)

// 	var firewallProfile FirewallProfile

// 	err = json.Unmarshal(res, &firewallProfile)
// 	utils.CheckError(err)

// 	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
// 	fmt.Fprintln(w, "ID\tNAME\tDESCRIPTION\tDEFAULT\r")
// 	fmt.Fprintf(w, "%s\t%s\t%s\t%t\n", firewallProfile.Id, firewallProfile.Name, firewallProfile.Description, firewallProfile.Default)
// 	fmt.Fprintln(w, "RULES:\r")
// 	fmt.Fprintln(w, "\tPROTOCOL\tMIN PORT\tMAX PORT\tSOURCE\r")
// 	for _, r := range firewallProfile.Rules {
// 		fmt.Fprintf(w, "\t%s\t%d\t%d\t%s\n", r.Protocol, r.MinPort, r.MaxPort, r.CidrIp)
// 	}
// 	w.Flush()

// }

// func cmdDelete(c *cli.Context) {
// 	utils.FlagsRequired(c, []string{"id"})

// 	webservice, err := webservice.NewWebService()
// 	utils.CheckError(err)

// 	err, mesg, res := webservice.Delete(fmt.Sprintf("/v1/network/firewall_profiles/%s", c.String("id")))
// 	utils.CheckError(err)
// 	utils.CheckReturnCode(res, mesg)
// }
