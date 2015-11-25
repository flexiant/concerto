package utils

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckReturnCode(res int, mesg []byte) {
	if res >= 300 {
		// check if response is a web page.
		message := string(mesg[:])
		log.Debugf("Concerto API response: %s", message)

		webpageIdentified := "<html>"
		scrapResponse := "<title>(.*?)</title>"
		if strings.Contains(message, webpageIdentified) {

			re, err := regexp.Compile(scrapResponse)
			scrapped := re.FindStringSubmatch(message)

			if scrapped == nil || err != nil || len(scrapped) < 2 {
				// couldn't scrape, return generic error
				message = "Error executing operation"
			} else {
				// return scrapped response
				message = scrapped[1]
			}
		}
		// if it's not a web page, return raw message
		log.Fatal(fmt.Sprintf("There was an issue with your http request: status[%d] message [%s]", res, message))
	}
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

func GetHomeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME")
}

func GetBaseDir() string {
	baseDir := os.Getenv("MACHINE_DIR")
	if baseDir == "" {
		baseDir = GetHomeDir()
	}
	return baseDir
}

func GetConcertoDir() string {
	userName := GetUsername()
	if userName == "root" || Exists("/etc/concerto/client.xml") {
		return "/etc/concerto/"
	} else if userName == "Administrator" || Exists("c:\\concerto\\client.xml") || userName[len(userName)-1:] == "$" {
		return "c:\\concerto\\"
	} else {
		return filepath.Join(GetBaseDir(), ".concerto")
	}
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
