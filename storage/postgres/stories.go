package postgres

import (
	"content-service/generated/communication"
	pb "content-service/generated/stories"
	"content-service/models"
	"database/sql"
)

type TravelStoriesRepo struct {
	DB *sql.DB
}

func NewTravelStoriesRepo(db *sql.DB) *TravelStoriesRepo {
	return &TravelStoriesRepo{
		DB: db,
	}
}

func (repo *TravelStoriesRepo) CreateTravelStory(req *pb.CreateTravelStoryRequest) (*pb.CreateTravelStoryResponse, error) {
	var storyResp pb.CreateTravelStoryResponse

	err := repo.DB.QueryRow(`
		INSERT INTO stories (title, content, location, author_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, title, content, location, author_id, created_at
	`, req.Title, req.Content, req.Location, req.AuthorId).
		Scan(&storyResp.Id, &storyResp.Title, &storyResp.Content, &storyResp.Location, &storyResp.AuthorId, &storyResp.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &storyResp, nil
}

func (repo *TravelStoriesRepo) UpdateTravelStory(req *pb.UpdateTravelStoryRequest) (*pb.UpdateTravelStoryResponse, error) {
	var storyResp pb.UpdateTravelStoryResponse

	err := repo.DB.QueryRow(`
		UPDATE stories SET title = $1, content = $2
		WHERE id = $3 AND deleted_at = 0
		RETURNING id, title, content, location, author_id, updated_at
	`, req.Title, req.Content, req.Id).
		Scan(&storyResp.Id, &storyResp.Title, &storyResp.Content, &storyResp.Location, &storyResp.AuthorId, &storyResp.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &storyResp, nil
}

func (repo *TravelStoriesRepo) DeleteTravelStory(id string) (*pb.DeleteTravelStoryResponse, error) {
	res, err := repo.DB.Exec(`
		UPDATE stories SET deleted_at = DATE_PART('epoch', CURRENT_TIMESTAMP)::INT
		WHERE deleted_at = 0 AND id = $1
	`, id)

	if err != nil {
		return &pb.DeleteTravelStoryResponse{
			Message: "Error in travel story deletion",
		}, nil
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return &pb.DeleteTravelStoryResponse{
		Message: "Story successfully deleted",
	}, nil
}

func (repo *TravelStoriesRepo) GetTravelStories(req *pb.ListTravelStoryRequest) (*pb.ListTravelStoryResponse, error) {
	var stories []*pb.TravelStory
	offset := (req.Page - 1) * req.Limit

	rows, err := repo.DB.Query(`
		SELECT id, title, author_id, location, created_at
		FROM stories
		WHERE deleted_at = 0
		ORDER BY created_at DESC
		OFFSET $1 LIMIT $2
	`, offset, req.Limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var story pb.TravelStory
		var author pb.Authors

		err := rows.Scan(&story.Id, &story.Title, &author.Id, &story.Location, &story.CreatedAt)
		if err != nil {
			return nil, err
		}

		story.Author = &author
		stories = append(stories, &story)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var total int32
	err = repo.DB.QueryRow(`
		SELECT COUNT(*) FROM stories WHERE deleted_at = 0
	`).Scan(&total)

	if err != nil {
		return nil, err
	}

	return &pb.ListTravelStoryResponse{
		Stories: stories,
		Total:   total,
		Limit:   req.Limit,
		Page:    req.Page,
	}, nil
}

func (repo *TravelStoriesRepo) GetTravelStory(id string) (*pb.GetTravelStoryResponse, error) {
	var storyResp pb.GetTravelStoryResponse
	var author pb.Author

	err := repo.DB.QueryRow(`
		SELECT id, title, content, location, author_id, created_at, updated_at
		FROM stories
		WHERE deleted_at = 0 AND id = $1
	`, id).
		Scan(&storyResp.Id, &storyResp.Title, &storyResp.Content, &storyResp.Location, &author.Id, &storyResp.CreatedAt, &storyResp.UpdatedAt)

	if err != nil {
		return nil, err
	}

	storyResp.Author = &author
	return &storyResp, err
}

func (repo *TravelStoriesRepo) AddComment(req *pb.AddCommentRequest) (*pb.AddCommentResponse, error) {
	var commentResp pb.AddCommentResponse

	err := repo.DB.QueryRow(`
		INSERT INTO comments (content, author_id, story_id)
		VALUES ($1, $2, $3)
		RETURNING id, content, author_id, story_id, created_at
	`, req.Content, req.AuthorId, req.StoryId).
		Scan(&commentResp.Id, &commentResp.Content, &commentResp.AuthorId, &commentResp.StoryId, &commentResp.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &commentResp, nil
}

func (repo *TravelStoriesRepo) GetComments(req *pb.ListCommentsRequest) (*pb.ListCommentsResponse, error) {
	var comments []*pb.Comment
	offset := (req.Page - 1) * req.Limit

	rows, err := repo.DB.Query(`
		SELECT c.id, c.content, c.author_id, c.created_at
		FROM comments c
		INNER JOIN stories s ON c.story_id = s.id
		WHERE s.deleted_at = 0
		ORDER BY c.created_at DESC
		OFFSET $1 LIMIT $2
	`, offset, req.Limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment pb.Comment
		var author pb.Authors

		err := rows.Scan(&comment.Id, &comment.Content, &author.Id, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comment.Author = &author
		comments = append(comments, &comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var total int32
	err = repo.DB.QueryRow(`
		SELECT COUNT(*)
		FROM comments c
		INNER JOIN stories s ON c.story_id = s.id
		WHERE s.deleted_at = 0
	`).Scan(&total)

	if err != nil {
		return nil, err
	}

	return &pb.ListCommentsResponse{
		Comments: comments,
		Total:    total,
		Limit:    req.Limit,
		Page:     req.Page,
	}, nil
}

func (repo *TravelStoriesRepo) AddLike(req *pb.AddLikeRequest) (*pb.AddLikeResponse, error) {
	var likeResp pb.AddLikeResponse

	err := repo.DB.QueryRow(`
		INSERT INTO likes (user_id, story_id)
		VALUES ($1, $2)
		RETURNING user_id, story_id, created_at
	`, req.UserId, req.StoryId).
		Scan(&likeResp.UserId, &likeResp.StoryId, &likeResp.LikedAt)

	if err != nil {
		return nil, err
	}

	return &likeResp, nil
}

func (repo *TravelStoriesRepo) CreateStoryTags(req models.StoryTag) error {
	_, err := repo.DB.Exec(`
		INSERT INTO story_tags (story_id, tag)
		VALUES ($1, $2)
	`, req.StoryId, req.Tag)

	return err
}

func (repo *TravelStoriesRepo) CountStories(id string) (int32, error) {
	var totalStories int32

	err := repo.DB.QueryRow(`
		SELECT COUNT(*)
		FROM stories
		WHERE deleted_at = 0 AND (author_id = $1 OR id = $1)
	`, id).Scan(&totalStories)

	if err != nil {
		return -1, err
	}

	return totalStories, nil
}

func (repo *TravelStoriesRepo) CountComments(id string) (int32, error) {
	var totalComments int32

	err := repo.DB.QueryRow(`
		SELECT COUNT(*)
		FROM comments c
		INNER JOIN stories s ON c.story_id = s.id
		WHERE s.deleted_at = 0 AND (c.author_id = $1 OR c.story_id = $1)
	`, id).Scan(&totalComments)

	if err != nil {
		return -1, err
	}

	return totalComments, nil
}

func (repo *TravelStoriesRepo) MostPopularStory(id string) (*communication.MostPopularStory, error) {
	var popularStory communication.MostPopularStory

	err := repo.DB.QueryRow(`
		SELECT id, title, likes_count
		FROM stories
		WHERE author_id = $1 AND deleted_at = 0
		ORDER BY likes_count DESC
		LIMIT 1
	`, id).
		Scan(&popularStory.Id, &popularStory.Title, &popularStory.LikesCount)

	return &popularStory, err
}
