package response

import (
	"net/http"
)

func BadRequest(message string, w http.ResponseWriter) {
	Error(message, http.StatusBadRequest, w)
}
