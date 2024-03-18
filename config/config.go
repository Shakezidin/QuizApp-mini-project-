package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host          string
	User          string
	Password      string
	Database      string
	Port          string
	Sslmode       string
	REDISHOST     string
	REDISPassword string
	SERVERPORT    string
	SECRETKEY     string
}

func LoadConfig() (*Config, error) {
	godotenv.Load("../../.env")

	var config Config

	// Use os.Getenv to retrieve environment variables
	config.Host = os.Getenv("HOST")
	config.User = os.Getenv("USER")
	config.Password = os.Getenv("PASSWORD")
	config.Database = os.Getenv("DATABASE")
	config.Port = os.Getenv("PORT")
	config.Sslmode = os.Getenv("SSLMODE")
	config.REDISHOST = os.Getenv("REDISHOST")
	config.REDISPassword = os.Getenv("REDISPASSWORD")
	config.SECRETKEY = os.Getenv("SECRETKEY")

	return &config, nil
}
