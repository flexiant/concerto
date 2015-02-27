package config

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/flexiant/concerto/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	XMLName     xml.Name `xml:"tapp"`
	ApiEndpoint string   `xml:"server,attr"`
	LogFile     string   `xml:"log_file,attr"`
	LogLevel    string   `xml:"log_level,attr"`
	Certificate Cert     `xml:"ssl"`
}

type Cert struct {
	Cert string `xml:"cert,attr"`
	Key  string `xml:"key,attr"`
	Ca   string `xml:"server_ca,attr"`
}

// Returns Concerto Server Configuration
func ConcertoServerConfiguration() (*Config, error) {

	var fileLocation string

	if utils.GetUsername() == "root" {
		fileLocation = "/etc/tapp/client.xml"
	} else {
		fileLocation = filepath.Join(utils.GetBaseDir(), ".krane", "client.xml")
	}

	if utils.Exists(fileLocation) {
		var config *Config
		xmlFile, err := os.Open(fileLocation)
		if err != nil {
			return nil, err
		}
		defer xmlFile.Close()
		b, _ := ioutil.ReadAll(xmlFile)
		xml.Unmarshal(b, &config)
		return config, nil
	} else {
		return nil, errors.New(fmt.Sprintf("Configuration File %s does not exist.", fileLocation))
	}
}
