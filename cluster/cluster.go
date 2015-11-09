package cluster

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/config"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
)

type Cluster struct {
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	State             string   `json:"state"`
	MasterCount       int      `json:"master_count"`
	SlaveCount        int      `json:"slave_count"`
	WorkspaceId       string   `json:"workspace_id"`
	FirewallProfileId string   `json:"firewall_profile_id"`
	MasterTemplateId  string   `json:"master_template_id"`
	SlaveTemplateId   string   `json:"slave_template_id"`
	Masters           []string `json:"masters"`
}

func cmdCreate(c *cli.Context) {
	utils.FlagsRequired(c, []string{"cluster"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	v := make(map[string]string)

	v["name"] = c.String("cluster")
	if c.IsSet("domain_id") {
		v["domain_id"] = c.String("domain_id")
	}

	json, err := json.Marshal(v)
	utils.CheckError(err)

	err, _, code := webservice.Post("/v1/kaas/fleets", json)
	utils.CheckError(err)
	utils.CheckReturnCode(code)

}

func cmdDelete(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Delete(fmt.Sprintf("/v1/kaas/fleets/%s", c.String("id")))
	utils.CheckError(err)
	utils.CheckReturnCode(res)

}

func cmdStart(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Put(fmt.Sprintf("/v1/kaas/fleets/%s/start", c.String("id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(res)

}

func cmdStop(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Put(fmt.Sprintf("/v1/kaas/fleets/%s/stop", c.String("id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(res)

}

func cmdEmpty(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Put(fmt.Sprintf("/v1/kaas/fleets/%s/empty", c.String("id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(res)

}

func cmdAttachNet(c *cli.Context) {
	utils.FlagsRequired(c, []string{"id"})

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	err, _, res := webservice.Put(fmt.Sprintf("/v1/kaas/fleets/%s/attach_network", c.String("id")), nil)
	utils.CheckError(err)
	utils.CheckReturnCode(res)

}

func cmdList(c *cli.Context) {
	var clusters []Cluster

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/kaas/fleets")
	utils.CheckError(err)

	err = json.Unmarshal(data, &clusters)
	utils.CheckError(err)

	w := tabwriter.NewWriter(os.Stdout, 15, 1, 3, ' ', 0)
	fmt.Fprintln(w, "CLUSTER\tID\tSTATE\tMASTER COUNT\tSLAVE COUNT")

	for _, cluster := range clusters {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%d\n", cluster.Name, cluster.Id, cluster.State, cluster.MasterCount, cluster.SlaveCount)
	}

	w.Flush()
}

func cmdKubectlHijack(c *cli.Context) {
	var clusters []Cluster
	var cluster Cluster

	discovered := false

	utils.FlagsRequired(c, []string{"cluster"})

	clusterName := c.String("cluster")

	var firstArgument string
	if c.Args().Present() {
		firstArgument = c.Args().First()
	} else {
		firstArgument = "help"
	}

	webservice, err := webservice.NewWebService()
	utils.CheckError(err)

	data, err := webservice.Get("/v1/kaas/fleets")
	utils.CheckError(err)

	err = json.Unmarshal(data, &clusters)
	utils.CheckError(err)

	// Validating if cluster exist
	for _, element := range clusters {
		if (element.Name == clusterName) || (element.Id == clusterName) {
			discovered = true
			cluster = element
		}
	}

	if discovered == true {
		//Discover where kubectl is located
		output, err := exec.Command("whereis", "kubectl").Output()
		utils.CheckError(err)

		kubeLocation := strings.TrimSpace(string(output))

		if !(len(kubeLocation) > 0) {
			log.Info("Not found kubectl with whereis going to try which")
			//Discover where kubectl is located
			output, err = exec.Command("which", "kubectl").Output()
			utils.CheckError(err)

			kubeLocation = strings.TrimSpace(string(output))
		}

		if len(kubeLocation) > 0 {
			log.Debug(fmt.Sprintf("Found kubectl at %s", kubeLocation))
			config, err := config.ConcertoServerConfiguration()
			utils.CheckError(err)

			clusterParameters := fmt.Sprintf("--server=https://%s:6443", cluster.Masters[0])
			clientCertificate := fmt.Sprintf("--client-certificate=%s", config.Certificate.Cert)
			clientKey := fmt.Sprintf("--client-key=%s", config.Certificate.Key)
			clientCA := fmt.Sprintf("--certificate-authority=%s", config.Certificate.Ca)

			arguments := append([]string{clusterParameters, "--api-version=v1", clientCertificate, clientKey, clientCA, firstArgument}, c.Args().Tail()...)

			log.Debug(fmt.Sprintf("Going to execute %s %s", kubeLocation, arguments))

			cmd := exec.Command(kubeLocation, arguments...)

			stdout, err := cmd.StdoutPipe()
			utils.CheckError(err)

			stderr, err := cmd.StderrPipe()
			utils.CheckError(err)

			// Start command
			err = cmd.Start()
			utils.CheckError(err)
			defer cmd.Wait()

			go io.Copy(os.Stderr, stderr)

			ls := bufio.NewReader(stdout)

			for {
				line, isPrefix, err := ls.ReadLine()
				if isPrefix {
					log.Errorf("%s", errors.New("isPrefix: true"))
				}
				if err != nil {
					if err != io.EOF {
						log.Errorf("%s", err.Error())
					}
					break
				}
				fmt.Printf("%s\n", strings.Replace(string(line), "kubectl", fmt.Sprintf("concerto cluster kubectl --cluster %s", clusterName), -1))
			}

			go func() {
				time.Sleep(30 * time.Second)
				log.Fatal(fmt.Sprintf("Timeout out. Check conectivity to %s", clusterParameters))
			}()

			return
		} else {
			log.Warn(fmt.Sprintf("We could not find kubectl in your enviroment. Please install it. Thank you."))
			os.Exit(1)
		}
	} else {
		log.Warn(fmt.Sprintf("Cluster \"%s\" is not in your account please create it. Thank you.", clusterName))
		os.Exit(1)
	}

}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Usage:  "Lists all available Clusters",
			Action: cmdList,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cluster",
					Usage: "Cluster Name to Attach Ship",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Name of Host",
				},
				cli.StringFlag{
					Name:  "fqdn",
					Usage: "Full Qualify Domain Name of Host",
				},
				cli.StringFlag{
					Name:  "plan",
					Usage: "Server Plan to Use to Create Host",
				},
			},
		},
		{
			Name:   "start",
			Usage:  "Starts a given Cluster",
			Action: cmdStart,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "stop",
			Usage:  "Stops a given Cluster",
			Action: cmdStop,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "empty",
			Usage:  "Empties a given Cluster",
			Action: cmdEmpty,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "attach_net",
			Usage:  "Attaches network to a given Cluster",
			Action: cmdAttachNet,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "create",
			Usage:  "Creates a Cluster",
			Action: cmdCreate,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cluster",
					Usage: "Cluster Name",
				},
				cli.StringFlag{
					Name:  "domain_id",
					Usage: "Domain Id",
				},
			},
		},
		{
			Name:   "delete",
			Usage:  "Deletes a given Cluster",
			Action: cmdDelete,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "id",
					Usage: "Cluster Id",
				},
			},
		},
		{
			Name:   "kubectl",
			Usage:  "Kubectl command line wrapper",
			Action: cmdKubectlHijack,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "cluster",
					Usage: "Cluster Name",
				},
			},
		},
	}
}
