package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
)
