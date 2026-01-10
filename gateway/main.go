package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
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

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := s.Run(conf.Port, handler.Init(logger, float64(time.Second)/5, 10)); err != nil {
		slog.Error(fmt.Sprintf("Got an error while trying to start the server: %e\n", err))
	} else {
		slog.Info(fmt.Sprintf("Server is active on port :%d\n", conf.Port))
	}
}
