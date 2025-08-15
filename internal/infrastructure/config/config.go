package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Postgres PostgresConfig
	Server   ServerConfig
	Telegram TelegramConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	SSLMode  string
}

type TelegramConfig struct {
	BotToken string
}

func (p *PostgresConfig) DBString() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		p.User, p.Password, p.Host, p.Port, p.Database, p.SSLMode,
	)
}

func mustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("missing required env var: %s", key))
	}
	return val
}

func getEnvOrDefault(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func getEnvOrDefaultAsInt(key string, def int) int {
	v := os.Getenv(key)
	i, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return i
}

func mustGetEnvAsInt(key string) int {
	val := mustGetEnv(key)
	i, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Errorf("env %s: %w", key, err))
	}
	return i
}

func Load() *Config {
	return &Config{
		Postgres: PostgresConfig{
			User:     mustGetEnv("POSTGRES_USER"),
			Password: mustGetEnv("POSTGRES_PASSWORD"),
			Host:     mustGetEnv("POSTGRES_HOST"),
			Port:     mustGetEnvAsInt("POSTGRES_PORT"),
			Database: mustGetEnv("POSTGRES_DATABASE"),
			SSLMode:  getEnvOrDefault("POSTGRES_SSL", "disable"),
		},
		Server: ServerConfig{
			Host: getEnvOrDefault("HOST", ""),
			Port: getEnvOrDefaultAsInt("PORT", 8080),
		},
		Telegram: TelegramConfig{
			BotToken: mustGetEnv("TELEGRAM_BOT_TOKEN"),
		},
	}
}
