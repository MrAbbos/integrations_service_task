package postgres

import (
	"context"
	"integrations_service/genproto/integrations_service"
	"integrations_service/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type integrationsRepo struct {
	db *pgxpool.Pool
}

func NewIntegrationsRepo(db *pgxpool.Pool) storage.IntegrationsRepoI {
	return &integrationsRepo{
		db: db,
	}
}

func (b *integrationsRepo) CreateContent(ctx context.Context, req *integrations_service.ContentPullingRequest) (resp *integrations_service.ContentPullingRequest, err error) {
	return
}