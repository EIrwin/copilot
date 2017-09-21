package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/eirwin/copilot"
	"github.com/eirwin/copilot/pkg/config"
	"github.com/eirwin/copilot/pkg/k8s"
)

var (
	server     copilot.Server
	kubernetes k8s.Kubernetes
)

func init() {
	path := flag.String("path", "", "Absolute path to local configuration")

	flag.Parse()

	err := config.Init(*path)
	if err != nil {
		log.Panicln(err)
	}

	// initialize service dependencies
	client, err := k8s.NewClient(config.ConfigPath())
	if err != nil {
		log.Fatalln(err)
	}

	// initialize kubernetes service
	kubernetes, err = k8s.NewService(client)
	if err != nil {
		log.Fatalln(err)
	}

	// initialize copilot command server
	server = copilot.NewServer(kubernetes)
}

func main() {

	http.HandleFunc("/", server.Handler)

	log.Println("Listening on localhost:" + config.Port())

	log.Fatal(http.ListenAndServe(":"+config.Port(), nil))
}
