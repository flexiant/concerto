package format

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io"
	"os"
	"text/tabwriter"
)

// TextFormatter prints items and lists
type TextFormatter struct {
	output io.Writer
}

// NewTextFormatter creates a new TextFormatter
func NewTextFormatter(out io.Writer) *TextFormatter {
	return &TextFormatter{
		output: out,
	}
}

// PrintItem prints an item
func (f *TextFormatter) PrintItem(item []string, header []string) {
	log.Debug("PrintItem")
	w := tabwriter.NewWriter(f.output, 15, 1, 3, ' ', 0)
	for i := range item {
		fmt.Fprintf(w, "%s:\t%s\n", header[i], item[i])
	}
	fmt.Fprintln(w)
	w.Flush()
}

// PrintList prints an item
func (f *TextFormatter) PrintList(items [][]string, headers []string) {
	log.Debug("PrintList")
	w := tabwriter.NewWriter(f.output, 15, 1, 3, ' ', 0)

	for _, h := range headers {
		fmt.Fprintf(w, "%s\t", h)
	}
	fmt.Fprintln(w)

	for _, item := range items {
		for _, column := range item {
			fmt.Fprintf(w, "%s\t", column)
		}
		fmt.Fprintln(w)
	}

	w.Flush()

}

// PrintError prints an error
func (f *TextFormatter) PrintError(context string, err error) {
	fmt.Printf("ERROR: %s\n -> %s", context, err)
}

// PrintFatal prints an error and exists
func (f *TextFormatter) PrintFatal(context string, err error) {
	f.PrintError(context, err)
	os.Exit(1)
}
