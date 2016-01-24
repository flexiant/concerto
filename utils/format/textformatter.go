package format

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"text/tabwriter"
)

// TextFormatter prints items and lists
type TextFormatter struct {
	output *os.File
}

// NewTextFormatter creates a new TextFormatter
func NewTextFormatter(out *os.File) *TextFormatter {
	return &TextFormatter{
		output: out,
	}
}

// PrintItem prints an item
func (f *TextFormatter) PrintItem(item interface{}, header []string) {
	fmt.Print("TODO not implemented")
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
	fmt.Print("probando error")

}
