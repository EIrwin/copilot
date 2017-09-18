package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/eirwin/copilot"
	"github.com/eirwin/copilot/pkg/config"
)

func init() {
	path := flag.String("path", "", "Absolute path to local configuration")

	flag.Parse()

	err := config.Init(*path)
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	http.HandleFunc("/", copilot.CopilotHandler)
	http.HandleFunc("/help", copilot.HelpHandler)

	log.Fatal(http.ListenAndServe(":"+config.Port(), nil))
}
