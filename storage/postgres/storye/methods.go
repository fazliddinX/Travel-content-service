package storye

import (
	"Content-Service/genproto/auth_service"
	pb "Content-Service/genproto/content-service"
	"context"
)

func (s *StoryRepo) Create(in *pb.Story) (*pb.Story, error) {
	err := s.DB.QueryRow("insert into stories (title, content, location, author_id) values ($1, $2, $3, $4) returning id, created_at",
		in.Title, in.Content, in.Location, in.AuthorId).Scan(&in.Id, &in.CreatedAt)
	if err != nil {
		s.Logger.Error("Error inserting new story", "error", err)
		return nil, err
	}
	_, err = s.DB.Exec("insert into story_tags (story_id, tag) values ($1, $2)", in.Id, in.Tags)
	if err != nil {
		s.Logger.Error("Error inserting new story_tags", "error", err)
		return nil, err
	}
	return &pb.Story{}, nil
}

func (s *StoryRepo) Update(in *pb.PutStory) (*pb.Story, error) {
	story := pb.Story{}
	err := s.DB.QueryRow(`update stories set title = $1, content = $2 where id = $3
returning location, author_id, created_at, updated_at`,
		in.Title, in.Content, in.Id).Scan(&story.Location, &story.AuthorId, &story.CreatedAt, &story.UpdatedAt)
	if err != nil {
		s.Logger.Error("Error updating story", "error", err)
		return nil, err
	}
	err = s.DB.QueryRow("select tag from story_tags where story_id = $1", in.Id).Scan(&story.Tags)
	if err != nil {
		s.Logger.Error("Error updating story_tags", "error", err)
		return nil, err
	}
	story.Content = in.Content
	story.Title = in.Title
	story.Id = in.Id

	return &story, nil
}

func (s *StoryRepo) Delete(in *pb.StoryId) (*pb.MessageSuccess, error) {
	_, err := s.DB.Exec(`UPDATE stories SET 
		deleted_at = date_part('epoch', current_timestamp)::INT 
		WHERE id = $1 AND deleted_at = 0`, in.Id)
	if err != nil {
		s.Logger.Error("Error deleting story", "error", err)
		return nil, err
	}
	return &pb.MessageSuccess{Message: "successfully deleted"}, nil
}

func (s *StoryRepo) Get(in *pb.FilterStories) (*pb.Stories, error) {
	query := "select id, title, author_id, location, likes_count, comments_count, created_at from stories where deleted_at = 0"
	args := []interface{}{}
	if in.Title != "" {
		query += " and title like $1"
		args = append(args, "%"+in.Title+"%")
	}
	if in.Location != "" {
		query += " and location like $2"
		args = append(args, "%"+in.Location+"%")
	}
	query += " limit = $3, offset = $4"
	args = append(args, in.Limit, in.Offset)

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		s.Logger.Error("Error getting stories", "error", err)
		return nil, err
	}
	defer rows.Close()
	result := []*pb.GetStory{}
	for rows.Next() {
		a := pb.GetStory{}
		err := rows.Scan(&a.Id, &a.Title, &a.Author.Id, &a.Location, &a.LikeCount, &a.CommentCount, &a.CratedAt)
		if err != nil {
			s.Logger.Error("Error getting story", "error", err)
			return nil, err
		}
		result = append(result, &a)
	}
	return &pb.Stories{Stories: result}, err
}

func (s *StoryRepo) GetFullInfo(in *pb.LikeReq) (*pb.FullStory, error) {
	fullstory := &pb.FullStory{}
	err := s.DB.QueryRow(`select id, title, content, location, likes_count, comments_count, created_at, updated_at
from stories where deleted_at = 0 and id = $1`, in.StoryId).
		Scan(&fullstory.Id, &fullstory.Title, &fullstory.Content, &fullstory.Location,
			&fullstory.LikeCount, &fullstory.CommentCount, &fullstory.CreatedAt, &fullstory.UpdatedAt)
	if err != nil {
		s.Logger.Error("Error getting full story", "error", err)
		return nil, err
	}
	err = s.DB.QueryRow("select tag from story_tags where story_id = $1", in.StoryId).Scan(&fullstory.Tags)
	if err != nil {
		s.Logger.Error("Error getting full story_tags", "error", err)
		return nil, err
	}
	user, err := s.User.GetProfile(context.Background(), &auth_service.Id{Id: in.UserId})
	if err != nil {
		s.Logger.Error("Error getting user", "error", err)
		return nil, err
	}
	fullstory.Author.Id = user.Id
	fullstory.Author.FullName = user.FullName
	fullstory.Author.UserName = user.UserName
	return fullstory, nil
}

func (s *StoryRepo) AddComment(id string) error {
	var count int
	err := s.DB.QueryRow("select comments_count from stories where id = $1 and deleted_at = 0", id).Scan(&count)
	if err != nil {
		s.Logger.Error("Error adding comment", "error", err)
		return err
	}
	count++
	_, err = s.DB.Exec("update stories set comments_count = $1 where id = $1", count, id)
	if err != nil {
		s.Logger.Error("Error adding comment", "error", err)
	}
	return err
}

func (s *StoryRepo) RemoveComment(id string) error {
	var count int
	err := s.DB.QueryRow("select comments_count from stories where id = $1 and deleted_at = 0", id).Scan(&count)
	if err != nil {
		s.Logger.Error("Error adding comment", "error", err)
		return err
	}
	count--
	_, err = s.DB.Exec("update stories set comments_count = $1 where id = $1", count, id)
	if err != nil {
		s.Logger.Error("Error adding comment", "error", err)
	}
	return err
}

func (s *StoryRepo) AddLike(id string) error {
	var count int
	err := s.DB.QueryRow("select like_count from stories where id = $1 and deleted_at = 0", id).Scan(&count)
	if err != nil {
		s.Logger.Error("Error adding comment", "error", err)
		return err
	}
	count++
	_, err = s.DB.Exec("update stories set comments_count = $1 where id = $1", count, id)
	if err != nil {
		s.Logger.Error("Error adding comment", "error", err)
	}
	return err
}

func (s *StoryRepo) RemoveLike(id string) error {
	var count int
	err := s.DB.QueryRow("select like_count from stories where id = $1 and deleted_at = 0", id).Scan(&count)
	if err != nil {
		s.Logger.Error("Error adding comment", "error", err)
		return err
	}
	count--
	_, err = s.DB.Exec("update stories set comments_count = $1 where id = $1", count, id)
	if err != nil {
		s.Logger.Error("Error adding comment", "error", err)
	}
	return err
}
