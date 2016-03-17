package format

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"
)

// TextFormatter prints items and lists
type TextFormatter struct {
	output io.Writer
}

// NewTextFormatter creates a new TextFormatter
func NewTextFormatter(out io.Writer) *TextFormatter {
	log.Debug("Creating Text formatter")
	return &TextFormatter{
		output: out,
	}
}

// PrintItem prints an item
func (f *TextFormatter) PrintItem(item interface{}) error {

	it := reflect.ValueOf(item)
	nf := it.NumField()

	w := tabwriter.NewWriter(f.output, 15, 1, 3, ' ', 0)
	for i := 0; i < nf; i++ {
		// TODO not the best way to use reflection. Check this later
		switch it.Field(i).Type().String() {
		case "json.RawMessage":
			fmt.Fprintf(w, "%s:\t%s\n", it.Type().Field(i).Tag.Get("header"), it.Field(i).Interface())
		case "*json.RawMessage":
			fmt.Fprintf(w, "%s:\t%s\n", it.Type().Field(i).Tag.Get("header"), it.Field(i).Elem())
		default:
			fmt.Fprintf(w, "%s:\t%+v\n", it.Type().Field(i).Tag.Get("header"), it.Field(i).Interface())
		}
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}

// PrintList prints item list
func (f *TextFormatter) PrintList(items interface{}) error {

	// should be an array
	its := reflect.ValueOf(items)
	t := its.Type().Kind()
	if t != reflect.Slice {
		return fmt.Errorf("Couldn't print list. Expected slice, but received %s", t.String())
	}

	w := tabwriter.NewWriter(f.output, 15, 1, 3, ' ', 0)

	header := reflect.TypeOf(items).Elem()
	nf := header.NumField()

	// avoid printing elements with 'show:nolist'  attribute
	// special format tags
	avoid := make([]bool, nf)
	format := make([]string, nf)
	for i := 0; i < nf; i++ {
		avoid[i] = false
		showTags := strings.Split(header.Field(i).Tag.Get("show"), ",")
		for _, showTag := range showTags {
			if showTag == "nolist" {
				avoid[i] = true
			}
			if showTag == "minifySeconds" {
				format[i] = "minifySeconds"
			}
		}
	}

	// print header
	for i := 0; i < nf; i++ {
		if !avoid[i] {
			fmt.Fprintf(w, "%+v\t", header.Field(i).Tag.Get("header"))
		}
	}
	fmt.Fprintln(w)

	// print contents
	for i := 0; i < its.Len(); i++ {
		it := its.Index(i)
		nf := it.NumField()
		for i := 0; i < nf; i++ {
			if !avoid[i] {

				if format[i] == "minifySeconds" {

					remainingSeconds := int(it.Field(i).Float())
					s := remainingSeconds % 60
					remainingSeconds = (remainingSeconds - s)
					m := int(remainingSeconds/60) % 60
					remainingSeconds = (remainingSeconds - m*60)
					h := (remainingSeconds / 3600) % 24
					remainingSeconds = (remainingSeconds - h*3600)
					d := int(remainingSeconds / 86400)

					if d > 0 {
						fmt.Fprintf(w, "%dd%dh%dm\t", d, h, m)
					} else {
						fmt.Fprintf(w, "%dh%dm%ds\t", h, m, s)
					}

				} else {

					switch it.Field(i).Type().String() {
					case "json.RawMessage":
						fmt.Fprintf(w, "%s\t", it.Field(i).Interface())
					case "*json.RawMessage":
						if it.Field(i).IsNil() {
							fmt.Fprintf(w, " \t")
						} else {
							fmt.Fprintf(w, "%s\t", it.Field(i).Elem())
						}
					default:
						fmt.Fprintf(w, "%+v\t", it.Field(i).Interface())
					}
				}
			}
		}
		fmt.Fprintln(w)
	}
	w.Flush()

	return nil
}

// PrintError prints an error
func (f *TextFormatter) PrintError(context string, err error) {
	f.output.Write([]byte(fmt.Sprintf("ERROR: %s\n -> %s\n", context, err)))
}

// PrintFatal prints an error and exists
func (f *TextFormatter) PrintFatal(context string, err error) {
	f.PrintError(context, err)
	os.Exit(1)
}
