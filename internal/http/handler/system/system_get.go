package system

import (
	"encoding/json"
	"net/http"
)

type System struct {
	SetupComplete bool `json:"initialSetupComplete"`
}

func (h Handler) systemGet(writer http.ResponseWriter, request *http.Request) {
	initialSetupComplete, err := h.databaseService.InitialSetupComplete(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	response := System{
		SetupComplete: initialSetupComplete,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}
