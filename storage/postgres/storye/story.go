package storye

import (
	"Content-Service/genproto/auth_service"
	"database/sql"
	"google.golang.org/grpc"
	"log/slog"
)

type StoryRepo struct {
	User   auth_service.AuthServiceClient
	Logger *slog.Logger
	DB     *sql.DB
}

func NewStoryRepo(logger *slog.Logger, db *sql.DB, c *grpc.ClientConn) *StoryRepo {
	user := auth_service.NewAuthServiceClient(c)
	return &StoryRepo{Logger: logger, DB: db, User: user}
}
