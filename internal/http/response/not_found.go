package response

import (
	"net/http"
)

func NotFound(message string, w http.ResponseWriter) {
	Error(message, http.StatusNotFound, w)
}
