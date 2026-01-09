package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/kyeeego/urfu-microservices/order-service/config"
	"github.com/kyeeego/urfu-microservices/order-service/delivery/http"
	"github.com/kyeeego/urfu-microservices/order-service/repository"
	"github.com/kyeeego/urfu-microservices/order-service/server"
	"github.com/kyeeego/urfu-microservices/order-service/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	conf := config.Init()

	dsn := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(conf.DBUsername, conf.DBPassword),
		Host:     fmt.Sprintf("%s:%s", conf.DBHost, conf.DBPort),
		Path:     conf.DBName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	db, err := gorm.Open(postgres.Open(dsn.String()))
	if err != nil {
		log.Fatalf("Could not connect to database: %e\n", err)
	}

	repos := repository.New(db)
	services := service.New(repos)
	handler := http.New(services)

	s := &server.Server{}

	if err := s.Run(conf.Port, handler.Init()); err != nil {
		log.Fatalf("Got an error while trying to start the server: %e\n", err)
	} else {
		log.Printf("Server is active on port :%d\n", conf.Port)
	}
}
