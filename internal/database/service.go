package database

import (
	"database/sql"
	"embed"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var migrations embed.FS

type DatabaseService struct {
	*Queries

	db *sql.DB
}

func NewDatabaseService() (*DatabaseService, error) {
	db, err := sql.Open("sqlite3", "centric.sqlite")
	if err != nil {
		return nil, err
	}

	goose.SetLogger(LoggerAdapter{Logger: log.Logger})
	goose.SetBaseFS(migrations)
	if err := goose.SetDialect("sqlite"); err != nil {
		return nil, err
	}
	if err := goose.Up(db, "migrations"); err != nil {
		return nil, err
	}

	return &DatabaseService{
		Queries: New(db),
		db:      db,
	}, nil
}
