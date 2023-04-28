package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetValue returns configuration value based on a given key from the .env file
func GetValue(key string) string {
	// load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file\n")
	}

	// return the value based on a given key
	return os.Getenv(key)
}
