package response

import (
	"net/http"
)

func InternalServerError(err error, message string, w http.ResponseWriter) {
	Error(message, http.StatusInternalServerError, w)
}
