package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	CASSANDRA_URL  = os.Getenv("CASSANDRA_URL")
)
