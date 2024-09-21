package environments

import (
	"net/http"

	"github.com/gdwr/centric/internal/http/response"
)

func (h Handler) environmentsGetVolumes(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	if id == "" {
		response.BadRequest("path parameter \"id\" is required", writer)
		return
	}

	environment, err := h.databaseService.GetEnvironmentByID(request.Context(), id)
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve environments", writer)
		return
	}

	volumes, err := h.dockerService.GetVolumes(request.Context(), environment.Uri)
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve volumes", writer)
		return
	}

	response.Json(volumes, writer)
}
