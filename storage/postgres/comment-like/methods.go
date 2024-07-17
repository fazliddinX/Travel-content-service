package comment_like

import (
	pb "Content-Service/genproto/content-service"
)

func (cl *CommentLikeRepo) AddCommentStory(in *pb.StoryComment) (*pb.StoryCommentInfo, error) {
	info := &pb.StoryCommentInfo{}

	err := cl.DB.QueryRow("insert into comments (content, author_id, story_id) values ($1, $2, $3)"+
		"returning id, created_at",
		in.Content, in.UserId, in.StoryId).Scan(&info.Id, info.CreatedAt)
	info.Content = in.Content
	info.AuthorId = in.UserId
	info.StoryId = in.StoryId

	if err != nil {
		cl.Logger.Error("error in add comment in comment", "error", err)
		return nil, err
	}
	err = cl.Story.AddComment(in.StoryId)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (cl *CommentLikeRepo) GetCommentsStory(in *pb.FilterComment) (*pb.StoryComments, error) {
	query := "select id, content, author_id, story_id, created_at from comments where deleted_at = 0"
	args := []interface{}{}
	if in.UserName != "" {
		query = query + " where author_id = $1"
		args = append(args, in.UserName)
	}
	query += " limit $2, offset $1"
	args = append(args, in.Limit, in.Offset)
	rows, err := cl.DB.Query(query, args...)
	if err != nil {
		cl.Logger.Error("error in get comments", "error", err)
		return nil, err
	}
	defer rows.Close()
	var comments []*pb.StoryCommentInfo
	for rows.Next() {
		var comment pb.StoryCommentInfo
		err = rows.Scan(&comment.Id, comment.Content, &comment.AuthorId, &comment.StoryId, &comment.CreatedAt)
		if err != nil {
			cl.Logger.Error("error in get comments", "error", err)
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return &pb.StoryComments{Comments: comments}, nil
}

func (cl *CommentLikeRepo) ToStoryAddLike(in *pb.LikeReq) (*pb.LikeInfo, error) {
	info := &pb.LikeInfo{}
	err := cl.DB.QueryRow("insert into story_likes (user_id, story_id) values ($1, $2) returning created_at", in.UserId, in.StoryId).
		Scan(&info.LikedAt)
	if err != nil {
		cl.Logger.Error("error in add like in comment", "error", err)
		return nil, err
	}
	err = cl.Story.AddLike(in.StoryId)
	if err != nil {
		return nil, err
	}
	info.StoryId = in.StoryId
	info.UserId = in.UserId
	return info, nil
}

func (cl *CommentLikeRepo) ToItinerariesAddLike(in *pb.LikeItinerariesInfo) (*pb.LikeItinerariesInfo, error) {
	info := &pb.LikeItinerariesInfo{}
	err := cl.DB.QueryRow("insert into itineraries_likes (user_id, itineraries_id) values ($1, $2) returning created_at", in.UserId, in.ItinerariesId).
		Scan(&info.LikedAt)
	if err != nil {
		cl.Logger.Error("error in add like in comment", "error", err)
		return nil, err
	}
	err = cl.Itineraries.AddLike(in.ItinerariesId)
	if err != nil {
		return nil, err
	}
	info.ItinerariesId = in.ItinerariesId
	info.UserId = in.UserId
	return info, nil
}

func (cl *CommentLikeRepo) AddCommentItineraries(in *pb.ItinerariesComment) (*pb.ItinerariesCommentInfo, error) {
	info := &pb.ItinerariesCommentInfo{}
	err := cl.DB.QueryRow("insert into comments(content, author_id, story_id) values ($1, $2, $3) returning id, created_at",
		in.Content, in.UserId, in.ItineraryId).Scan(&info.Id, info.CreatedAt)
	if err != nil {
		cl.Logger.Error("error in add comment itineraries in comment", "error", err)
		return nil, err
	}

	err = cl.Itineraries.AddComment(in.ItineraryId)
	if err != nil {
		return nil, err
	}
	info.AuthorId = in.UserId
	info.Content = in.Content
	info.ItineraryId = in.ItineraryId

	return info, nil
}
