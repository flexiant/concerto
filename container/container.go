package container

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/config"
	"github.com/flexiant/concerto/ship"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

func CmbHijack(c *cli.Context) {

	var ships []ship.Ship
	var ship ship.Ship

	discovered := false

	utils.FlagsRequired(c, []string{"ship"})

	shipName := c.String("ship")

	if len(c.Args()) == 0 {
		discovered = true
	} else {
		webservice, err := webservice.NewWebService()
		utils.CheckError(err)

		data, err := webservice.Get("/v1/kaas/ships")
		utils.CheckError(err)

		err = json.Unmarshal(data, &ships)
		utils.CheckError(err)

		// Validating if fleet exist
		for _, element := range ships {
			if element.Name == shipName {
				discovered = true
				ship = element
			}
		}
	}

	if discovered == true {

		var dockerLocation string

		if utils.Exists("/usr/local/bin/docker") {
			dockerLocation = "/usr/local/bin/docker"
		} else {
			//Discover where kubectl is located
			output, err := exec.Command("whereis", "docker").Output()
			utils.CheckError(err)
			dockerLocation = strings.TrimSpace(string(output))
		}

		if len(dockerLocation) > 0 {
			config, err := config.ConcertoServerConfiguration()
			utils.CheckError(err)

			shipParameters := fmt.Sprintf("--host=tcp://%s:2376", ship.Fqdn)
			tls := "--tls=true"
			clientCertificate := fmt.Sprintf("--tlscert=%s", config.Certificate.Cert)
			clientKey := fmt.Sprintf("--tlskey=%s", config.Certificate.Key)
			clientCA := fmt.Sprintf("--tlscacert=%s", config.Certificate.Ca)

			arguments := append([]string{shipParameters, tls, clientCertificate, clientKey, clientCA, c.Args().First()}, c.Args().Tail()...)

			log.Debug(fmt.Sprintf("Going to execute %s %s", dockerLocation, arguments))

			cmd := exec.Command(dockerLocation, arguments...)

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
				fmt.Printf("%s\n", strings.Replace(string(line), "docker", "concerto container", -1))
			}

			go func() {
				time.Sleep(30 * time.Second)
				log.Fatal(fmt.Sprintf("Timeout out. Check conectivity to %s", shipParameters))
			}()

			return
		} else {
			log.Warn(fmt.Sprintf("We could not find docker command line in your enviroment. Please install it. Thank you."))
			os.Exit(1)
		}
	} else {
		log.Warn(fmt.Sprintf("Fleet \"%s\" is not in your account please create it. Thank you.", shipName))
		os.Exit(1)
	}

}
