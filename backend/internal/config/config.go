package config

import (
    "os"
    "strconv"
    "github.com/joho/godotenv"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    Port       string
    JWTSecret  string
    RateLimit  int
}

func LoadConfig() (*Config, error) {
    // Load .env file if exists (local development)
    godotenv.Load()

    rateLimit, _ := strconv.Atoi(getEnv("RATE_LIMIT", "100"))

    return &Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "3306"),
        DBUser:     getEnv("DB_USER", "root"),
        DBPassword: getEnv("DB_PASSWORD", "password"),
        DBName:     getEnv("DB_NAME", "userdb"),
        Port:       getEnv("PORT", "8080"),
        JWTSecret:  getEnv("JWT_SECRET", "your-secret-key"),
        RateLimit:  rateLimit,
    }, nil
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}