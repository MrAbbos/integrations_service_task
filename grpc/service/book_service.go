package service

import (
	"context"
	"integrations_service/config"
	"integrations_service/genproto/book_service"
	"integrations_service/grpc/client"
	"integrations_service/pkg/logger"
	"integrations_service/storage"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type bookService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	book_service.UnimplementedBookServiceServer
}

func NewBookService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *bookService {
	return &bookService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (b *bookService) CreateBook(ctx context.Context, req *book_service.CreateBookRequest) (resp *book_service.Book, err error) {
	b.log.Info("---CreateBook--->", logger.Any("req", req))

	pKey, err := b.strg.Book().Create(ctx, req)

	if err != nil {
		b.log.Error("!!!CreateBook--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return b.strg.Book().Get(ctx, pKey)
}

func (b *bookService) GetBook(ctx context.Context, req *book_service.BookPrimaryKey) (resp *book_service.Book, err error) {
	b.log.Info("---GetBook--->", logger.Any("req", req))

	resp, err = b.strg.Book().Get(ctx, req)

	if err != nil {
		b.log.Error("!!!GetBook--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *bookService) GetBooksList(ctx context.Context, req *book_service.GetBooksListRequest) (resp *book_service.GetBooksListResponse, err error) {
	b.log.Info("---GetBooksList--->", logger.Any("req", req))

	resp, err = b.strg.Book().GetList(ctx, req)

	if err != nil {
		b.log.Error("!!!GetBooksList--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (b *bookService) UpdateBook(ctx context.Context, req *book_service.UpdateBookRequest) (resp *book_service.Book, err error) {
	b.log.Info("---UpdateBook--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Book().Update(ctx, req)

	if err != nil {
		b.log.Error("!!!UpdateBook--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Book().Get(ctx, &book_service.BookPrimaryKey{Id: req.Book.Id})
	if err != nil {
		b.log.Error("!!!UpdateBook--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *bookService) PatchUpdate(ctx context.Context, req *book_service.PatchUpdateRequest) (resp *book_service.Book, err error) {
	b.log.Info("---PatchUpdate--->", logger.Any("req", req))

	rowsAffected, err := b.strg.Book().PatchUpdate(ctx, req)

	if err != nil {
		b.log.Error("!!!PatchUpdate--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = b.strg.Book().Get(ctx, &book_service.BookPrimaryKey{Id: req.Id})
	if err != nil {
		b.log.Error("!!!Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (b *bookService) DeleteBook(ctx context.Context, req *book_service.BookPrimaryKey) (resp *empty.Empty, err error) {
	b.log.Info("---DeleteBook--->", logger.Any("req", req))

	resp = &empty.Empty{}

	rowsAffected, err := b.strg.Book().Delete(ctx, req)

	if err != nil {
		b.log.Error("!!!DeleteBook--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return resp, err
}
