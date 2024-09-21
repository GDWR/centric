package response

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type errorResponse struct {
	Reason string `json:"reason"`
}

func Error(message string, statusCode int, w http.ResponseWriter) {
	bodyEncoded, err := json.Marshal(errorResponse{Reason: message})
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to encode error")

		return
	}

	log.Debug().
		Str("reason", message).
		Int("status", statusCode).
		Int("content-length", len(message)).
		Send()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(bodyEncoded)
}
