package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	apigen "github.com/yazmeyaa/duck_in_the_sea/internal/api/http_server/generated"
	"github.com/yazmeyaa/duck_in_the_sea/internal/api/http_server/handlers"
	"github.com/yazmeyaa/duck_in_the_sea/internal/api/http_server/middlewares"
	"github.com/yazmeyaa/duck_in_the_sea/internal/infrastructure/config"
	"github.com/yazmeyaa/duck_in_the_sea/internal/infrastructure/database"
	"github.com/yazmeyaa/duck_in_the_sea/internal/infrastructure/logger"
	"github.com/yazmeyaa/duck_in_the_sea/internal/users/repository"
	usecase "github.com/yazmeyaa/duck_in_the_sea/internal/users/use_case"
)

func main() {
	logger := logger.New(os.Stdout)

	logger.Info("Started process...", "PID", os.Getpid())

	cfg := config.Load()
	dbConn, err := database.NewConnection(context.Background(), cfg)
	defer func() {
		err := dbConn.Close(context.Background())
		if err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(err)
	}
	r := repository.NewPostgresRepository(dbConn)
	uc := usecase.NewUserUseCase(r, logger)

	var si apigen.ServerInterface = handlers.NewServer(uc, logger)
	h := apigen.HandlerWithOptions(si, apigen.StdHTTPServerOptions{
		Middlewares: []apigen.MiddlewareFunc{
			middlewares.SpawnWithAuthMiddleware(uc, cfg, logger),
		},
	})
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port), h); err != nil {
			panic(err)
		}
	}()
	logger.Info("Server started successfully", "host", cfg.Server.Host, "port", cfg.Server.Port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	s := <-c
	logger.Info("Got OS interrupt signal, shutting down...", "signal", s)
}
