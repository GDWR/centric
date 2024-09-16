package http

import (
	"net/http"

	"github.com/gdwr/centric/internal/auth"
	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/docker"
	"github.com/gdwr/centric/internal/http/handler/environments"
	"github.com/gdwr/centric/internal/http/handler/system"
	"github.com/gdwr/centric/internal/user"
)

type Server struct {
	AuthService     *auth.AuthService
	DatabaseService *database.DatabaseService
	DockerService   *docker.DockerService
	UserService     *user.UserService
}

func (s *Server) Run() error {
	environmentsHandler := environments.NewHandler(s.DatabaseService, s.DockerService)
	http.Handle("GET /api/v1/environments", environmentsHandler)
	http.Handle("GET /api/v1/environments/{id}", environmentsHandler)
	http.Handle("GET /api/v1/environments/{id}/containers", environmentsHandler)
	http.Handle("GET /api/v1/environments/{id}/images", environmentsHandler)

	systemHandler := system.NewHandler(s.DatabaseService, s.UserService)
	http.Handle("GET /api/v1/system", systemHandler)
	http.Handle("POST /api/v1/system/initial-setup", systemHandler)

	http.Handle("/", http.FileServer(http.Dir("static")))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}
