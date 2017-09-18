package resource

const (
	defaultJSONIndent = "   "
	defaultJSONPrefix = ""
)

type JSONFormatter interface {
	JSON() (string, error)
}

type ColumnFormatter interface {
	Columns() (string, error)
}

type Result interface {
	JSONFormatter
	ColumnFormatter
	Data() [][]string
	Headers() []string
}
