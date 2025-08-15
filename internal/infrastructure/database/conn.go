package database

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/yazmeyaa/duck_in_the_sea/internal/infrastructure/config"
)

func NewConnection(ctx context.Context, cfg *config.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, cfg.Postgres.DBString())

	return conn, err
}
