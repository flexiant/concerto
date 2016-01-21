package config

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/flexiant/concerto/utils"
)

// Config stores Concerto CLI configuration items
type Config struct {
	XMLName     xml.Name `xml:"concerto"`
	ApiEndpoint string   `xml:"server,attr"`
	LogFile     string   `xml:"log_file,attr"`
	LogLevel    string   `xml:"log_level,attr"`
	Certificate Cert     `xml:"ssl"`
}

// Cert stores certificates to use against Concerto API
type Cert struct {
	Cert string `xml:"cert,attr"`
	Key  string `xml:"key,attr"`
	Ca   string `xml:"server_ca,attr"`
}

// cachedConf makes sure we read configuratino only once
var cachedConf *Config

// ConcertoServerConfiguration returns Concerto Server Configuration
func ConcertoServerConfiguration() (*Config, error) {

	if cachedConf != nil {
		return cachedConf, nil
	}

	fileLocation := utils.GetConcertoConfig()

	if utils.Exists(fileLocation) {
		var config *Config

		xmlFile, err := os.Open(fileLocation)
		if err != nil {
			return nil, err
		}
		defer xmlFile.Close()
		b, _ := ioutil.ReadAll(xmlFile)
		xml.Unmarshal(b, &config)

		config.ApiEndpoint = utils.GetConcertoEndpoint()
		config.Certificate.Ca = utils.GetConcertoCACert()
		config.Certificate.Cert = utils.GetConcertoClientCert()
		config.Certificate.Key = utils.GetConcertoClientKey()

		if config.ApiEndpoint != "" && config.Certificate.Cert != "" {
			return config, nil
		}
		return nil, fmt.Errorf("Configuration File %s does not have valid format.", fileLocation)

	} else if utils.Exists(utils.GetConcertoClientCert()) {

		certificate := Cert{
			utils.GetConcertoClientCert(),
			utils.GetConcertoClientKey(),
			utils.GetConcertoCACert(),
		}
		config := Config{}
		config.ApiEndpoint = utils.GetConcertoEndpoint()
		config.Certificate = certificate

		return &config, nil

	}

	return nil, fmt.Errorf("Configuration File %s does not exist.", fileLocation)
}
