package config

import (
	"os"
	"strconv"
)

type Config struct {
	ProductsUrl string
	OrdersUrl   string
	UsersUrl    string

	Port int
}

func Init() *Config {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	return &Config{
		ProductsUrl: os.Getenv("PRODUCTS_URL"),
		OrdersUrl:   os.Getenv("ORDERS_URL"),
		UsersUrl:    os.Getenv("USERS_URL"),

		Port: port,
	}
}
