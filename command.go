package copilot

type Command struct {
	namespace string
	resource  string
	action    string
	options   map[string]string
}

func ParseCommand(cmd string) (Command, error) {

	//TODO: validate cmd format

	//TODO: parse resource

	//TODO: parse action

	//TODO: parse options

	return Command{}, nil
}
