package main

import (
	content_service "Content-Service/genproto/content-service"
	"Content-Service/logger"
	"Content-Service/service"
	"Content-Service/storage/postgres"
	"Content-Service/storage/postgres/Itineraries"
	comment_like "Content-Service/storage/postgres/comment-like"
	"Content-Service/storage/postgres/destinations"
	"Content-Service/storage/postgres/messages"
	"Content-Service/storage/postgres/storye"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func main() {
	logger := logger.InitLogger()
	logger.Info("main started")

	db, err := postgres.Connection()
	if err != nil {
		logger.Error("error connecting to database", "error", err)
		log.Fatalln(err)
	}
	defer db.Close()

	lister, err := net.Listen("tcp", ":5052")
	if err != nil {
		logger.Error("error listening on tcp", "error", err)
		log.Fatalln(err)
	}
	defer lister.Close()

	client, err := grpc.NewClient(":5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("error creating grpc client", "error", err)
		log.Fatalln(err)
	}
	st := storye.NewStoryRepo(logger, db, client)

	des := destinations.NewDestinationRepo(db, logger)

	iti := Itineraries.NewItinerariesRepo(db, logger, des, client)

	cl := comment_like.NewCommentLikeRepo(logger, db, st, iti)

	msg := messages.NewMessageRepo(db, logger)

	service1 := service.NewService(cl, des, iti, msg, st)
	grpcServer := grpc.NewServer()

	content_service.RegisterContentServer(grpcServer, service1)

	logger.Info("service running on port :5052")
	log.Fatal(grpcServer.Serve(lister))
}
