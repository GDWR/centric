package system

import (
	"net/http"

	"github.com/gdwr/centric/internal/http/response"
)

type System struct {
	SetupComplete bool `json:"initialSetupComplete"`
}

func (h Handler) systemGet(writer http.ResponseWriter, request *http.Request) {
	initialSetupComplete, err := h.databaseService.InitialSetupComplete(request.Context())
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve system information", writer)
		return
	}

	response.Json(System{SetupComplete: initialSetupComplete}, writer)
}
