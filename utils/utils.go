package utils

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckReturnCode(res int) {
	if res >= 300 {
		log.Fatal(fmt.Sprintf("There was an issue with your http request; error code: %d", res))
	} else {
		log.Info(fmt.Sprintf("The command executed successfully; http response code: %d", res))
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
