package docker

import (
	"net/http"

	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/http/response"
)

func (h Handler) getContainers(environment database.Environment, w http.ResponseWriter, r *http.Request) {
	containers, err := h.dockerService.GetContainers(r.Context(), environment.Uri)
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve containers", w)
		return
	}

	response.Json(containers, w)
}
