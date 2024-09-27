package storage

import (
	"context"
	"errors"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var (
	ErrPostgresNoUrl = errors.New("No PSQL URL provided")
	ErrCHNoUrl       = errors.New("No ClickHouse URL provided")
)

func NewPostgres(ctx context.Context) (*pgxpool.Pool, error) {
	connURL := os.Getenv("PSQL_URL")
	if connURL == "" {
		return nil, ErrPostgresNoUrl
	}

	_cfg, err := pgxpool.ParseConfig(connURL)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	_cfg.MaxConns = 100

	conn, err := pgxpool.NewWithConfig(ctx, _cfg)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	err = conn.Ping(ctx)
	return conn, err
}

func RunSelect[T any](ctx context.Context, pool *pgxpool.Pool, sql string, args ...any) (t []T, err error) {
	var rows pgx.Rows
	rows, err = pool.Query(ctx, sql, args...)
	if err != nil {
		return
	}
	t, err = pgx.CollectRows(rows, pgx.RowToStructByName[T])
	rows.Close()
	return
}
