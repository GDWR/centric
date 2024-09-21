package swarm

import (
	"net/http"

	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/http/response"
)

func (h Handler) getConfigs(environment database.Environment, w http.ResponseWriter, r *http.Request) {
	configs, err := h.dockerService.GetConfigs(r.Context(), environment.Uri)
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve configs", w)
		return
	}

	response.Json(configs, w)
}
