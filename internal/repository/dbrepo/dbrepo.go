package dbrepo

import (
	"database/sql"

	"github.com/timfewi/bookingsGo/internal/config"
	"github.com/timfewi/bookingsGo/internal/repository"
)

type postgresDbRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewPostgresRepo creates a new repository
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDbRepo{
		App: a,
		DB:  conn,
	}
}
