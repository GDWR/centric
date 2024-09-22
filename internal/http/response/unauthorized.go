package response

import (
	"net/http"
)

func Unauthorized(message string, w http.ResponseWriter) {
	Error(message, http.StatusUnauthorized, w)
}
