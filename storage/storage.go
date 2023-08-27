package storage

import (
	"context"
	"integrations_service/genproto/book_service"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
}

type BookRepoI interface {
	Create(ctx context.Context, req *book_service.CreateBookRequest) (resp *book_service.BookPrimaryKey, err error)
	Get(ctx context.Context, req *book_service.BookPrimaryKey) (resp *book_service.Book, err error)
	GetList(ctx context.Context, req *book_service.GetBooksListRequest) (resp *book_service.GetBooksListResponse, err error)
	Update(ctx context.Context, req *book_service.UpdateBookRequest) (rowsAffected int64, err error)
	PatchUpdate(ctx context.Context, req *book_service.PatchUpdateRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, req *book_service.BookPrimaryKey) (rowsAffected int64, err error)
}
