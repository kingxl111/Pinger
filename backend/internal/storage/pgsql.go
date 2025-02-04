package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	containersTable = "containers"
	pingsTable      = "pings"

	containersTableIPColumn     = "ip"
	containersTableNameColumn   = "name"
	containersTableActiveColumn = "active"

	pingsTableContainerIDColumn = "container_id"
	pingsTablePingTime          = "ping_time"
	pingsTableLastSuccessPing   = "last_success_ping"
)

type DB struct {
	pool *pgxpool.Pool
}

func NewDB(ctx context.Context, username, password, host, port, dbname, sslmode string) (*DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, dbname, sslmode)
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}
	return &DB{pool: pool}, nil
}

func (db *DB) Close() {
	db.pool.Close()
}
