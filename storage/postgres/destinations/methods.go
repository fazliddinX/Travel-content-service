package destinations

import (
	pb "Content-Service/genproto/content-service"
	"fmt"
	"time"
)

func (d *DestinationRepo) Create(in *pb.Itineraries, index int) (string, error) {
	var id string
	err := d.db.QueryRow("insert into destinations (name) values ($1)", in.Descriptions[index].Name).Scan(&id)
	if err != nil {
		d.Logger.Error("error in insert destination name", "err", err)
		return "", err
	}
	return id, nil
}

func (d *DestinationRepo) GetDescriptions(in *pb.FilterDestinations) (*pb.Destinations, error) {
	query := "select id, name, country, description  from destinations where deleted_at = 0"
	args := []interface{}{}

	if in.Name != "" {
		query += ", and name = $1"
		args = append(args, in.Name)
	}
	if in.Country != "" {
		query += ", and country = $2"
		args = append(args, in.Country)
	}
	query += " limit $3 offset $4"
	args = append(args, in.Limit, in.Offset)

	rows, err := d.db.Query(query, args...)
	if err != nil {
		d.Logger.Error("error in get by filter", "err", err)
	}
	defer rows.Close()

	destinations := []*pb.Destination{}
	for rows.Next() {
		destination := pb.Destination{}
		err = rows.Scan(&destination.Id, &destination.Name, &destination.Country, &destination.Descriptions)
		if err != nil {
			d.Logger.Error("error in get by filter", "err", err)
			return nil, err
		}
		destinations = append(destinations, &destination)
	}
	return &pb.Destinations{Destinations: destinations}, nil
}

func (d *DestinationRepo) GetDestinationById(in *pb.DestinationId) (*pb.FullDestinations, error) {
	ds := pb.FullDestinations{}
	err := d.db.QueryRow(`select id, name, country, description, best_time_to_visit, average_cost_per_day, currency, language
		 from destinations where id = $1`, in.Id).
		Scan(&ds.Id, &ds.Name, &ds.Country, &ds.Description, &ds.BestTimeToVisite, &ds.AvarageCostPerDay, &ds.Currency, &ds.Language)
	if err != nil {
		d.Logger.Error("error in get destination by id", "err", err)
		return nil, err
	}
	return &ds, nil
}

func (d *DestinationRepo) GetTrendingDestinations(a *pb.Void) (*pb.DestinationsRes, error) {
	fmt.Println(a)
	month := time.Now().Month().String()
	rows, err := d.db.Query("select id, name, country from destinations where deleted_at = 0 and best_time_to_visit = $1 and language = $2", month, "english")
	if err != nil {
		d.Logger.Error("error in get trending destinations", "err", err)
		return nil, err
	}
	defer rows.Close()

	destinations := []*pb.DestinationRs{}
	for rows.Next() {
		destination := pb.DestinationRs{}
		err = rows.Scan(&destination.Id, &destination.Name, &destination.Country)
		if err != nil {
			d.Logger.Error("error in Scanning destinations", "err", err)
		}
		reason := fmt.Sprintf("this wares is so beautifull in month %s", month)
		destination.PopularityScore = int32(90)
		destination.TrendingReason = reason
		destinations = append(destinations, &destination)
	}
	return &pb.DestinationsRes{Destinations: destinations}, nil
}
