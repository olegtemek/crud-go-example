package main

import (
	"context"

	"github.com/olegtemek/crud-go/internal/config"
	"github.com/olegtemek/crud-go/internal/handler/rest"
	"github.com/olegtemek/crud-go/internal/logger"
	"github.com/olegtemek/crud-go/internal/repository"
	"github.com/olegtemek/crud-go/internal/service"
	db "github.com/olegtemek/crud-go/pgk"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()

	if err != nil {
		panic("Cannot get config")
	}

	log := logger.NewLogger(cfg)

	db, err := db.NewPostgresConnect(log, cfg, ctx)

	if err != nil {
		panic("Cannot connect tot db")
	}

	repositories := repository.NewRepository(log, db)
	services := service.NewService(log, repositories)
	restHandler := rest.NewHandler(log, cfg, services)

	// тут будет gracefull stutdown
	log.Info("Trying to start server", restHandler.Init().ListenAndServe())
}
