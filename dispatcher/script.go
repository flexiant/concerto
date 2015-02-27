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

type ConclusionWrapper struct {
	Conclusion ScriptConclusion `json:"script_conclusion"`
}

func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "boot",
			Usage:  "Executes scripts characterization associated to booting state of host",
			Action: cmbBoot,
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

// func execScriptCharacterization(script ScriptCharacterization, directoryPath string) (output ScriptConclusion) {
// 	output.UUID = script.UUID
// 	output.Output, output.ExitCode, output.StartedAt, output.FinishedAt = utils.ExecCode(script.Script.Code, directoryPath, script.Script.UUID)
// 	return output
// }

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
		log.Infof("Attachment Folder: %s", os.Getenv("ATTACHMENT_DIR"))

		// Seting up Enviroment Variables
		log.Infof("Enviroment Variables")
		for index, value := range ex.Parameters {
			os.Setenv(index, value)
			log.Debugf("\t - %s=%s", index, value)
		}

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
		utils.ExecCode(ex.Script.Code, path, ex.Script.UUID)
		log.Infof("------------------------------------------------------------------------------------------------")
	}
}

func cmbBoot(c *cli.Context) {
	execute("boot")
}

func cmdOperational(c *cli.Context) {
	execute("operational")
}

func cmdShutdown(c *cli.Context) {
	execute("shutdown")
}
