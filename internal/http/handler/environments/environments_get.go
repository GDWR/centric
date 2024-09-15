package environments

import (
	"encoding/json"
	"net/http"
)

func (h Handler) environmentsGet(writer http.ResponseWriter, request *http.Request) {
	environments, err := h.databaseService.GetEnvironments(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if environments == nil {
		writer.WriteHeader(http.StatusNoContent)
		return
	}

	data, err := json.Marshal(environments)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(data)
}
