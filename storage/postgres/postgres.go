package postgres

import (
	"context"
	"fmt"
	"integrations_service/config"
	"integrations_service/storage"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db   *pgxpool.Pool
	integrations storage.IntegrationsRepoI
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (l *Store) Log(ctx context.Context, msg string, data map[string]interface{}) {
	args := make([]interface{}, 0, len(data)+2) // making space for arguments + msg
	args = append(args, msg)
	for k, v := range data {
		args = append(args, fmt.Sprintf("%s=%v", k, v))
	}
	log.Println(args...)
}

func (s *Store) Integrations() storage.IntegrationsRepoI {
	if s.integrations == nil {
		s.integrations = NewIntegrationsRepo(s.db)
	}

	return s.integrations
}
