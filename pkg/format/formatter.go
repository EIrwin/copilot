package format

import (
	"strings"

	"github.com/ryanuber/columnize"
)

const (
	defaultDelimiter = "|"
)

type ColumnFormatter struct {
	config *columnize.Config
}

func NewColumnFormatter() ColumnFormatter {
	config := columnize.DefaultConfig()
	config.Delim = defaultDelimiter
	return ColumnFormatter{
		config: config,
	}
}

func (c ColumnFormatter) Format(headers []string, data [][]string) string {
	var lines []string

	// append headers
	lines = append(lines, strings.Join(headers, defaultDelimiter))

	// append data
	for _, d := range data {
		lines = append(lines, strings.Join(d, defaultDelimiter))
	}

	return columnize.Format(lines, c.config)
}
