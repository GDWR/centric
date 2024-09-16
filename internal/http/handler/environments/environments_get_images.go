package environments

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func (h Handler) environmentsGetImages(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("id")
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(`{"error": "id is required"}`))
		return
	}

	environment, err := h.databaseService.GetEnvironmentByID(request.Context(), id)
	if err != nil {
		log.Error().Err(err).Msg("failed to get environment")
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"error": "failed to get environment"}`))
		return
	}

	systemInfo, err := h.dockerService.GetImages(request.Context(), environment.Uri)
	if err != nil {
		log.Error().Err(err).Msg("failed to get images")
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(systemInfo); err != nil {
		log.Error().Err(err).Msg("failed to encode images")
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(`{"error": "failed to encode images"}`))
		return
	}
}
