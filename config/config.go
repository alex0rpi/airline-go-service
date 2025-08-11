package config

import (
	"os"
)

type Config struct {
    PostgresDSN string
    Port        string
}

func LoadConfig() *Config {
    return &Config{
        PostgresDSN: getEnv("POSTGRES_DSN", "host=localhost user=postgres password=postgres dbname=fleet port=5432 sslmode=disable"),
        Port:        getEnv("PORT", "8080"),
    }
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}