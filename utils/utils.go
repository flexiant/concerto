package utils

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
)

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

func RunCmd(command string) string {
	log.Debugln(command)
	out, err := exec.Command("/bin/sh", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Infoln(out)
	return string(out)
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
	return filepath.Join(GetBaseDir(), ".krane")
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
