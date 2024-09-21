package http

import (
	"net/http"

	"github.com/gdwr/centric/internal/auth"
	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/docker"
	"github.com/gdwr/centric/internal/http/handler/environments"
	"github.com/gdwr/centric/internal/http/handler/system"
	"github.com/gdwr/centric/internal/http/middleware"
	"github.com/gdwr/centric/internal/user"
)

type Server struct {
	AuthService     *auth.AuthService
	DatabaseService *database.DatabaseService
	DockerService   *docker.DockerService
	UserService     *user.UserService
}

func (s *Server) Run() error {
	mux := http.NewServeMux()
	middleware := middleware.LoggerMiddleware()

	environmentsHandler := middleware(environments.NewHandler(s.DatabaseService, s.DockerService))
	mux.Handle("GET /api/v1/environments", environmentsHandler)
	mux.Handle("GET /api/v1/environments/{id}", environmentsHandler)
	mux.Handle("GET /api/v1/environments/{id}/configs", environmentsHandler)
	mux.Handle("GET /api/v1/environments/{id}/containers", environmentsHandler)
	mux.Handle("GET /api/v1/environments/{id}/images", environmentsHandler)
	mux.Handle("GET /api/v1/environments/{id}/networks", environmentsHandler)
	mux.Handle("GET /api/v1/environments/{id}/system-information", environmentsHandler)
	mux.Handle("GET /api/v1/environments/{id}/volumes", environmentsHandler)

	systemHandler := middleware(system.NewHandler(s.DatabaseService, s.UserService))
	mux.Handle("GET /api/v1/system", systemHandler)
	mux.Handle("POST /api/v1/system/initial-setup", systemHandler)

	mux.Handle("/", http.FileServer(http.Dir("static")))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		return err
	}

	return nil
}
