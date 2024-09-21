package response

import (
	"encoding/json"
	"net/http"
)

func Json(content any, w http.ResponseWriter) {
	bodyEncoded, err := json.Marshal(content)
	if err != nil {
		InternalServerError(err, "Unable to encode response", w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bodyEncoded)
}
