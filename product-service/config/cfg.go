package config

import (
	"os"
)

type Config struct {
	DBName     string
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string

	Port string
}

func Init() *Config {
	return &Config{
		DBName:     os.Getenv("PG_NAME"),
		DBHost:     os.Getenv("PG_HOST"),
		DBUsername: os.Getenv("PG_USER"),
		DBPort:     os.Getenv("PG_PORT"),
		DBPassword: os.Getenv("PG_PASS"),

		Port: os.Getenv("PORT"),
	}
}
