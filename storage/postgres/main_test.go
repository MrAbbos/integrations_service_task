package postgres

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"integrations_service/config"
)

var (
	db *pgxpool.Pool
)

func TestMain(m *testing.M) {
	cfg := config.Load()
	conf, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		panic(err)
	}

	conf.MaxConns = cfg.PostgresMaxConnections

	db, err = pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
