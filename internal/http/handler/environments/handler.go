package environments

import (
	"net/http"

	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/docker"
)

type Handler struct {
	databaseService *database.DatabaseService
	dockerService   *docker.DockerService
}

func NewHandler(db *database.DatabaseService, do *docker.DockerService) Handler {
	return Handler{databaseService: db, dockerService: do}
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Pattern {
	case "GET /api/v1/environments":
		h.environmentsGet(writer, request)
	case "GET /api/v1/environments/{id}":
		h.environmentsGetInfo(writer, request)
	}
}
