package config

import (
	"encoding/xml"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
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

	fileLocation := filepath.Join(utils.GetConcertoDir(), "client.xml")

	if utils.Exists(fileLocation) {
		var config *Config
		xmlFile, err := os.Open(fileLocation)
		if err != nil {
			return nil, err
		}
		defer xmlFile.Close()
		b, _ := ioutil.ReadAll(xmlFile)
		xml.Unmarshal(b, &config)

		if config.ApiEndpoint != "" && config.Certificate.Cert != "" {
			return config, nil
		} else {
			return nil, errors.New(fmt.Sprintf("Configuration File %s does not have valid format.", fileLocation))
		}

	} else if utils.Exists(filepath.Join(utils.GetConcertoDir(), "ssl", "cert.crt")) {

		certificate := Cert{
			filepath.Join(utils.GetConcertoDir(), "ssl", "cert.crt"),
			filepath.Join(utils.GetConcertoDir(), "ssl", "/private/cert.key"),
			filepath.Join(utils.GetConcertoDir(), "ssl", "ca_cert.pem"),
		}
		config := Config{}
		config.ApiEndpoint = os.Getenv("CONCERTO_ENDPOINT")
		config.Certificate = certificate
		log.Debugf("%#v", config)

		return &config, nil

	} else {
		return nil, errors.New(fmt.Sprintf("Configuration File %s does not exist.", fileLocation))
	}
}
