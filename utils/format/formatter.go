package format

// Formatter defines output printing interface
type Formatter interface {
	PrintItem(item []string, header []string)
	PrintList(items [][]string, headers []string)
	PrintError(context string, err error)
	PrintFatal(context string, err error)
}
