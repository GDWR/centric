package auth

import (
	"net/http"

	"github.com/gdwr/centric/internal/auth"
	"github.com/gdwr/centric/internal/database"
)

type Handler struct {
	authService     *auth.AuthService
	databaseService *database.DatabaseService
}

func NewHandler(a *auth.AuthService, d *database.DatabaseService) Handler {
	return Handler{authService: a, databaseService: d}
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Pattern {
	case "POST /api/v1/auth/login":
		h.login(writer, request)
	}
}
