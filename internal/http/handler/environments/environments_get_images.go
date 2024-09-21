package environments

import (
	"net/http"

	"github.com/gdwr/centric/internal/http/response"
)

func (h Handler) environmentsGetImages(writer http.ResponseWriter, request *http.Request) {
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

	images, err := h.dockerService.GetImages(request.Context(), environment.Uri)
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve images", writer)
		return
	}

	response.Json(images, writer)
}
