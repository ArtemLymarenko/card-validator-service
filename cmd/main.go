package main

import (
	"card-validator-service/internal/app"
	"card-validator-service/internal/config"
	v1Router "card-validator-service/internal/interface/rest/v1"
	v1Handlers "card-validator-service/internal/interface/rest/v1/handlers"
	"fmt"
	"net/http"
)

func main() {
	cfg := config.MustGet(config.EnvLocal)
	fmt.Println(cfg)

	handlers := v1Handlers.New(cfg)

	router := v1Router.GetGinRouter(handlers)

	server := &http.Server{
		Addr:         cfg.HttpServer.Addr,
		Handler:      router,
		ReadTimeout:  cfg.HttpServer.Timeout,
		WriteTimeout: cfg.HttpServer.Timeout,
		IdleTimeout:  cfg.HttpServer.IdleTimeout,
	}

	application := app.New(server)
	application.Start()
}
