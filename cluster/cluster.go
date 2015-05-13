package cluster

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/config"
	"github.com/flexiant/concerto/fleet"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

func CmbHijack(c *cli.Context) {
	var fleets []fleet.Fleet
	var fleet fleet.Fleet

	discovered := false

	utils.FlagsRequired(c, []string{"fleet"})

	fleetName := c.String("fleet")

	if len(c.Args()) == 0 {
		discovered = true
	} else {
		webservice, err := webservice.NewWebService()
		utils.CheckError(err)

		data, err := webservice.Get("/v1/kaas/fleets")
		utils.CheckError(err)

		err = json.Unmarshal(data, &fleets)
		utils.CheckError(err)

		// Validating if fleet exist
		for _, element := range fleets {
			if element.Name == fleetName {
				discovered = true
				fleet = element
			}
		}
	}

	if discovered == true {
		//Discover where kubectl is located
		output, err := exec.Command("whereis", "kubectl").Output()
		utils.CheckError(err)
		kubeLocation := strings.TrimSpace(string(output))

		if len(kubeLocation) > 0 {
			config, err := config.ConcertoServerConfiguration()
			utils.CheckError(err)

			fleetParameters := fmt.Sprintf("--server=https://%s:6443", fleet.Masters[0])
			clientCertificate := fmt.Sprintf("--client-certificate=%s", config.Certificate.Cert)
			clientKey := fmt.Sprintf("--client-key=%s", config.Certificate.Key)
			clientCA := fmt.Sprintf("--certificate-authority=%s", config.Certificate.Ca)

			arguments := append([]string{fleetParameters, clientCertificate, clientKey, clientCA, c.Args().First()}, c.Args().Tail()...)

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
				fmt.Printf("%s\n", strings.Replace(string(line), "kubectl", "concerto cluster", -1))
			}

			go func() {
				time.Sleep(30 * time.Second)
				log.Fatal(fmt.Sprintf("Timeout out. Check conectivity to %s", fleetParameters))
			}()

			return
		} else {
			log.Warn(fmt.Sprintf("We could not find kubectl in your enviroment. Please install it. Thank you."))
			os.Exit(1)
		}
	} else {
		log.Warn(fmt.Sprintf("Fleet \"%s\" is not in your account please create it. Thank you.", fleetName))
		os.Exit(1)
	}

}
