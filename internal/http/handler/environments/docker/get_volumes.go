package docker

import (
	"net/http"

	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/http/response"
)

func (h Handler) getVolumes(environment database.Environment, w http.ResponseWriter, r *http.Request) {
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

	volumes, err := h.dockerService.GetVolumes(r.Context(), environment.Uri)
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve volumes", w)
		return
	}

	response.Json(volumes, w)
}
