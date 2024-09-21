package environments

import (
	"net/http"

	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/docker"
	"github.com/gdwr/centric/internal/http/handler/environments/swarm"

	dockerHandler "github.com/gdwr/centric/internal/http/handler/environments/docker"
)

type Handler struct {
	databaseService *database.DatabaseService
	dockerService   *docker.DockerService

	swarmHandler  swarm.Handler
	dockerHandler dockerHandler.Handler
}

func NewHandler(db *database.DatabaseService, do *docker.DockerService) Handler {
	return Handler{
		databaseService: db,
		dockerService:   do,
		swarmHandler:    swarm.NewHandler(db, do),
		dockerHandler:   dockerHandler.NewHandler(db, do),
	}
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	switch request.Pattern {
	case "GET /api/v1/environments":
		h.getEnvironments(writer, request)
	case "GET /api/v1/environments/{id}/containers",
		"GET /api/v1/environments/{id}/images",
		"GET /api/v1/environments/{id}/networks",
		"GET /api/v1/environments/{id}/system-information",
		"GET /api/v1/environments/{id}/volumes":
		h.dockerHandler.ServeHTTP(writer, request)
	case "GET /api/v1/environments/{id}/configs":
		h.swarmHandler.ServeHTTP(writer, request)
	}
}
