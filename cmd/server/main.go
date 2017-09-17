package main

import (
	"log"
	"net/http"

	"github.com/eirwin/kubes"
)

func main() {
	//TODO: parse config

	const url = ""
	http.HandleFunc("/", kubesHandler)

	log.Fatal(http.ListenAndServe(url, nil))
}

func kubesHandler(w http.ResponseWriter, r *http.Request) {
	text := ""
	cmd, err := kubes.ParseCommand(text)
	if err != nil {
		log.Fatalln(err)
	}

	config := kubes.Config{}
	service, err := kubes.NewService(config)
	if err != nil {
		log.Fatalln(err)
	}

	output, err := service.Run(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(output)
}
