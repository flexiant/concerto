package utils

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"
)

const (
	TimeStampLayout = "2006-01-02T15:04:05.000000-07:00"
)

func extractExitCode(err error) int {
	if err != nil {
		exiterr := err.(*exec.ExitError)
		status := exiterr.Sys().(syscall.WaitStatus)
		return status.ExitStatus()
	} else {
		return 0
	}
}

func ExecCode(code string, path string, filename string) (output string, exitCode int, startedAt time.Time, finishedAt time.Time) {
	tmp, err := os.Create(fmt.Sprintf("%s/%s", path, filename))
	if err != nil {
		log.Fatalf("Error creating temp file : ", err)
	}

	defer os.Remove(tmp.Name())
	defer tmp.Close()

	_, err = tmp.WriteString(code)
	if err != nil {
		log.Fatalf("Error writing to file : ", err)
	}

	os.Chmod(tmp.Name(), 0777)
	if err != nil {
		log.Fatalf("Error changing permision to file : ", err)
	}

	return RunCmd(tmp.Name())
}

func RunCmd(command string) (output string, exitCode int, startedAt time.Time, finishedAt time.Time) {

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		log.Infof("Command: %s %s", "cmd /C ", command)
		cmd = exec.Command("cmd", command)
	} else {
		log.Infof("Command: %s %s", "/bin/sh", command)
		cmd = exec.Command("/bin/sh", command)
	}

	startedAt = time.Now()
	bytes, err := cmd.CombinedOutput()
	finishedAt = time.Now()
	output = strings.TrimSpace(string(bytes))
	exitCode = extractExitCode(err)

	log.Debugf("Starting Time: %s", startedAt.Format(TimeStampLayout))
	log.Debugf("End Time: %s", finishedAt.Format(TimeStampLayout))
	log.Debugf("Output")
	log.Debugf("")
	log.Debugf("%s", output)
	log.Debugf("")
	log.Infof("Exit Code: %d", exitCode)
	return
}
