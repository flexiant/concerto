package format

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
	"text/tabwriter"
)

// TabFormatter prints items and lists
type TabFormatter struct {
	output *os.File
}

// NewTabFormatter creates a new TabFormatter
func NewTabFormatter(out *os.File) *TabFormatter {
	return &TabFormatter{
		output: out,
	}
}

// PrintItem prints an item
func (f *TabFormatter) PrintItem(item interface{}, header []string) {
	fmt.Print("TODO not implemented")
}

// PrintList prints an item
func (f *TabFormatter) PrintList(items [][]string, headers []string) {
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
func (f *TabFormatter) PrintError(context string, err error) {
	fmt.Print("probando error")

}
