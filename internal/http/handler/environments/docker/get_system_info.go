package docker

import (
	"net/http"

	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/http/response"
)

func (h Handler) getSystemInformation(environment database.Environment, w http.ResponseWriter, r *http.Request) {
	systemInfo, err := h.dockerService.GetSystemInformation(r.Context(), environment.Uri)
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve system information", w)
		return
	}

	response.Json(systemInfo, w)
}
