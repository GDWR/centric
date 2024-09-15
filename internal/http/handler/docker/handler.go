package docker

import (
	"net/http"

	"github.com/gdwr/centric/internal/docker"
)

type Handler struct {
	dockerService *docker.DockerService
}

func NewHandler(d *docker.DockerService) Handler {
	return Handler{dockerService: d}
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	containers, err := h.dockerService.GetContainers()

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, container := range containers {
		writer.Write([]byte(container.ID))
	}

	writer.WriteHeader(http.StatusOK)
}
