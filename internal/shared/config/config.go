package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	FatalLevel = zapcore.FatalLevel
)

type Config struct {
	PORT 		 string
	DatabaseUrl    string
	DBMAXOpenConns int
	DBMAXIdleConns int
	DBConnMAXLife int
	AppEnv         string
	LogLevel       Level
	RedisHost	  string
	RedisPort	  int
	RedisPassword  string
	RedisDB        int
	RedisDbUrl	 string
	WebhookSecret   string
}

func requireEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("%s is required", key)
	}
	return v
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v, ok := os.LookupEnv(key); ok {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}

func levelFromEnv() Level {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		return DebugLevel
	case "warn", "warning":
		return WarnLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	default:
		return InfoLevel
	}
}

func SetupConfig() *Config {
	_=godotenv.Load(".env")
	return &Config{
		DatabaseUrl:      requireEnv("DATABASE_URL"),
		DBMAXOpenConns:   getEnvInt("DB_MAX_OPEN_CONNS", 25),
		DBMAXIdleConns:   getEnvInt("DB_MAX_IDLE_CONNS", 25),
		DBConnMAXLife:    getEnvInt("DB_CONN_MAX_LIFE", 300),
		AppEnv:           requireEnv("APP_ENV"),
		LogLevel:         levelFromEnv(),
		RedisHost:        getEnv("REDIS_HOST", "localhost"),
		RedisPort:        getEnvInt("REDIS_PORT", 6379),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		RedisDB:          getEnvInt("REDIS_DB", 0),
		RedisDbUrl:       requireEnv("REDIS_DB_URL"),
		PORT:              getEnv("PORT", "3333"),
		WebhookSecret:    requireEnv("WEBHOOK_SECRET"),
	}
}