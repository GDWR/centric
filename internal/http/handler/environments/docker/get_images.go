package docker

import (
	"net/http"

	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/http/response"
)

func (h Handler) getImages(environment database.Environment, w http.ResponseWriter, r *http.Request) {
	images, err := h.dockerService.GetImages(r.Context(), environment.Uri)
	if err != nil {
		response.InternalServerError(err, "Unable to retrieve images", w)
		return
	}

	response.Json(images, w)
}
