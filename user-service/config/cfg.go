package config

import (
	"os"
	"strconv"
)

type Config struct {
	DBName     string
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string

	Port int
}

func Init() *Config {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	return &Config{
		DBName:     os.Getenv("PG_NAME"),
		DBHost:     os.Getenv("PG_HOST"),
		DBUsername: os.Getenv("PG_USER"),
		DBPort:     os.Getenv("PG_PORT"),
		DBPassword: os.Getenv("PG_PASS"),

		Port: port,
	}
}
