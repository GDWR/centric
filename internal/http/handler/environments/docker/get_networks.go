package docker

import (
	"net/http"

	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/http/response"
)

func (h Handler) getNetworks(environment database.Environment, w http.ResponseWriter, r *http.Request) {
	networks, err := h.dockerService.GetNetworks(r.Context(), environment.Uri)
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve networks", w)
		return
	}

	response.Json(networks, w)
}
