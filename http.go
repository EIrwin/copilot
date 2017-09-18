package copilot

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"errors"

	"strings"

	"github.com/eirwin/copilot/pkg/config"
	"github.com/eirwin/copilot/pkg/k8s"
)

func CopilotHandler(w http.ResponseWriter, r *http.Request) {

	// parse request text
	text, err := parseText(config.SlackToken(), w, r)
	if err != nil {
		log.Fatalln(err)
	}

	cmd, err := ParseCommand(text)
	if err != nil {
		log.Fatalln(err)
	}

	// initialize service dependencies
	client, err := k8s.NewClient(config.ConfigPath())
	if err != nil {
		log.Fatalln(err)
	}

	// initialize kubernetes service
	kubernetes, err := k8s.NewService(client)
	if err != nil {
		log.Fatalln(err)
	}

	// initialize copilot service
	service := NewService(kubernetes)
	output, err := service.Run(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	json, _ := json.Marshal(struct {
		Type string `json:"response_type"`
		Text string `json:"text"`
	}{
		Type: "in_channel",
		Text: fmt.Sprintf("```%s```", output),
	})

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {

}

func parseText(token string, w http.ResponseWriter, r *http.Request) (string, error) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return "", errors.New("method not allowed")
	}

	if token != r.FormValue("token") {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return "", errors.New("unauthorized")
	}

	return strings.Replace(r.FormValue("text"), "\r", "", -1), nil
}
