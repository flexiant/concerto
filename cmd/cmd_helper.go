package cmd

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/flexiant/concerto/utils/format"
	"os"
	"reflect"
	"runtime"
	"strings"
)

//
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
			missing = fmt.Sprintf("%s\n\t--%s\n", missing, flag)
		}
	}

	if missing != "" {
		f.PrintError("Incorrect usage.", fmt.Errorf("Mandatory parameters missing: %s", missing))
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(2)
	}
}

// flagConvertParams converts cli parameters in API callable params
func flagConvertParams(c *cli.Context) *map[string]string {
	v := make(map[string]string)
	for _, flag := range c.FlagNames() {
		if c.IsSet(flag) {
			v[flag] = c.String(flag)
		}
	}
	return &v
}

func itemConvertParams(item interface{}) (*map[string]string, error) {

	it := reflect.ValueOf(item)
	nf := it.NumField()
	v := make(map[string]string)

	for i := 0; i < nf; i++ {
		v[it.Type().Field(i).Name] = fmt.Sprintf("%s", it.Field(i).Interface())
		// if value, ok :=  it.Field(i).Interface().(string); ok {
		// 	v[it.Type().Field(i).Name] = value
		// } else {
		// 	return nil, fmt.Errorf("Interface couldn't be converted to map of strings. Field: %s", it.Type().Field(i).Name)
		// }
	}
	return &v, nil
}
