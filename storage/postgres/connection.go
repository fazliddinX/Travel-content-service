package postgres

import (
	"Content-Service/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Connection() (*sql.DB, error) {
	cfg := config.Load()

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME)

	db, err := sql.Open("postgres", conn)
	return db, err
}
