package grpc

import (
	"integrations_service/config"
	"integrations_service/genproto/integrations_service"
	"integrations_service/grpc/client"
	"integrations_service/grpc/service"
	"integrations_service/pkg/logger"
	"integrations_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	integrations_service.RegisterIntegrationsServiceServer(grpcServer, service.NewIntegrationsService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
