package comment_like

import (
	"Content-Service/storage/postgres/Itineraries"
	"Content-Service/storage/postgres/storye"
	"database/sql"
	"log/slog"
)

type CommentLikeRepo struct {
	Itineraries *Itineraries.ItinerariesRepo
	Story       *storye.StoryRepo
	Logger      *slog.Logger
	DB          *sql.DB
}

func NewCommentLikeRepo(logger *slog.Logger, db *sql.DB, story1 *storye.StoryRepo, itin *Itineraries.ItinerariesRepo) *CommentLikeRepo {
	return &CommentLikeRepo{Logger: logger, DB: db, Story: story1, Itineraries: itin}
}
