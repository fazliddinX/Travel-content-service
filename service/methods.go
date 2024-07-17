package service

import (
	pb "Content-Service/genproto/content-service"
	"context"
)

func (s *Service) CreateStory(ctx context.Context, in *pb.Story) (*pb.Story, error) {
	result, err := s.Story.Create(in)
	return result, err
}
func (s *Service) UpdateStory(ctx context.Context, in *pb.PutStory) (*pb.Story, error) {
	result, err := s.Story.Update(in)
	return result, err
}
func (s *Service) DeleteStory(ctx context.Context, in *pb.StoryId) (*pb.MessageSuccess, error) {
	success, err := s.Story.Delete(in)
	return success, err
}
func (s *Service) GetStories(ctx context.Context, in *pb.FilterStories) (*pb.Stories, error) {
	result, err := s.Story.Get(in)
	return result, err
}
func (s *Service) GetFullStoryInfo(ctx context.Context, in *pb.LikeReq) (*pb.FullStory, error) {
	result, err := s.Story.GetFullInfo(in)
	return result, err
}
func (s *Service) AddCommentStory(ctx context.Context, in *pb.StoryComment) (*pb.StoryCommentInfo, error) {
	result, err := s.CommentLike.AddCommentStory(in)
	return result, err
}
func (s *Service) GetComments(ctx context.Context, in *pb.FilterComment) (*pb.StoryComments, error) {
	result, err := s.CommentLike.GetCommentsStory(in)
	return result, err
}
func (s *Service) AddLikeToStory(ctx context.Context, in *pb.LikeReq) (*pb.LikeInfo, error) {
	result, err := s.CommentLike.ToStoryAddLike(in)
	return result, err
}
func (s *Service) CreateItineraries(ctx context.Context, in *pb.Itineraries) (*pb.InfoItineraries, error) {
	result, err := s.Itineraries.CreateItineraries(in)
	return result, err
}
func (s *Service) UpdateItineraries(ctx context.Context, in *pb.PutItineraries) (*pb.InfoItineraries, error) {
	result, err := s.Itineraries.UpdateItineraries(in)
	return result, err
}
func (s *Service) DeleteItineraries(ctx context.Context, in *pb.ItinerariesId) (*pb.MessageSuccess, error) {
	result, err := s.Itineraries.DeleteItineraries(in)
	return result, err
}
func (s *Service) GetItineraries(ctx context.Context, in *pb.FilterItineraries) (*pb.AllItineraries, error) {
	result, err := s.Itineraries.GetItineraries(in)
	return result, err
}
func (s *Service) AddLikeItineraries(ctx context.Context, in *pb.LikeItinerariesInfo) (*pb.LikeItinerariesInfo, error) {
	result, err := s.CommentLike.ToItinerariesAddLike(in)
	return result, err
}
func (s *Service) AddCommentItineraries(ctx context.Context, in *pb.ItinerariesComment) (*pb.ItinerariesCommentInfo, error) {
	result, err := s.CommentLike.AddCommentItineraries(in)
	return result, err
}

// -----------------------------------------------------------------------------------------------------------------------------
func (s *Service) GetDescriptions(ctx context.Context, in *pb.FilterDestinations) (*pb.Destinations, error) {
	result, err := s.Destination.GetDescriptions(in)
	return result, err
}
func (s *Service) GetDestinationById(ctx context.Context, in *pb.DestinationId) (*pb.FullDestinations, error) {
	result, err := s.Destination.GetDestinationById(in)
	return result, err
}
func (s *Service) CreateMessage(ctx context.Context, in *pb.MessageReq) (*pb.MessageRes, error) {
	result, err := s.Message.CreateMessage(in)
	return result, err
}
func (s *Service) GetMessages(ctx context.Context, in *pb.GetMessage) (*pb.Messages, error) {
	result, err := s.Message.GetMessages(in)
	return result, err
}
func (s *Service) CreateTravelTips(ctx context.Context, in *pb.TravelTipReq) (*pb.TravelTipRes, error) {
	result, err := s.Itineraries.CreateTravelTips(in)
	return result, err
}
func (s *Service) GetTravelTips(ctx context.Context, in *pb.FilterTravelTip) (*pb.TravelTips, error) {
	result, err := s.Itineraries.GetTravelTips(in)
	return result, err
}
func (s *Service) GetUserStatistics(ctx context.Context, in *pb.UserId) (*pb.UserStatistics, error) {
	result, err := s.Itineraries.GetUserStatistics(in)
	return result, err
}
func (s *Service) GetTrendingDestinations(ctx context.Context, in *pb.Void) (*pb.DestinationsRes, error) {
	result, err := s.Destination.GetTrendingDestinations(in)
	return result, err
}
