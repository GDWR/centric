package response

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func Json(content any, w http.ResponseWriter) {
	bodyEncoded, err := json.Marshal(content)
	if err != nil {
		InternalServerError(err, "Unable to encode response", w)
		return
	}

	log.Debug().
		Int("status", http.StatusOK).
		Int("content-length", len(bodyEncoded)).
		Send()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bodyEncoded)
}
