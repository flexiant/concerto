package dispatcher

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils"
	"github.com/flexiant/concerto/webservice"
	"io/ioutil"
	"os"
)

const (
	characterizationsEndpoint = "blueprint/script_characterizations?type=%s"
	conclusionsEndpoint       = "blueprint/script_conclusions"
)

type ScriptCharacterization struct {
	Order      int               `json:"execution_order"`
	UUID       string            `json:"uuid"`
	Script     Script            `json:"script"`
	Parameters map[string]string `json:"parameter_values"`
}

type Script struct {
	Code            string   `json:"code"`
	UUID            string   `json:"uuid"`
	AttachmentPaths []string `json:"attachment_paths"`
}

type ScriptConclusion struct {
	UUID       string `json:"script_characterization_id"`
	Output     string `json:"output"`
	ExitCode   int    `json:"exit_code"`
	StartedAt  string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
}

type ScriptConclusionRoot struct {
	Root ScriptConclusion `json:"script_conclusion"`
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "boot",
			Usage:  "Executes scripts characterization associated to booting state of host",
			Action: cmdBoot,
		},
		{
			Name:   "operational",
			Usage:  "Executes scripts characterization associated to operational state of host",
			Action: cmdOperational,
		},
		{
			Name:   "shutdown",
			Usage:  "Executes scripts characterization associated to shutdown state of host",
			Action: cmdShutdown,
		},
	}
}

func executeScriptCharacterization(script ScriptCharacterization, directoryPath string) (conclusion ScriptConclusionRoot) {
	output, exitCode, startedAt, finishedAt := utils.ExecCode(script.Script.Code, directoryPath, script.Script.UUID)

	conclusion.Root.UUID = script.UUID
	conclusion.Root.Output = output
	conclusion.Root.ExitCode = exitCode
	conclusion.Root.StartedAt = startedAt.Format(utils.TimeStampLayout)
	conclusion.Root.FinishedAt = finishedAt.Format(utils.TimeStampLayout)

	return conclusion
}

func execute(phase string) {
	var scriptChars []ScriptCharacterization
	webservice, err := webservice.NewWebService()
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("Current Script Characterization %s", phase)
	data, err := webservice.Get(fmt.Sprintf(characterizationsEndpoint, phase))

	json.Unmarshal(data, &scriptChars)

	log.Debugf(string(data))

	err = json.Unmarshal(data, &scriptChars)
	if err != nil {
		log.Fatal(err)
	}
	scripts := ByOrder(scriptChars)

	for _, ex := range scripts {
		log.Infof("------------------------------------------------------------------------------------------------")
		path, err := ioutil.TempDir("", "concerto")
		if err != nil {
			log.Fatal(err)
		}

		os.Setenv("ATTACHMENT_DIR", fmt.Sprintf("%s/%s", path, "attachments"))

		log.Infof("UUID: %s", ex.UUID)
		log.Infof("Home Folder: %s", path)
		err = os.Mkdir(os.Getenv("ATTACHMENT_DIR"), 0777)
		if err != nil {
			log.Fatal(err)
		}

		// Seting up Enviroment Variables
		log.Infof("Enviroment Variables")
		for index, value := range ex.Parameters {
			os.Setenv(index, value)
			log.Infof("\t - %s=%s", index, value)
		}

		if len(ex.Script.AttachmentPaths) > 0 {
			log.Infof("Attachment Folder: %s", os.Getenv("ATTACHMENT_DIR"))
			// Downloading Attachements
			log.Infof("Attachments")
			if err != nil {
				log.Fatal(err)
			}
			for _, endpoint := range ex.Script.AttachmentPaths {
				filename, err := webservice.GetFile(endpoint, os.Getenv("ATTACHMENT_DIR"))
				if err != nil {
					log.Fatal(err)
				}
				log.Infof("\t - %s --> %s", endpoint, filename)
			}
		}

		json, err := json.Marshal(executeScriptCharacterization(ex, path))
		if err != nil {
			log.Fatal(err)
		}

		err = webservice.Post(conclusionsEndpoint, json)
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("------------------------------------------------------------------------------------------------")
	}
}

func cmdBoot(c *cli.Context) {
	execute("boot")
}

func cmdOperational(c *cli.Context) {
	execute("operational")
}

func cmdShutdown(c *cli.Context) {
	execute("shutdown")
}
