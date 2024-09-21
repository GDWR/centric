package swarm

import (
	"net/http"

	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/docker"
	"github.com/gdwr/centric/internal/http/response"
)

type Handler struct {
	databaseService *database.DatabaseService
	dockerService   *docker.DockerService
}

func NewHandler(db *database.DatabaseService, do *docker.DockerService) Handler {
	return Handler{databaseService: db, dockerService: do}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		response.BadRequest("path parameter \"id\" is required", w)
		return
	}

	environment, err := h.databaseService.GetEnvironmentByID(r.Context(), id)
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve environments", w)
		return
	}

	switch r.Pattern {
	case "GET /api/v1/environments/{id}/configs":
		h.getConfigs(environment, w, r)
	}
}
