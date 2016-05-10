package converge

import (
	"bufio"
	"errors"
	"io"
	"os/exec"
	"path"
	"regexp"
	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
)

func CmbConverge(c *cli.Context) error {

	var firstBootJsonChef string

	if runtime.GOOS == "windows" {
		firstBootJsonChef = path.Join("c:\\chef", "first-boot.json")
	} else {
		firstBootJsonChef = path.Join("/etc/chef", "first-boot.json")
	}

	if utils.FileExists(firstBootJsonChef) {
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
	return nil
}
