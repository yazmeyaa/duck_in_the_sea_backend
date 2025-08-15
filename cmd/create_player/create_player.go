package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/yazmeyaa/duck_in_the_sea/internal/infrastructure/config"
	"github.com/yazmeyaa/duck_in_the_sea/internal/infrastructure/logger"
	"github.com/yazmeyaa/duck_in_the_sea/internal/player"
)

func main() {
	log := logger.New(os.Stdout)
	cfg := config.Load()
	conn, err := pgx.Connect(context.Background(), cfg.Postgres.DBString())
	if err != nil {
		log.Error("Error while create connection", slog.Any("error", err))
		panic(err)
	}
	r := player.NewRepository(conn)
	data := player.CreatePlayerDTO{
		UserID: 1,
	}
	p, err := r.Create(context.Background(), data)
	if err != nil {
		panic(err)
	}
	log.Info("User created", "id", p)
}
