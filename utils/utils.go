package utils

import (
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
	if GetUsername() == "root" {
		return "/etc/concerto/"
	} else if GetUsername() == "Administrator" {
		return "c:\\concerto\\"
	} else {
		return filepath.Join(GetBaseDir(), ".concerto")
	}
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
