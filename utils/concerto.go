package utils

import (
	"crypto/x509"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
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

func GetConfig(fileLocation string) (*Config, error) {

	if Exists(fileLocation) {
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
	}
	return nil, errors.New(fmt.Sprintf("Configuration File %s does not exist.", fileLocation))
}

func GetUsername() string {
	u := "unknown"
	osUser := ""

	switch runtime.GOOS {
	case "darwin", "linux":
		osUser = os.Getenv("USER")
	case "windows":
		osUser = os.Getenv("USERNAME")
	}

	if osUser != "" {
		u = osUser
	}

	return u
}

func GetBaseDir() string {
	baseDir := os.Getenv("MACHINE_DIR")
	if baseDir == "" {
		if runtime.GOOS == "windows" {
			return os.Getenv("USERPROFILE")
		}
		return os.Getenv("HOME")
	}
	return baseDir
}

func GetConcertoDir() (path string, server bool) {
	userName := GetUsername()
	if userName == "root" || Exists("/etc/concerto/client.xml") {
		return "/etc/concerto/", true
	} else if userName == "Administrator" || Exists("c:\\concerto\\client.xml") || userName[len(userName)-1:] == "$" {
		return "c:\\concerto\\", true
	} else {
		return filepath.Join(GetBaseDir(), ".concerto"), false
	}
}

func IsClientCertificate(filename string) bool {
	if Exists(filename) {
		data, err := ioutil.ReadFile(filename)
		CheckError(err)
		block, _ := pem.Decode(data)

		cert, err := x509.ParseCertificate(block.Bytes)
		CheckError(err)

		fmt.Printf("%#v")
		if len(cert.Subject.OrganizationalUnit) > 0 {
			if cert.Subject.OrganizationalUnit[0] == "Hosts" {
				return true
			} else if cert.Issuer.Organization[0] == "Tapp" {
				return true
			}
		}
	}
	return false
}

func GetConcertoCACert() string {
	enviromentVariable, enviromentVariableExists := LookupEnv("CONCERTO_CA_CERT")
	directoryPath, server := GetConcertoDir()
	if enviromentVariableExists {
		return enviromentVariable
	} else {
		if server {
			if Exists(filepath.Join(directoryPath, "client.xml")) {
				config, err := GetConfig(filepath.Join(directoryPath, "client.xml"))
				CheckError(err)
				return config.Certificate.Ca
			} else {
				return filepath.Join(directoryPath, "client_ssl", "ca_cert.pem")
			}

		} else {
			return filepath.Join(directoryPath, "ssl", "ca_cert.pem")
		}
	}
}

func GetConcertoClientCert() string {
	enviromentVariable, enviromentVariableExists := LookupEnv("CONCERTO_CLIENT_CERT")
	directoryPath, server := GetConcertoDir()
	if enviromentVariableExists {
		return enviromentVariable
	} else {
		if server {
			if Exists(filepath.Join(directoryPath, "client.xml")) {
				config, err := GetConfig(filepath.Join(directoryPath, "client.xml"))
				CheckError(err)
				return config.Certificate.Cert
			} else {
				return filepath.Join(directoryPath, "client_ssl", "cert.pem")
			}

		} else {
			return filepath.Join(directoryPath, "ssl", "cert.crt")
		}
	}
}

func GetConcertoClientKey() string {
	enviromentVariable, enviromentVariableExists := LookupEnv("CONCERTO_CLIENT_KEY")
	directoryPath, server := GetConcertoDir()
	if enviromentVariableExists {
		return enviromentVariable
	} else {
		if server {
			if Exists(filepath.Join(directoryPath, "client.xml")) {
				config, err := GetConfig(filepath.Join(directoryPath, "client.xml"))
				CheckError(err)
				return config.Certificate.Key
			} else {
				return filepath.Join(directoryPath, "client_ssl", "private", "key.pem")
			}

		} else {
			return filepath.Join(directoryPath, "ssl", "private", "cert.key")
		}
	}
}

func GetConcertoConfig() string {
	enviromentVariable, enviromentVariableExists := LookupEnv("CONCERTO_CONFIG")
	directoryPath, _ := GetConcertoDir()
	if enviromentVariableExists {
		return enviromentVariable
	} else {
		return filepath.Join(directoryPath, "client.xml")
	}
}

func LookupEnv(key string) (string, bool) {
	if len(os.Getenv(key)) > 0 {
		return os.Getenv(key), true
	}
	return "", false
}

func GetConcertoEndpoint() string {
	enviromentVariable, enviromentVariableExists := LookupEnv("CONCERTO_ENDPOINT")
	directoryPath, _ := GetConcertoDir()
	if enviromentVariableExists {
		return enviromentVariable
	} else {
		if Exists(filepath.Join(directoryPath, "client.xml")) {
			config, err := GetConfig(filepath.Join(directoryPath, "client.xml"))
			CheckError(err)
			return config.ApiEndpoint
		}
	}
	return ""
}
