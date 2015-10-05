package utils

import (
	"bufio"
	"bytes"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io"
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
		switch err.(type) {
		case *exec.ExitError:
			return err.(*exec.ExitError).Sys().(syscall.WaitStatus).ExitStatus()
		case *os.PathError:
			return 127
		}
	}
	return 0
}

func ExecCode(code string, path string, filename string) (output string, exitCode int, startedAt time.Time, finishedAt time.Time) {
	var err error
	var tmp *os.File

	if runtime.GOOS == "windows" {
		tmp, err = os.Create(fmt.Sprintf("%s/%s.bat", path, filename))
	} else {
		tmp, err = os.Create(fmt.Sprintf("%s/%s", path, filename))
	}

	if err != nil {
		log.Fatalf("Error creating temp file : ", err)
	}

	defer tmp.Close()

	_, err = tmp.WriteString(code)
	if err != nil {
		log.Fatalf("Error writing to file : ", err)
	}

	os.Chmod(tmp.Name(), 0777)
	if err != nil {
		log.Fatalf("Error changing permision to file : ", err)
	}

	return RunFile(tmp.Name())
}

func RunFile(command string) (output string, exitCode int, startedAt time.Time, finishedAt time.Time) {

	var cmd *exec.Cmd

	var b bytes.Buffer
	buffer := bufio.NewWriter(&b)

	if runtime.GOOS == "windows" {
		log.Infof("Command: %s", command)
		cmd = exec.Command("cmd", "/C", command)
	} else {
		log.Infof("Command: %s %s", "/bin/sh", command)
		cmd = exec.Command("/bin/sh", command)
	}

	stdout, err := cmd.StdoutPipe()
	CheckError(err)

	stderr, err := cmd.StderrPipe()
	CheckError(err)

	multi := io.MultiReader(stdout, stderr)

	startedAt = time.Now()
	err = cmd.Start()
	CheckError(err)

	io.Copy(buffer, multi)

	//go io.Copy(buffer, stderr)
	//go io.Copy(buffer, stdout)

	err = cmd.Wait()
	finishedAt = time.Now()
	exitCode = extractExitCode(err)

	err = buffer.Flush()
	CheckError(err)

	log.Debugf("Starting Time: %s", startedAt.Format(TimeStampLayout))
	log.Debugf("End Time: %s", finishedAt.Format(TimeStampLayout))
	log.Debugf("Output")
	log.Debugf("")
	log.Debugf("%s", b.String())
	log.Debugf("")
	log.Infof("Exit Code: %d", exitCode)
	return
}

func RunCmd(command string) (output string, exitCode int, startedAt time.Time, finishedAt time.Time) {

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		log.Infof("Command: %s", command)
		cmd = exec.Command("cmd", "/C", command)
	} else {
		log.Infof("Command: %s %s", "/bin/sh -c", command)
		cmd = exec.Command("/bin/sh", "-c", command)
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
