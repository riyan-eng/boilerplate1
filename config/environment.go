package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	env := ".env"
	if os.Getenv("GO_ENV") != "" {
		env += "." + os.Getenv("GO_ENV")
	}
	err := godotenv.Load(env)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
