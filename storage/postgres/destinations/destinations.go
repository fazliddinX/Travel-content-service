package destinations

import (
	"database/sql"
	"log/slog"
)

type DestinationRepo struct {
	Logger *slog.Logger
	db     *sql.DB
}

func NewDestinationRepo(db *sql.DB, logger *slog.Logger) *DestinationRepo {
	return &DestinationRepo{Logger: logger, db: db}
}
