package environments

import (
	"net/http"

	"github.com/gdwr/centric/internal/http/response"
)

func (h Handler) environmentsGetConfigs(writer http.ResponseWriter, request *http.Request) {
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

	configs, err := h.dockerService.GetConfigs(request.Context(), environment.Uri)
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve configs", writer)
		return
	}

	response.Json(configs, writer)
}
