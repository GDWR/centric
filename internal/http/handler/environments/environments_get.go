package environments

import (
	"net/http"

	"github.com/gdwr/centric/internal/http/response"
)

func (h Handler) environmentsGet(writer http.ResponseWriter, request *http.Request) {
	environments, err := h.databaseService.GetEnvironments(request.Context())
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve environments", writer)
		return
	}

	if environments == nil {
		response.NotFound("No environments found", writer)
		return
	}

	response.Json(environments, writer)
}
