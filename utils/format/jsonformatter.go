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

// JSONMessage hosts generic messages
type JSONMessage struct {
	Type    string `json:"type"`
	Context string `json:"context,omitempty"`
	Message string `json:"message"`
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

	msg := JSONMessage{
		Type:    "Error",
		Context: context,
		Message: err.Error(),
	}

	msgJSON, err := json.Marshal(msg)
	if err != nil {
		// fallback to hand made message
		msgJSON = []byte(fmt.Sprintf("(Formatting error, cannot show JSON)\n %s -> %s \n", context, err))
	}

	f.output.Write(msgJSON)
	fmt.Fprintf(f.output, "\n")
}

// PrintFatal prints an error and exists
func (f *JSONFormatter) PrintFatal(context string, err error) {
	// TODO JSON
	f.PrintError(context, err)
	os.Exit(1)
}
