package copilot

import (
	"encoding/json"
	"net/http"

	"fmt"

	"errors"

	"strings"

	"github.com/eirwin/copilot/pkg/config"
	"github.com/eirwin/copilot/pkg/k8s"
)

type Server struct {
	kubernetes k8s.Kubernetes
}

func NewServer(kubernetes k8s.Kubernetes) Server {
	return Server{
		kubernetes: kubernetes,
	}
}

func (s Server) Handler(w http.ResponseWriter, r *http.Request) {

	var output string
	parser := CommandParser{}

	// parse request text
	text, err := parseText(config.SlackToken(), w, r)
	if err != nil {
		output = parser.HelpWitMessage(err.Error())
		respond(output, w, r)
		return
	}

	// parse command from text
	cmd, err := parser.Parse(text)
	if err != nil {
		output = parser.Help()
		respond(output, w, r)
		return
	}

	// initialize copilot service
	service := NewService(s.kubernetes)
	output, err = service.Run(cmd)
	if err != nil {
		output = parser.Help()
		respond(output, w, r)
	}

	respond(output, w, r)
}

func respond(output string, w http.ResponseWriter, r *http.Request) {
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
