package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"log"
	"os"
)

type Config struct {
	DB_HOST string
	DB_PORT int
	DB_USER string
	DB_NAME string
	DB_PASS string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	cfg := Config{}
	cfg.DB_HOST = cast.ToString(Coalesce("DB_HOST", "localhost"))
	cfg.DB_PORT = cast.ToInt(Coalesce("DB_PORT", "5432"))
	cfg.DB_USER = cast.ToString(Coalesce("DB_USER", "postgres"))
	cfg.DB_NAME = cast.ToString(Coalesce("DB_NAME", "travel_content"))
	cfg.DB_PASS = cast.ToString(Coalesce("DB_PASS", "123321"))

	return cfg
}

func Coalesce(key, defaultValue string) interface{} {
	valuer, exists := os.LookupEnv(key)
	if exists {
		return valuer
	}
	return defaultValue
}
