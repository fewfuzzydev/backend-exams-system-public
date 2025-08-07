package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPort    string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	JWTSecret string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env vars")
	}

	AppConfig = &Config{
		DBPort:    getEnv("DB_PORT", ""),
		DBHost:    getEnv("DB_HOST", ""),
		DBUser:    getEnv("DB_USER", ""),
		DBPass:    getEnv("DB_PASSWORD", ""),
		DBName:    getEnv("DB_NAME", ""),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
