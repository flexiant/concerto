package utils

import (
	"crypto/x509"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

const windowsServerConfigFile = "c:\\concerto\\client.xml"
const nixServerConfigFile = "/etc/concerto/client.xml"
const defaultConcertoEndpoint = "https://clients.concerto.io:886/"

// Config stores configuration file contents
type Config struct {
	XMLName      xml.Name `xml:"concerto"`
	APIEndpoint  string   `xml:"server,attr"`
	LogFile      string   `xml:"log_file,attr"`
	LogLevel     string   `xml:"log_level,attr"`
	Certificate  Cert     `xml:"ssl"`
	ConfLocation string
	ConfFile     string
	IsHost       bool
	ConcertoURL  string
}

// Cert stores cert files location
type Cert struct {
	Cert string `xml:"cert,attr"`
	Key  string `xml:"key,attr"`
	Ca   string `xml:"server_ca,attr"`
}

var cachedConfig *Config

// GetConcertoConfig returns concerto configuration
func GetConcertoConfig() (*Config, error) {
	if cachedConfig == nil {
		return nil, fmt.Errorf("Configuration hasn't been initialized")
	}
	return cachedConfig, nil
}

// InitializeConcertoConfig creates the concerto configuration structure
func InitializeConcertoConfig(c *cli.Context) (*Config, error) {
	log.Debug("InitializeConcertoConfig")
	if cachedConfig != nil {
		return cachedConfig, nil
	}

	// where config file must me
	cachedConfig = &Config{}
	err := cachedConfig.evaluateConcertoConfigFile(c)
	if err != nil {
		return nil, err
	}

	// read config contents
	log.Debugf("Reading configuration from %s", cachedConfig.ConfFile)
	err = cachedConfig.readConcertoConfig(c)
	if err != nil {
		return nil, err
	}

	// add login URL. Needed for setup
	err = cachedConfig.readConcertoURL()
	if err != nil {
		return nil, err
	}

	// check if isHost. Needed to show appropiate options
	err = cachedConfig.evaluateCertificate()
	if err != nil {
		return nil, err
	}

	debugShowConfig()
	return cachedConfig, nil
}

func debugShowConfig() {
	if log.GetLevel() < log.DebugLevel {
		return
	}

	if cachedConfig == nil {
		log.Debug("Concerto configuration not loaded")
	}

	debugStruct("", *cachedConfig)
	// c := reflect.ValueOf(*cachedConfig)
	// for i := 0; i < c.NumField(); i++ {
	// 	if c.Type().Field(i).Type.String() != "xml.Name" {
	// 		log.WithField(c.Type().Field(i).Name, c.Field(i).Interface()).Debug("Configuration item")
	// 	}
	// }
}

// debugStruct iterates struct and show in debug console all items and subitems
func debugStruct(prefix string, item interface{}) {
	c := reflect.ValueOf(item)
	for i := 0; i < c.NumField(); i++ {
		if c.Type().Field(i).Type.String() != "xml.Name" {

			name := c.Type().Field(i).Name
			value := c.Field(i).Interface()

			// if value is struct, iterate with recursion
			if c.Type().Field(i).Type.Kind() == reflect.Struct {
				debugStruct(name, value)
			} else {
				if prefix != "" {
					name = fmt.Sprintf("%s.%s", prefix, name)
				}
				log.WithField(name, value).Debug("Configuration item")
			}
		}
	}
}

// IsConfigReady returns whether configurations items are filled
func (config *Config) IsConfigReady() bool {
	if config.APIEndpoint == "" ||
		config.Certificate.Cert == "" ||
		config.Certificate.Key == "" ||
		config.Certificate.Ca == "" {
		return false
	}
	return true
}

// IsConfigReadySetup returns whether we can use setup command
func (config *Config) IsConfigReadySetup() bool {
	return config.ConcertoURL != ""
}

// readConcertoConfig reads Concerto config file located at fileLocation
func (config *Config) readConcertoConfig(c *cli.Context) error {
	log.Debug("Reading Concerto Configuration")
	if FileExists(config.ConfFile) {
		// file exists, read it's contents

		xmlFile, err := os.Open(config.ConfFile)
		if err != nil {
			return err
		}
		defer xmlFile.Close()
		b, err := ioutil.ReadAll(xmlFile)
		if err != nil {
			return fmt.Errorf("Configuration File %s couldn't be read.", config.ConfFile)
		}

		if err = xml.Unmarshal(b, &config); err != nil {
			return fmt.Errorf("Configuration File %s does not have valid XML format.", config.ConfFile)
		}

	} else {
		log.Debugf("Configuration File %s does not exist. Reading environment variables", config.ConfFile)
	}

	// overwrite with environment/arguments vars
	if overwEP := c.String("concerto-endpoint"); overwEP != "" {
		log.Debug("Concerto APIEndpoint taken from env/args")
		config.APIEndpoint = overwEP
	}

	if overwCert := c.String("client-cert"); overwCert != "" {
		log.Debug("Certificate path taken from env/args")
		config.Certificate.Cert = overwCert
	}

	if overwKey := c.String("client-key"); overwKey != "" {
		log.Debug("Certificate key path taken from env/args")
		config.Certificate.Key = overwKey
	}

	if overwCa := c.String("ca-cert"); overwCa != "" {
		log.Debug("CA certificate path taken from env/args")
		config.Certificate.Ca = overwCa
	}

	// if endpoint empty set default
	// we can't set the default from flags, because it would overwrite config file
	if config.APIEndpoint == "" {
		config.APIEndpoint = defaultConcertoEndpoint
	}

	return nil
}

// evaluateConcertoConfigFile returns path to concerto config file
func (config *Config) evaluateConcertoConfigFile(c *cli.Context) error {
	log.Debug("evaluateConcertoConfigFile")

	if configFile := c.String("concerto-config"); configFile != "" {

		log.Debug("Concerto configuration file location taken from env/args")
		config.ConfFile = configFile

	} else {

		currUser, err := user.Current()
		if err != nil {
			log.Debugf("Couldn't use os.user to get user details: %s", err.Error())
			dir, err := homedir.Dir()
			if err != nil {
				return fmt.Errorf("Couldn't get home dir for current user: %s", err.Error())
			}
			currUser = &user.User{
				Username: getUsername(),
				HomeDir:  dir,
			}
		}

		if runtime.GOOS == "windows" {
			currUser.Username = currUser.Username[strings.LastIndex(currUser.Username, "\\")+1:]
			log.Debugf("Windows username is %s", currUser.Username)

			if (currUser.Gid == "S-1-5-32-544" || isWinAdministrator(currUser.Username)) && FileExists(windowsServerConfigFile) {
				log.Debug("Current user is administrator, setting config file as %s", windowsServerConfigFile)
				config.ConfFile = windowsServerConfigFile
			} else {
				// User mode Windows
				log.Debugf("Current user is regular user: %s", currUser.Username)
				config.ConfFile = filepath.Join(currUser.HomeDir, ".concerto/client.xml")
			}
		} else {
			// Server mode *nix
			if currUser.Uid == "0" || currUser.Username == "root" && FileExists(nixServerConfigFile) {
				config.ConfFile = nixServerConfigFile
			} else {
				// User mode *nix
				config.ConfFile = filepath.Join(currUser.HomeDir, ".concerto/client.xml")
			}
		}
	}
	config.ConfLocation = path.Dir(config.ConfFile)
	return nil
}

// getUsername gets username by env variable.
// os.user is dependant on cgo, so cross compiling won't work
func getUsername() string {
	log.Debug("getUsername")
	u := "unknown"
	osUser := ""

	switch runtime.GOOS {
	case "darwin", "linux", "solaris":
		osUser = os.Getenv("USER")
	case "windows":
		osUser = os.Getenv("USERNAME")

		// remove domain
		osUser = osUser[strings.LastIndex(osUser, "\\")+1:]
		log.Debugf("Windows user has been transformed into %s", osUser)

		// HACK ugly ... if localized administrator, translate to administrator
		if isWinAdministrator(osUser) {
			osUser = "Administrator"
		}
	}

	if osUser != "" {
		u = osUser
	}
	return u
}

func isWinAdministrator(user string) bool {
	return user == "Järjestelmänvalvoja" ||
		user == "Administrateur" ||
		user == "Rendszergazda" ||
		user == "Administrador" ||
		user == "Администратор" ||
		user == "Administratör" ||
		user == "Administrator" ||
		user == "SYSTEM"
}

// readConcertoURL reads URL from CONCERTO_URL envrionment or calculates using API URL
func (config *Config) readConcertoURL() error {

	if config.ConcertoURL != "" {
		return nil
	}

	if overwURL := os.Getenv("CONCERTO_URL"); overwURL != "" {
		config.ConcertoURL = overwURL
		log.Debug("Concerto URL taken from CONCERTO_URL")
		return nil
	}

	cURL, err := url.Parse(config.APIEndpoint)
	if err != nil {
		return err
	}

	tokenHost := strings.Split(cURL.Host, ":")
	tokenFqdn := strings.Split(tokenHost[0], ".")

	if !strings.Contains(cURL.Host, "staging") {
		tokenFqdn[0] = "start"
	}

	config.ConcertoURL = fmt.Sprintf("%s://%s/", cURL.Scheme, strings.Join(tokenFqdn, "."))
	return nil
}

// evaluateCertificate determines if a certificate has been issued for a host
func (config *Config) evaluateCertificate() error {

	if FileExists(config.Certificate.Cert) {

		data, err := ioutil.ReadFile(config.Certificate.Cert)
		if err != nil {
			return err
		}

		block, _ := pem.Decode(data)

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return err
		}

		if len(cert.Subject.OrganizationalUnit) > 0 {
			if cert.Subject.OrganizationalUnit[0] == "Hosts" {
				config.IsHost = true
				return nil
			}
		} else if len(cert.Issuer.Organization) > 0 {
			if cert.Issuer.Organization[0] == "Tapp" {
				config.IsHost = true
				return nil
			}
		}
	}
	config.IsHost = false
	return nil
}
