package messages

import (
	"database/sql"
	"log/slog"
)

type MessageRepo struct {
	DB     *sql.DB
	Logger *slog.Logger
}

func NewMessageRepo(db *sql.DB, logger *slog.Logger) *MessageRepo {
	return &MessageRepo{DB: db, Logger: logger}
}
