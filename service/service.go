package service

import (
	pb "Content-Service/genproto/content-service"
	"Content-Service/storage/postgres/Itineraries"
	"Content-Service/storage/postgres/comment-like"
	"Content-Service/storage/postgres/destinations"
	"Content-Service/storage/postgres/messages"
	"Content-Service/storage/postgres/storye"
)

type Service struct {
	pb.UnimplementedContentServer
	CommentLike *comment_like.CommentLikeRepo
	Destination *destinations.DestinationRepo
	Itineraries *Itineraries.ItinerariesRepo
	Message     *messages.MessageRepo
	Story       *storye.StoryRepo
}

func NewService(cl *comment_like.CommentLikeRepo, des *destinations.DestinationRepo,
	it *Itineraries.ItinerariesRepo, msg *messages.MessageRepo, repo *storye.StoryRepo) *Service {
	return &Service{CommentLike: cl, Destination: des, Itineraries: it, Message: msg, Story: repo}
}
