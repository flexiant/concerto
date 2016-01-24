package format

// Formatter defines output printing interface
type Formatter interface {
	PrintItem(item interface{}, header []string)
	PrintList(items [][]string, headers []string)
	PrintError(context string, err error)
}
