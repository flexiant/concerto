package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
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

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
