package client

import (
	"integrations_service/config"
	"integrations_service/genproto/integrations_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	IntegrationsService() integrations_service.IntegrationsServiceClient
}

type grpcClients struct {
	integrationsService integrations_service.IntegrationsServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connIntegrationsService, err := grpc.Dial(
		cfg.ServiceHost+cfg.ServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		integrationsService: integrations_service.NewIntegrationsServiceClient(connIntegrationsService),
	}, nil
}

func (g *grpcClients) IntegrationsService() integrations_service.IntegrationsServiceClient {
	return g.integrationsService
}
