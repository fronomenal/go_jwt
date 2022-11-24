package inits

import (
	"log"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load an env file")
	}
}
