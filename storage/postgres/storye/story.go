package storye

import (
	"database/sql"
	"log/slog"
)

type StoryRepo struct {
	Logger *slog.Logger
	DB     *sql.DB
}

func NewStoryRepo(logger *slog.Logger, db *sql.DB) *StoryRepo {
	return &StoryRepo{Logger: logger, DB: db}
}
