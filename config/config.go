package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port          string
	ServerAddress string
	Environment   string
	LogLevel      string
	JWTSecret     string
	DatabaseLink  string
	CacheLink     string
	EmailService  string
	EmailUser     string
	EmailPass     string
	EmailFrom     string
}

func (c Config) LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	serverAddress := os.Getenv("SERVER_IP")
	if serverAddress == "" {
		serverAddress = "127.0.0.1"
	}
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development"
	}
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET environment variable not set")
	}
	databaseLink := os.Getenv("DB_LINK")
	if databaseLink == "" {
		return nil, errors.New("DB_LINK environment variable not set")
	}
	cacheLink := os.Getenv("CACHE_LINK")
	if cacheLink == "" {
		return nil, errors.New("CACHE_LINK environment variable not set")
	}
	emailService := os.Getenv("EMAIL_SERVICE")
	if emailService == "" {
		return nil, errors.New("EMAIL_SERVICE environment variable not set")
	}
	emailUser := os.Getenv("EMAIL_USER")
	if emailUser == "" {
		return nil, errors.New("EMAIL_USER environment variable not set")

	}
	emailPass := os.Getenv("EMAIL_PASS")
	if emailPass == "" {
		return nil, errors.New("EMAIL_PASS environment variable not set")
	}
	emailFrom := os.Getenv("EMAIL_FROM")
	if emailFrom == "" {
		return nil, errors.New("EMAIL_FROM environment variable not set")
	}
	return &Config{
		Port:          port,
		ServerAddress: serverAddress,
		Environment:   environment,
		LogLevel:      logLevel,
		JWTSecret:     jwtSecret,
		DatabaseLink:  databaseLink,
		CacheLink:     cacheLink,
		EmailService:  emailService,
		EmailUser:     emailUser,
		EmailPass:     emailPass,
		EmailFrom:     emailFrom,
	}, nil
}
