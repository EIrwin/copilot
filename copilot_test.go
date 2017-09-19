package copilot

import (
	"testing"

	"log"

	"github.com/eirwin/copilot/pkg/k8s"
)

func TestCopilot(t *testing.T) {
	path := "/Users/eirwin/.kube/config"
	client, err := k8s.NewClientFromConfig(&path)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	kubernetesService, err := k8s.NewService(client)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	service := NewService(kubernetesService)

	cmd, _ := ParseCommand("")

	result, _ := service.Run(cmd)
	log.Println("\n" + result)

}
