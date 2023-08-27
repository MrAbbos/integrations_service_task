package client

import (
	"integrations_service/config"
	"integrations_service/genproto/book_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	BookService() book_service.BookServiceClient
}

type grpcClients struct {
	bookService book_service.BookServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connBookService, err := grpc.Dial(
		cfg.ServiceHost+cfg.ServicePort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		bookService: book_service.NewBookServiceClient(connBookService),
	}, nil
}

func (g *grpcClients) BookService() book_service.BookServiceClient {
	return g.bookService
}
