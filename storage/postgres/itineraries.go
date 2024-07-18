package postgres

import (
	"content-service/generated/communication"
	pb "content-service/generated/itineraries"
	"content-service/models"
	"database/sql"
	"errors"
)

type ItinerariesRepo struct {
	DB *sql.DB
}

func NewItinerariesRepo(db *sql.DB) *ItinerariesRepo {
	return &ItinerariesRepo{
		DB: db,
	}
}
func (repo *ItinerariesRepo) CreateItinerary(req *pb.CreateItineraryRequest) (*pb.CreateItineraryResponse, error) {
	var response pb.CreateItineraryResponse
	err := repo.DB.QueryRow(
		`INSERT INTO itineraries (title, description, start_date, end_date, author_id)
         VALUES ($1, $2, $3, $4, $5)
         RETURNING id, title, description, start_date, end_date, author_id, created_at`,
		req.Title, req.Description, req.StartDate, req.EndDate, req.AthorId,
	).Scan(&response.Id, &response.Title, &response.Description, &response.StartDate, &response.EndDate, &response.AuthorId, &response.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (repo *ItinerariesRepo) UpdateItinerary(req *pb.UpdateItineraryRequest) (*pb.UpdateItineraryResponse, error) {
	var response pb.UpdateItineraryResponse
	err := repo.DB.QueryRow(
		`UPDATE itineraries SET title = $1, description = $2
         WHERE id = $3 AND deleted_at = 0
         RETURNING id, title, description, start_date, end_date, author_id, updated_at`,
		req.Title, req.Description, req.Id,
	).Scan(&response.Id, &response.Title, &response.Description, &response.StartDate, &response.EndDate, &response.AuthorId, &response.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (repo *ItinerariesRepo) DeleteItinerary(itineraryId string) (*pb.DeleteItineraryResponse, error) {
	result, err := repo.DB.Exec(
		`DELETE FROM itineraries WHERE id = $1`, itineraryId,
	)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return &pb.DeleteItineraryResponse{
			Message: "Itinerary not found or already deleted",
		}, errors.New("itinerary not found or already deleted")
	}
	return &pb.DeleteItineraryResponse{
		Message: "Itinerary deleted successfully",
	}, nil
}

func (repo *ItinerariesRepo) GetItinerary(itineraryId string) (*pb.GetItineraryResponse, error) {
	var response pb.GetItineraryResponse
	var author pb.Author
	err := repo.DB.QueryRow(
		`SELECT id, title, description, start_date, end_date, author_id, created_at, updated_at
         FROM itineraries
         WHERE deleted_at = 0 AND id = $1`, itineraryId,
	).Scan(&response.Id, &response.Title, &response.Description, &response.StartDate, &response.EndDate, &author.Id, &response.CreatedAt, &response.UpdatedAt)
	if err != nil {
		return nil, err
	}
	response.Author = &author
	return &response, nil
}

func (repo *ItinerariesRepo) ListItineraries(req *pb.ListItinerariesRequest) (*pb.ListItinerariesResponse, error) {
	var itineraries []*pb.Itinerary
	offset := (req.Page - 1) * req.Limit
	rows, err := repo.DB.Query(
		`SELECT id, title, author_id, start_date, end_date, created_at
         FROM itineraries
         WHERE deleted_at = 0
         OFFSET $1 LIMIT $2`, offset, req.Limit,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var itinerary pb.Itinerary
		var author pb.Authors
		err = rows.Scan(&itinerary.Id, &itinerary.Title, &author.Id, &itinerary.StartDate, &itinerary.EndDate, &itinerary.CreatedAt)
		if err != nil {
			return nil, err
		}
		itinerary.Author = &author
		itineraries = append(itineraries, &itinerary)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	var total int32
	err = repo.DB.QueryRow(
		`SELECT COUNT(*) FROM itineraries WHERE deleted_at = 0`,
	).Scan(&total)
	if err != nil {
		return nil, err
	}
	return &pb.ListItinerariesResponse{
		Itineraries: itineraries,
		Total:       total,
		Limit:       req.Limit,
		Page:        req.Page,
	}, nil
}

func (repo *ItinerariesRepo) CreateItineraryDestinations(req models.ItineraryDestination) error {
	_, err := repo.DB.Exec(
		`INSERT INTO itinerary_destinations (itinerary_id, name, start_date, end_date)
         VALUES ($1, $2, $3, $4)`,
		req.ItineraryId, req.Name, req.StartDate, req.EndDate,
	)
	return err
}

func (repo *ItinerariesRepo) CreateItineraryActivity(req models.ItineraryActivity) error {
	_, err := repo.DB.Exec(
		`INSERT INTO itinerary_activities (destination_id, activity)
         VALUES ($1, $2)`,
		req.DestinationId, req.Activity,
	)
	return err
}

func (repo *ItinerariesRepo) GetItineraryDestinations(itineraryId string) ([]models.Result, error) {
	var destinations []models.Result
	rows, err := repo.DB.Query(
		`SELECT id, name, start_date, end_date
         FROM itinerary_destinations i_d
         JOIN itineraries i ON i.id = i_d.itinerary_id
         WHERE i.deleted_at = 0 AND i_d.itinerary_id = $1`, itineraryId,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var result models.Result
		err = rows.Scan(&result.ID, &result.Name, &result.StartDate, &result.EndDate)
		if err != nil {
			return nil, err
		}
		destinations = append(destinations, result)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return destinations, nil
}

func (repo *ItinerariesRepo) GetItineraryActivity(destinationId string) ([]string, error) {
	var activities []string
	rows, err := repo.DB.Query(
		`SELECT activity
         FROM itinerary_activities
         WHERE destination_id = $1`, destinationId,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var activity string
		err = rows.Scan(&activity)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return activities, nil
}

func (repo *ItinerariesRepo) CreateItineraryComments(req *pb.LeaveCommentRequest) (*pb.LeaveCommentResponse, error) {
	var response pb.LeaveCommentResponse
	err := repo.DB.QueryRow(
		`INSERT INTO itinerary_comments (author_id, itinerary_id, content)
         VALUES ($1, $2, $3)
         RETURNING id, content, author_id, itinerary_id, created_at`,
		req.AuthorId, req.ItineraryId, req.Content,
	).Scan(&response.Id, &response.Content, &response.AuthorId, &response.ItineraryId, &response.CreatedAt)
	return &response, err
}

func (repo *ItinerariesRepo) CountItinerary(itineraryId string) (int32, error) {
	var total int32
	err := repo.DB.QueryRow(
		`SELECT COUNT(*)
         FROM itineraries
         WHERE (deleted_at = 0) AND (author_id = $1 OR id = $1)`, itineraryId,
	).Scan(&total)
	if err != nil {
		return -1, err
	}
	return total, nil
}

func (repo *ItinerariesRepo) CountItineraryComments(itineraryId string) (int32, error) {
	var total int32
	err := repo.DB.QueryRow(
		`SELECT COUNT(*)
         FROM itineraries
         WHERE (deleted_at = 0) AND (author_id = $1 OR id = $1)`, itineraryId,
	).Scan(&total)
	if err != nil {
		return -1, err
	}
	return total, nil
}

func (repo *ItinerariesRepo) TotalCountriesVisited(authorId string) (int32, error) {
	var total int32
	err := repo.DB.QueryRow(
		`SELECT COUNT(DISTINCT id.id)
         FROM itinerary_destinations id
         JOIN itineraries i ON id.itinerary_id = i.id
         WHERE i.author_id = $1 AND i.deleted_at = 0`, authorId,
	).Scan(&total)
	return total, err
}

func (repo *ItinerariesRepo) MostPopularItinerary(authorId string) (*communication.MostPopularItinerary, error) {
	var resp communication.MostPopularItinerary
	err := repo.DB.QueryRow(
		`SELECT id, title, likes_count
         FROM itineraries
         WHERE author_id = $1 AND deleted_at`).Scan(
		&resp.Id, &resp.Title, &resp.LikesCount,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
