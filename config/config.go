package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string
	PostgresURI string
	RedisURI    string
	JWTSecret   string
}

// With this we will acess the variables
var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("unable to load the env file", fmt.Sprint("vars"), err)
		log.Println(".env file not found")
		return
	}

	AppConfig = &Config{
		ServerPort:  getEnv("PORT", "8080"),
		PostgresURI: getEnv("POSTGRES_URL", "postgres://healthuser:healthpassword@localhost:5432/healthcare?sslmode=disable"),
		RedisURI:    getEnv("REDIS_URL", "localhost:6379"),
		JWTSecret:   getEnv("JWT_SECRET", "thisismysecret123"),
	}
	// slog.Info("Configuration file loaded successfully", "sucess")
	log.Println("config file loaded sucessfully")

}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback

}
