package config

import (
	"os"
	"strconv"
)

type Config struct {
	ProductsUrl string
	OrdersUrl   string
	UsersUrl    string
	RedisUrl    string
	RedisTtl    int

	Port int
}

func Init() *Config {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	ttl, _ := strconv.Atoi(os.Getenv("REDIS_TTL"))

	return &Config{
		ProductsUrl: os.Getenv("PRODUCTS_URL"),
		OrdersUrl:   os.Getenv("ORDERS_URL"),
		UsersUrl:    os.Getenv("USERS_URL"),
		RedisUrl:    os.Getenv("REDIS_URL"),
		RedisTtl:    ttl,

		Port: port,
	}
}
