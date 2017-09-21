package copilot

import (
	"log"
	"testing"
)

func TestCommandParser(t *testing.T) {
	input := "-n staging get pods"
	parser := CommandParser{}
	cmd, err := parser.Parse(input)
	if err != nil {
		t.Fail()
	}

	log.Println(cmd)
}
