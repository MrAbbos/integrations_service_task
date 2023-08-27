package service

import (
	"context"
	"integrations_service/config"
	"integrations_service/genproto/integrations_service"
	"integrations_service/grpc/client"
	"integrations_service/pkg/logger"
	"integrations_service/storage"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type integrationsService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	integrations_service.UnimplementedIntegrationsServiceServer
}

func NewIntegrationsService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *integrationsService {
	return &integrationsService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (b *integrationsService) ContentPulling(ctx context.Context, req *integrations_service.ContentPullingRequest) (resp *empty.Empty, err error) {
	b.log.Info("---CreateIntegrations--->", logger.Any("req", req))

	_, err = b.strg.Integrations().CreateContent(ctx, req)

	if err != nil {
		b.log.Error("!!!CreateIntegrations--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
