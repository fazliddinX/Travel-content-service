package Itineraries

import (
	pb "Content-Service/genproto/content-service"
)

func (i *ItinerariesRepo) CreateItineraries(in *pb.Itineraries) (*pb.InfoItineraries, error) {
	var info pb.InfoItineraries
	err := i.DB.QueryRow("insert into itineraries (title, description, start_date, end_date) values ($1, $2, $3, $4)"+
		"returning id, created_at", in.Title, in.Description, in.StartDate, in.EndDate).
		Scan(&info.Id, &info.CreatedAt)
	if err != nil {
		i.Logger.Error("error creating itineraries", "error", err)
		return nil, err
	}
	for k := 0; k < len(in.Descriptions); k++ {
		id, err := i.Destination.Create(in, k)
		if err != nil {
			return nil, err
		}
		_, err = i.DB.Exec("insert into itinerary_activities (destination_id, activity) values ($1, $2)", id, in.Descriptions[k].Activities)
		if err != nil {
			i.Logger.Error("error creating itinerary_activities", "error", err)
			return nil, err
		}
	}

	info.Title = in.Title
	info.Descriptions = in.Description
	info.StartDate = in.StartDate
	info.EndDate = in.EndDate
	info.AuthorId = in.AuthorId

	return &info, nil
}

func (i *ItinerariesRepo) UpdateItineraries(in *pb.PutItineraries) (*pb.InfoItineraries, error) {
	_, err := i.DB.Exec("update itineraries set title = $1, descriptions = $2 where id = $3", in.Title, in.Description, in.Id)
	if err != nil {
		i.Logger.Error("error updating itineraries", "error", err)
		return nil, err
	}
	info := pb.InfoItineraries{}
	err = i.DB.QueryRow("select id, title, description, start_data, end_data, author_id, created_at, updated_at from itineraries where id = $1", info.Id).
		Scan(&info.Id, info.Title, info.Descriptions, info.StartDate, info.EndDate, info.AuthorId, info.CreatedAt, info.UpdatedAt)
	if err != nil {
		i.Logger.Error("error updating itineraries", "error", err)
		return nil, err
	}
	return &info, nil
}

func (i *ItinerariesRepo) DeleteItineraries(in *pb.ItinerariesId) (*pb.MessageSuccess, error) {
	_, err := i.DB.Exec(`UPDATE itineraries SET 
		deleted_at = date_part('epoch', current_timestamp)::INT 
		WHERE id = $1 AND deleted_at = 0`, in.Id)
	if err != nil {
		i.Logger.Error("error deleting itineraries", "error", err)
	}
	return &pb.MessageSuccess{Message: "user deleted successfully"}, err
}

func (i *ItinerariesRepo) GetItineraries(in *pb.FilterItineraries) (*pb.AllItineraries, error) {
	query := "select id, titile, auther_id, start_date, end_date, like_count, comment_count, created_at from itineraries where id=$1"
	args := []interface{}{}
	if in.Title != "" {
		query += " and title = $1"
		args = append(args, in.Title)
	}
	if in.LikeCount != 0 {
		query += " and like_count = $2"
		args = append(args, in.LikeCount)
	}
	query += " limit $3 offset $4"
	args = append(args, in.Limit, in.Offset)

	rows, err := i.DB.Query(query, args...)
	if err != nil {
		i.Logger.Error("Error getting itineraries list", "query", query, "args", args, "error", err)
		return nil, err
	}
	defer rows.Close()

	its := []*pb.GItineraries{}
	for rows.Next() {
		it := pb.GItineraries{}
		err = rows.Scan(&it.Id, &it.Title, &it.Author, &it.StartDate, &it.EndDate, &it.LikeCount, &it.CommentCount, &it.CreatedAt)
		if err != nil {
			i.Logger.Error("Error getting itineraries list", "query", query, "args", args, "error", err)
			return nil, err
		}
		its = append(its, &it)
	}
	return &pb.AllItineraries{Itineraries: its}, nil
}

func (i *ItinerariesRepo) AddComment(id string) error {
	var count int
	err := i.DB.QueryRow("select comments_count from itineraries where id = $1 and deleted_at = 0", id).Scan(&count)
	if err != nil {
		i.Logger.Error("Error adding comment", "error", err)
		return err
	}
	count++
	_, err = i.DB.Exec("update itineraries set comments_count = $1 where id = $1", count, id)
	if err != nil {
		i.Logger.Error("Error adding comment", "error", err)
	}
	return err
}

func (i *ItinerariesRepo) RemoveComment(id string) error {
	var count int
	err := i.DB.QueryRow("select comments_count from itineraries where id = $1 and deleted_at = 0", id).Scan(&count)
	if err != nil {
		i.Logger.Error("Error adding comment", "error", err)
		return err
	}
	count--
	_, err = i.DB.Exec("update itineraries set comments_count = $1 where id = $1", count, id)
	if err != nil {
		i.Logger.Error("Error adding comment", "error", err)
	}
	return err
}

func (i *ItinerariesRepo) AddLike(id string) error {
	var count int
	err := i.DB.QueryRow("select like_count from itineraries where id = $1 and deleted_at = 0", id).Scan(&count)
	if err != nil {
		i.Logger.Error("Error adding comment", "error", err)
		return err
	}
	count++
	_, err = i.DB.Exec("update itineraries set like_count = $1 where id = $1", count, id)
	if err != nil {
		i.Logger.Error("Error adding comment", "error", err)
	}
	return err
}

func (i *ItinerariesRepo) RemoveLike(id string) error {
	var count int
	err := i.DB.QueryRow("select like_count from itineraries where id = $1 and deleted_at = 0", id).Scan(&count)
	if err != nil {
		i.Logger.Error("Error adding comment", "error", err)
		return err
	}
	count--
	_, err = i.DB.Exec("update itineraries set like_count = $1 where id = $1", count, id)
	if err != nil {
		i.Logger.Error("Error adding comment", "error", err)
	}
	return err
}

func (i *ItinerariesRepo) CreateTravelTips(in *pb.TravelTipReq) (*pb.TravelTipRes, error) {
	info := pb.TravelTipRes{}
	err := i.DB.QueryRow("insert into travel_tips (title, content, category) values ($1, $2, $3)",
		in.Title, in.Content, in.Category).Scan(&info.Id, info.CreatedAt)
	if err != nil {
		i.Logger.Error("Error creating travel tips", "error", err)
		return nil, err
	}
	info.Content = in.Content
	info.Category = in.Category
	info.Title = in.Title

	return &info, nil
}

func (i *ItinerariesRepo) GetTravelTips(in *pb.FilterTravelTip) (*pb.TravelTips, error) {
	query := "select id,  title, content, category, auther_id, created_id from travel_tips where deleted_at = 0"
	args := []interface{}{}
	if in.Category != "" {
		query += " and category = $1"
		args = append(args, in.Category)
	}
	query += " limit $2 offset $3"
	args = append(args, in.Limit, in.Offset)
	rows, err := i.DB.Query(query, args...)
	if err != nil {
		i.Logger.Error("Error getting travel tips", "query", query, "args", args, "error", err)
		return nil, err
	}
	defer rows.Close()

	travelTips := []*pb.TravelTipRes{}
	for rows.Next() {
		travelTip := pb.TravelTipRes{}
		err = rows.Scan(&travelTip.Id, &travelTip.Title, &travelTip.Content, &travelTip.Category, &travelTip.AuthorId, &travelTip.CreatedAt)
		if err != nil {
			i.Logger.Error("Error getting travel tips", "query", query, "args", args, "error", err)
			return nil, err
		}
		travelTips = append(travelTips, &travelTip)
	}
	return &pb.TravelTips{TravelTips: travelTips}, nil
}

func (i *ItinerariesRepo) GetUserStatistics(in *pb.UserId) (*pb.UserStatistics, error) {

	var (
		storyCount       int
		itinereriesCount int
		likesCount       int
		commentsCount    int
	)
	var popularStory pb.PopularStory
	var prpularItineries pb.PopularItinerary

	err := i.DB.QueryRow("SELECT COUNT(*) FROM stories where author_id = $1", in.UserId).Scan(&storyCount)
	if err != nil {
		i.Logger.Error("Error in scan itinereries", "error", err)
		return nil, err
	}
	err = i.DB.QueryRow("SELECT COUNT(*) FROM itineries where author_id = $1", in.UserId).Scan(&itinereriesCount)
	if err != nil {
		i.Logger.Error("Error in scan", "error", err)
		return nil, err
	}
	err = i.DB.QueryRow("SELECT COUNT(*) FROM likes where user_id = $1", in.UserId).Scan(&likesCount)
	if err != nil {
		i.Logger.Error("Error in scan itineraries", "error", err)
		return nil, err
	}
	err = i.DB.QueryRow("SELECT COUNT(*) FROM coments c join stories s on s.id = c.story_id where s.author_id = $1", in.UserId).Scan(&commentsCount)
	if err != nil {
		i.Logger.Error("Error in scan itineraries", "error", err)
		return nil, err
	}
	err = i.DB.QueryRow("select id, title, like_count from stories where author_id = $1 order by like_count desc limi 1",
		in.UserId).Scan(&popularStory.Id, &popularStory.Title, &popularStory.LikesCount)
	if err != nil {
		i.Logger.Error("Error in scan itineraries", "error", err)
		return nil, err
	}
	info := &pb.UserStatistics{
		UserId:                in.UserId,
		TotalStories:          int32(storyCount),
		TotalItineraries:      int32(itinereriesCount),
		TotalCountriesVisited: int32(storyCount),
		TotalLikesReceived:    int32(likesCount),
		TotalCommentsReceived: int32(commentsCount),
		MostPopularStory:      &popularStory,
		MostPopularItinerary:  &prpularItineries,
	}
	return info, nil
}
