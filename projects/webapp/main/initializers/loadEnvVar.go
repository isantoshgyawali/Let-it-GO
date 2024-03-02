package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVar() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading the .env file")
	}
}
