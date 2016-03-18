package ssh_profiles

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"os"
	"text/tabwriter"
)

func cmdList(c *cli.Context) {
	var sshProfiles []types.SSHProfile

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, data, res := webservice.Get("/v1/cloud/ssh_profiles")
	utils.CheckError(err)
	utils.CheckReturnCode(res, data)

	err = json.Unmarshal(data, &sshProfiles)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\tPUBLIC KEY\tPRIVATE KEY\r")

	for _, sshProfile := range sshProfiles {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", sshProfile.Id, sshProfile.Name, sshProfile.Public_key, sshProfile.Private_key)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var sshProfile types.SSHProfile

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, data, res := webservice.Get(fmt.Sprintf("/v1/cloud/ssh_profiles/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res, data)

	err = json.Unmarshal(data, &sshProfile)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\rPUBLIC KEY\tPRIVATE KEY\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", sshProfile.Id, sshProfile.Name, sshProfile.Public_key, sshProfile.Private_key)
	w.Flush()
}

func cmdCreate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"name", "public_key"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["name"] = c.String("name")
	v["public_key"] = c.String("public_key")

	if c.IsSet("private_key") {
		v["private_key"] = c.String("private_key")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, code := webservice.Post("/v1/cloud/ssh_profiles", jsonBytes)
	if res == nil {
		log.Fatal(err)
	}
	utils.CheckError(err)
	utils.CheckReturnCode(code, res)

	var sshProfile types.SSHProfile
	err = json.Unmarshal(res, &sshProfile)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\rPUBLIC KEY\tPRIVATE KEY\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", sshProfile.Id, sshProfile.Name, sshProfile.Public_key, sshProfile.Private_key)
	w.Flush()
}

func cmdUpdate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	if c.IsSet("name") {
		v["name"] = c.String("name")
	}
	if c.IsSet("public_key") {
		v["public_key"] = c.String("public_key")
	}
	if c.IsSet("private_key") {
		v["private_key"] = c.String("private_key")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, code := webservice.Put(fmt.Sprintf("/v1/cloud/ssh_profiles/%s", c.String("id")), jsonBytes)
	utils.CheckError(err)
	utils.CheckReturnCode(code, res)

	var sshProfile types.SSHProfile
	err = json.Unmarshal(res, &sshProfile)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\rPUBLIC KEY\tPRIVATE KEY\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", sshProfile.Id, sshProfile.Name, sshProfile.Public_key, sshProfile.Private_key)
	w.Flush()
}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, mesg, res := webservice.Delete(fmt.Sprintf("/v1/cloud/ssh_profiles/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res, mesg)
}
