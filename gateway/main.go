package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/kyeeego/urfu-microservices/gateway/config"
	delivery "github.com/kyeeego/urfu-microservices/gateway/delivery/http"
	"github.com/kyeeego/urfu-microservices/gateway/delivery/http/clients"
	"github.com/kyeeego/urfu-microservices/gateway/server"
)

func main() {

	conf := config.Init()

	redis := redis.NewClient(&redis.Options{
		Addr:     conf.RedisUrl,
		Password: "",
		DB:       0,
	})
	defer redis.Close()

	httpClient := &http.Client{}
	client := clients.NewHttpClientWithRetry(httpClient, 5, time.Second*3)
	handler := delivery.New(conf, client, redis)

	s := &server.Server{}

	if err := s.Run(conf.Port, handler.Init()); err != nil {
		log.Fatalf("Got an error while trying to start the server: %e\n", err)
	} else {
		log.Printf("Server is active on port :%d\n", conf.Port)
	}
}
