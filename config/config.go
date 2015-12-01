package config

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/flexiant/concerto/utils"
)

type Config struct {
	XMLName     xml.Name `xml:"concerto"`
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
		} else {
			return nil, errors.New(fmt.Sprintf("Configuration File %s does not have valid format.", fileLocation))
		}

	} else if utils.Exists(utils.GetConcertoClientCert()) {

		certificate := Cert{
			utils.GetConcertoClientCert(),
			utils.GetConcertoClientKey(),
			utils.GetConcertoCACert(),
		}
		config := Config{}
		config.ApiEndpoint = utils.GetConcertoEndpoint()
		config.Certificate = certificate
		fmt.Printf("%#v", config)

		return &config, nil

	} else {
		return nil, errors.New(fmt.Sprintf("Configuration File %s does not exist.", fileLocation))
	}
	return nil, errors.New("Can not load configuration")
}
