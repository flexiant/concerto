package cmd

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils/format"
	"os"
	"runtime"
	"strings"
)

// debugCmdFuncInfo writes context info about the calling function
func debugCmdFuncInfo(c *cli.Context) {
	if log.GetLevel() < log.DebugLevel {
		return
	}

	// get function name
	dbgMsg := ""
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		dbgMsg = runtime.FuncForPC(pc).Name()
		i := strings.LastIndex(dbgMsg, "/")
		if i != -1 {
			dbgMsg = dbgMsg[i+1:]
		}
	} else {
		dbgMsg = "<unknown function name>"
	}
	dbgMsg = fmt.Sprintf("func %s", dbgMsg)

	// get used flags
	for _, flag := range c.FlagNames() {
		dbgMsg = fmt.Sprintf("%s\n\t%s=%+v", dbgMsg, flag, c.Generic(flag))
	}
	log.Debugf(dbgMsg)
}

// checkRequiredFlags checks for required flags, and show usage if requirements not met
func checkRequiredFlags(c *cli.Context, flags []string, f format.Formatter) {
	missing := ""
	for _, flag := range flags {
		if !c.IsSet(flag) {
			missing = fmt.Sprintf("%s\n\t--%s", missing, flag)
		}
	}

	if missing != "" {
		f.PrintError("Incorrect usage.", fmt.Errorf("Mandatory parameters missing: %s\n", missing))
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(2)
	}
}

// checkRequiredFlagsOr checks that at least one of required flags is present, and show usage if requirements not met
func checkRequiredFlagsOr(c *cli.Context, flags []string, f format.Formatter) {
	missing := ""
	for _, flag := range flags {
		if c.IsSet(flag) {
			return
		}
		missing = fmt.Sprintf("%s\n\t--%s", missing, flag)
	}

	f.PrintError("Incorrect usage.", fmt.Errorf("Please use either parameter: %s\n", missing))
	cli.ShowCommandHelp(c, c.Command.Name)
	os.Exit(2)
}
