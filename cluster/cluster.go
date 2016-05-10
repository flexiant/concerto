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
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/api/types"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
)

func cmdKubectlHijack(c *cli.Context) error {
	var clusters []types.Cluster
	var cluster types.Cluster

	discovered := false
	operational := false

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

	err, data, res := webservice.Get("/v1/kaas/fleets")
	utils.CheckError(err)
	utils.CheckReturnCode(res, data)

	err = json.Unmarshal(data, &clusters)
	utils.CheckError(err)

	// Validating if cluster exist
	for _, element := range clusters {
		if (element.Name == clusterName) || (element.Id == clusterName) {
			discovered = true
			cluster = element
			if cluster.State == "operational" || cluster.State == "partially_operational" {
				operational = true
			}
		}
	}

	if (discovered && operational) || firstArgument == "help" {

		kubeLocation, err := exec.LookPath("kubectl")
		if err != nil {
			log.Warn(fmt.Sprintf("We could not find kubectl in your enviroment. Please install it."))
			os.Exit(1)
		}

		if discovered && !operational {
			log.Warn(fmt.Sprintf("Cluster \"%s\" is not operational. Wait till it gets operational.", clusterName))
		}

		log.Debug(fmt.Sprintf("Found kubectl at %s", kubeLocation))
		config, err := utils.GetConcertoConfig()
		utils.CheckError(err)

		var clusterParameters, clientCertificate, clientKey, clientCA string

		if len(cluster.Masters) > 0 {
			clusterParameters = fmt.Sprintf("--server=https://%s:6443", cluster.Masters[0])
		}

		clientCertificate = fmt.Sprintf("--client-certificate=%s", config.Certificate.Cert)
		clientKey = fmt.Sprintf("--client-key=%s", config.Certificate.Key)
		clientCA = fmt.Sprintf("--certificate-authority=%s", config.Certificate.Ca)

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

		return nil
	} else {
		log.Warn(fmt.Sprintf("Cluster \"%s\" is not in your account please create it. Thank you.", clusterName))
		os.Exit(1)
	}
	return nil
}
