package storage

import (
	"context"
	service "integrations_service/genproto/integrations_service"
)

type StorageI interface {
	CloseDB()
	Integrations() IntegrationsRepoI
}

type IntegrationsRepoI interface {
	CreateContent(ctx context.Context, req *service.ContentPullingRequest) (resp *service.ContentPullingRequest, err error)
}
