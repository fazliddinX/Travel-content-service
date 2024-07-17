package Itineraries

import (
	"Content-Service/genproto/auth_service"
	"Content-Service/storage/postgres/destinations"
	"database/sql"
	"google.golang.org/grpc"
	"log/slog"
)

type ItinerariesRepo struct {
	User        auth_service.AuthServiceClient
	Destination *destinations.DestinationRepo
	Logger      *slog.Logger
	DB          *sql.DB
}

func NewItinerariesRepo(db *sql.DB, logger *slog.Logger, d *destinations.DestinationRepo, conn *grpc.ClientConn) *ItinerariesRepo {
	user := auth_service.NewAuthServiceClient(conn)
	return &ItinerariesRepo{DB: db, Logger: logger, Destination: d, User: user}
}
