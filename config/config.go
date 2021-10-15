package config

import (
	"os"
)

type Config struct {
	AppName         string
	AppPort         string
	LogLevel        string
	Environment     string
	EnvironmentLoc  string
	EnvironmentProd string
	EnvironmentRel  string
	JWTSecret       string
	RedisAddress    string
	DBUsername      string
	DBPassword      string
	DBHost          string
	DBPort          string
	DBName          string
	MinioEndpoint   string
	MinioAccessKey  string
	MinioSecretKey  string
	MinioRegion     string
	MinioBucket     string
}

func Init() *Config {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Warning("failed load .env")
	}

	appConfig := &Config{
		AppName:         os.Getenv("APP_NAME"),
		AppPort:         os.Getenv("APP_PORT"),
		LogLevel:        os.Getenv("LOG_LEVEL"),
		Environment:     os.Getenv("ENVIRONMENT"),
		EnvironmentLoc:  os.Getenv("ENVIRONMENT_LOC"),
		EnvironmentProd: os.Getenv("ENVIRONMENT_PROD"),
		EnvironmentRel:  os.Getenv("ENVIRONMENT_REL"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		RedisAddress:    os.Getenv("REDIS_ADDRESS"),
		DBUsername:      os.Getenv("DB_USERNAME"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBHost:          os.Getenv("DB_HOST"),
		DBPort:          os.Getenv("DB_PORT"),
		DBName:          os.Getenv("DB_NAME"),
		MinioEndpoint:   os.Getenv("MINIO_ENDPOINT"),
		MinioAccessKey:  os.Getenv("MINIO_ACCESS_KEY"),
		MinioSecretKey:  os.Getenv("MINIO_SECRET_KEY"),
		MinioRegion:     os.Getenv("MINIO_REGION"),
		MinioBucket:     os.Getenv("MINIO_BUCKET"),
	}

	return appConfig
}
