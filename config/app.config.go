package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Default port
var PORT = ":8000"

func BootApp() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if portEnv := os.Getenv("PORT"); portEnv != "" {
		PORT = portEnv
	}

	bootDatabase()
	connectDatabase()
	runMigration()
}
