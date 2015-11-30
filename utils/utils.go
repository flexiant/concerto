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

func ScrapeErrorMessage(message string, regExpression string) string {

	re, err := regexp.Compile(regExpression)
	scrapped := re.FindStringSubmatch(message)

	if scrapped == nil || err != nil || len(scrapped) < 2 {
		// couldn't scrape, return generic error
		message = "Error executing operation"
	} else {
		// return scrapped response
		message = scrapped[1]
	}

	return message
}

func CheckReturnCode(res int, mesg []byte) {
	if res >= 300 {

		message := string(mesg[:])
		log.Debugf("Concerto API response: %s", message)

		f := func(c rune) bool {
			return c == ',' || c == ':' || c == '{' || c == '}' || c == '"' || c == ']' || c == '['
		}

		// check if response is a web page.
		if strings.Contains(message, "<html>") {
			scrapResponse := "<title>(.*?)</title>"
			message = ScrapeErrorMessage(message, scrapResponse)
		} else if strings.Contains(message, "{\"errors\":{") {
			scrapResponse := "{\"errors\":(.*?)}"

			message = ScrapeErrorMessage(message, scrapResponse)
			result := strings.Split(message, ",")
			if result != nil {
				message = result[0]
			}
			// Separate into fields with func.
			fields := strings.FieldsFunc(message, f)
			message = strings.Join(fields[:], " ")

		} else if strings.Contains(message, "{\"error\":") {
			scrapResponse := "{\"error\":\"(.*?)\"}"
			message = ScrapeErrorMessage(message, scrapResponse)
		}

		// if it's not a web page or json-formatted message, return the raw message
		log.Fatal(fmt.Sprintf("There was an issue with your http request: status [%d] message [%s]", res, message))
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
