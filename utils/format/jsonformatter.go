package format

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io"
	"os"
)

// JSONFormatter prints items and lists in JSON format
type JSONFormatter struct {
	output io.Writer
}

// NewJSONFormatter creates a new JSONFormatter
func NewJSONFormatter(out io.Writer) *JSONFormatter {
	log.Debug("Creating JSON formatter")
	return &JSONFormatter{
		output: out,
	}
}

// PrintItem prints an item
func (f *JSONFormatter) PrintItem(item interface{}) error {
	b, err := json.Marshal(item)
	if err != nil {
		return err
	}
	f.output.Write(b)
	fmt.Fprintf(f.output, "\n")

	return nil
}

// PrintList prints item list
func (f *JSONFormatter) PrintList(items interface{}) error {
	log.Debug("PrintList")
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	f.output.Write(b)
	fmt.Fprintf(f.output, "\n")

	return nil
}

// PrintError prints an error
func (f *JSONFormatter) PrintError(context string, err error) {
	// TODO JSON
	fmt.Printf("ERROR: %s\n -> %s", context, err)
}

// PrintFatal prints an error and exists
func (f *JSONFormatter) PrintFatal(context string, err error) {
	// TODO JSON
	f.PrintError(context, err)
	os.Exit(1)
}
