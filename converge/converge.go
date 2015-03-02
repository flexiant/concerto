package converge

import (
	"bufio"
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"io"
	"os/exec"
	"path"
	"regexp"
	"runtime"
)

func CmbConverge(c *cli.Context) {

	var firstBootJsonChef string

	if runtime.GOOS == "windows" {
		firstBootJsonChef = path.Join("/etc/chef", "first-boot.json")
	} else {
		firstBootJsonChef = path.Join("/etc/chef", "first-boot.json")
	}

	if utils.Exists(firstBootJsonChef) {
		garbageOutput, _ := regexp.Compile("[\\[][^\\[|^\\]]*[\\]]\\s[A-Z]*:\\s")
		output, _ := regexp.Compile("Chef Run")
		cmd := exec.Command("chef-client", "-j", firstBootJsonChef)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Errorf("%s", err.Error())
		}
		ls := bufio.NewReader(stdout)
		err = cmd.Start()
		if err != nil {
			log.Errorf("%s", err.Error())
		}

		x := 0

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
			x = x + 1
			outputLine := garbageOutput.ReplaceAllString(string(line), "")
			if output.MatchString(outputLine) {
				log.Infof("%s", outputLine)
			} else {
				log.Debugf("%s", outputLine)
			}

		}
		err = cmd.Wait()
		if err != nil {
			log.Errorf("%s", err.Error())
		}
	} else {
		log.Fatalf("Make sure %s chef client configuration exists.", firstBootJsonChef)
	}

}
