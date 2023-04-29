package database

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvConnection() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

	return os.Getenv("MONGOURI")
}
