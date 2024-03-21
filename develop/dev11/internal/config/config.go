package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Addr     string
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

func NewConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	addr := os.Getenv("SERVER_ADDR")

	return Config{
		Addr:     addr,
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DB:       db,
	}, nil
}
