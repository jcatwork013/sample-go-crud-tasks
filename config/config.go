package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Could not load .env file, using system environment variables")
	}
}

// GetEnv retrieves an environment variable
func GetEnv(key string) string {
	LoadEnv()
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("❌ Missing environment variable: %s", key)
	}
	return value
}

// GetDatabaseURI constructs the database connection string from environment variables
func GetDatabaseURI() string {
	host := GetEnv("DB_HOST")
	port := GetEnv("DB_PORT")
	user := GetEnv("DB_USER")
	password := GetEnv("DB_PASSWORD")
	dbname := GetEnv("DB_NAME")

	// Construct DSN string
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}
