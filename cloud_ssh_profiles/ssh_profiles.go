package cloud_ssh_profiles

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

type SSHProfile struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Public_key  string `json:"public_key"`
	Private_key string `json:"private_key"`
}

func cmdList(c *cli.Context) {
	var sshProfiles []SSHProfile

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/cloud/ssh_profiles")
	utils.CheckError(err)

	err = json.Unmarshal(data, &sshProfiles)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\rPUBLIC KEY\tPRIVATE KEY\r")

	for _, sshProfile := range sshProfiles {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", sshProfile.Id, sshProfile.Name, sshProfile.Public_key, sshProfile.Private_key)
	}

	w.Flush()
}

func cmdShow(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})
	var sshProfile SSHProfile

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get(fmt.Sprintf("/v1/cloud/ssh_profiles/%s", c.String("id")))
	utils.CheckError(err)

	err = json.Unmarshal(data, &sshProfile)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tNAME\rPUBLIC KEY\tPRIVATE KEY\r")
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", sshProfile.Id, sshProfile.Name, sshProfile.Public_key, sshProfile.Private_key)
	w.Flush()
}

func cmdCreate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"name"})
	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["name"] = c.String("name")
	if c.IsSet("public_key") {
		v["public_key"] = c.String("public_key")
	}
	if c.IsSet("private_key") {
		v["private_key"] = c.String("private_key")
	}

	jsonBytes, err := json.Marshal(v)
	utils.CheckError(err)
	err, res, _ := webservice.Post("/v1/cloud/ssh_profiles", jsonBytes)
	if res == "" {
		log.Fatal(err)
	}
	utils.CheckError(err)

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
	err, res := webservice.Put(fmt.Sprintf("/v1/cloud/ssh_profiles/%s", c.String("id")), bytes.NewReader(jsonBytes))

	utils.CheckError(err)
	fmt.Println(res)
}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, res := webservice.Delete(fmt.Sprintf("/v1/cloud/ssh_profiles/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

	fmt.Println(res)
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available SSH profiles.",
			Action: cmdList,
		},
		{
			Name:   "show",
			Usage:  "Shows information about the SSH profile identified by the given id.",
			Action: cmdShow,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "SSH profile id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a new SSH profile.",
			Action: cmdCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the SSH profile",
				},
				cli.StringFlag{
					Name:  "public_key",
					Usage: "Public key of the SSH profile",
				},
				cli.StringFlag{
					Name:  "private_key",
					Usage: "Private key of the SSH profile",
				},
			},
		},
		{
			Name:   "update",
			Usage:  "Updates an existing SSH profile",
			Action: cmdUpdate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of the SSH profile",
				},
				cli.StringFlag{
					Name:  "public_key",
					Usage: "Public key of the SSH profile",
				},
				cli.StringFlag{
					Name:  "private_key",
					Usage: "Private key of the SSH profile",
				},
			},
		},
		{
			Name:   "destroy",
			Usage:  "Destroys an SSH profile",
			Action: cmdDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "SSH profile id",
				},
			},
		},
	}
}
