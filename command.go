package copilot

import (
	"strings"

	"bytes"

	"fmt"

	"github.com/alexflint/go-arg"
)

const (
	appName       = "copilot"
	defaultOutput = "columns"
)

type NamespaceFlags struct {
	Namespace string `arg:"-n,help: Namespace to run against"`
}

type OutputFlags struct {
	Output string `arg:"-o,help: Output format. One of json|columns|yaml."`
}

type SelectorFlags struct {
	LabelSelector []string `arg:"-l,separate,help: Selector (label query) to filter on."`
	FieldSelector []string `arg:"-f,separate,help: Selector (label query) to filter on.""`
}

type Command struct {
	Operation string   `arg:"positional"`
	Resource  string   `arg:"positional"`
	Names     []string `arg:"positional"`
	NamespaceFlags
	SelectorFlags
	OutputFlags
}

type CommandParser struct {
}

func (c CommandParser) Parse(cmd string) (Command, error) {
	var command Command
	parser, err := initParser(&command)
	if err != nil {
		return command, err
	}

	if command.Output == "" {
		command.Output = defaultOutput
	}

	parser.Parse(strings.Split(cmd, " "))
	return command, nil
}

func (c CommandParser) Usage() string {

	command := defaultCommand()
	parser, err := initParser(&command)
	if err != nil {
		return ""
	}

	var usage bytes.Buffer
	parser.WriteUsage(&usage)
	return usage.String()
}

func (c CommandParser) Help() string {

	command := defaultCommand()
	parser, err := initParser(&command)
	if err != nil {
		return ""
	}

	var help bytes.Buffer
	parser.WriteHelp(&help)
	return help.String()
}

func (c CommandParser) UsageWithMessage(msg string) string {
	usage := c.Usage()
	return fmt.Sprintf("%v \n\n %v", msg, usage)
}

func (c CommandParser) HelpWitMessage(msg string) string {
	help := c.Help()
	return fmt.Sprintf("%v \n\n %v", msg, help)
}

func initParser(cmd *Command) (*arg.Parser, error) {
	return arg.NewParser(arg.Config{
		Program: appName,
	}, cmd)
}

func defaultCommand() Command {
	return Command{
		OutputFlags: OutputFlags{
			Output: defaultOutput,
		},
	}
}
