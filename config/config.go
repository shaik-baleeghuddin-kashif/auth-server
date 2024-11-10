package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	MongoDB          string
	MongoDBUser      string
	MongoDBPassword  string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
}

var AppConfig *Config

func LoadConfig() {
	envPath := filepath.Join("config/", ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error Loading .env file from  %s: %v", envPath, err)
	}
	AppConfig = &Config{
		Port:             getEnv("PORT", "8000"),
		MongoDB:          getEnv("MONGODB", ""),
		MongoDBUser:      getEnv("MONGODBUSER", ""),
		MongoDBPassword:  getEnv("MONGODBPASSWORD", ""),
		PostgresDB:       getEnv("POSTGRESDB", ""),
		PostgresUser:     getEnv("POSTGRESDBUSER", ""),
		PostgresPassword: getEnv("POSTGRESDBPASSWORD", ""),
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
