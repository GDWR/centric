package response

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func InternalServerError(err error, message string, w http.ResponseWriter) {
	log.Error().
		Err(err).
		Msg(message)

	Error(message, http.StatusInternalServerError, w)
}
