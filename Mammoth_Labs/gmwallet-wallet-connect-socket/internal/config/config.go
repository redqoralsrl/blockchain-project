package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	JwtSecretKey string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err env load: %v", err)
	}

	return &Config{
		JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}
