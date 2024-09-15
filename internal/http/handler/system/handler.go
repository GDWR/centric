package system

import (
	"net/http"

	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/user"
)

type Handler struct {
	databaseService *database.DatabaseService
	userService     *user.UserService
}

func NewHandler(d *database.DatabaseService, u *user.UserService) Handler {
	return Handler{
		databaseService: d,
		userService:     u,
	}
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Pattern {
	case "GET /api/v1/system":
		h.systemGet(writer, request)
	case "POST /api/v1/system/initial-setup":
		h.systemInitialSetup(writer, request)
	}
}
