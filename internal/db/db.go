package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Database struct {
	URL     string
	MaxCon  int
	MinCon  int
	ConnTTL time.Duration
	Pool    *pgxpool.Pool
}

func New() *Database {
	return &Database{}
}

func (db *Database) Initialize() {
	ctx, cancel := context.WithTimeout(context.Background(), db.ConnTTL)
	defer cancel()
	pool, err := pgxpool.New(ctx, db.URL)
	if err != nil {
		// Todo describe the error
	}
	if err := pool.Ping(ctx); err != nil {
		// Todo describe the error
	}
	db.Pool = pool
}

func (db *Database) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}
