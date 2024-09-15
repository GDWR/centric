package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gdwr/centric/internal/auth"
	"github.com/gdwr/centric/internal/database"
	"github.com/gdwr/centric/internal/docker"
	"github.com/gdwr/centric/internal/http"
	"github.com/gdwr/centric/internal/user"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")
	json := flag.Bool("json", false, "sets log format to json")

	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	if !*json {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}

func main() {
	authService := auth.NewAuthService()
	databaseService, err := database.NewDatabaseService()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to create database service")
	}
	dockerService, err := docker.NewDockerService()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to create docker service")
	}

	server := http.Server{
		AuthService:     authService,
		DatabaseService: databaseService,
		DockerService:   dockerService,
		UserService:     user.NewUserService(authService, databaseService),
	}

	log.Info().
		Str("address", "localhost:8080").
		Msg("Starting HTTP server")

	if err := server.Run(); err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start HTTP server")
	}
}
