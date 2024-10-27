package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

type Config struct {
	Database  string
	SecretKey string
}

func GetEnv() *Config {
	return &Config{
		Database:  os.Getenv("DATABASE"),
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}
