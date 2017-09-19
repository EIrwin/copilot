package copilot

type Command struct {
	namespace string
	resource  string
	operation string
	flags     map[string]string
	output    string
}

const (
	defaultOutput = "columns"
)

func ParseCommand(cmd string) (Command, error) {

	namepace := "staging"
	resouce := "service"
	operation := "get"
	output := defaultOutput

	return Command{
		namespace: namepace,
		resource:  resouce,
		operation: operation,
		output:    output,
	}, nil
}
