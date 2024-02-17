package db

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/olegtemek/crud-go/internal/config"
)

func NewPostgresConnect(_ *slog.Logger, cfg *config.Config, ctx context.Context) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(ctx, cfg.DatabaseUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
