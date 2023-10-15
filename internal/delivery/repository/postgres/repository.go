package postgres

import (
	"database/sql"
)

func NewPostgresRepository(postgresDB *sql.DB) *Queries {
	return New(postgresDB)
}
