package postgres

import (
	"context"
	"database/sql"
	"integrations_service/config"
	"integrations_service/genproto/book_service"
	"integrations_service/storage"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type bookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) storage.BookRepoI {
	return &bookRepo{
		db: db,
	}
}

type Book struct {
	Id            string
	Name          string
	NumberOfPages int32
	CreatedAt     sql.NullString
	UpdatedAt     sql.NullString
}

func (b *bookRepo) Create(ctx context.Context, req *book_service.CreateBookRequest) (resp *book_service.BookPrimaryKey, err error) {
	query := `insert into books 
				(id, 
				name, 
				number_of_pages
				) VALUES (
					$1, 
					$2, 
					$3
				)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return resp, err
	}

	_, err = b.db.Exec(ctx, query,
		uuid.String(),
		req.Name,
		req.NumberOfPages)

	if err != nil {
		return resp, err
	}

	resp = &book_service.BookPrimaryKey{
		Id: uuid.String(),
	}

	return
}

func (b *bookRepo) Get(ctx context.Context, req *book_service.BookPrimaryKey) (resp *book_service.Book, err error) {
	result := &Book{}
	resp = &book_service.Book{}

	query := `select 
		id, 
		name, 
		number_of_pages,
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	from books 
	where id = $1`

	err = b.db.QueryRow(ctx, query, req.Id).Scan(
		&result.Id,
		&result.Name,
		&result.NumberOfPages,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		return resp, err
	}

	if result.CreatedAt.Valid {
		resp.CreatedAt = result.CreatedAt.String
	}

	if result.UpdatedAt.Valid {
		resp.UpdatedAt = result.UpdatedAt.String
	}

	resp.Id = result.Id
	resp.Name = result.Name
	resp.NumberOfPages = result.NumberOfPages

	return
}

func (b *bookRepo) GetList(ctx context.Context, req *book_service.GetBooksListRequest) (resp *book_service.GetBooksListResponse, err error) {
	resp = &book_service.GetBooksListResponse{}
	var (
		params      (map[string]interface{})
		filter      string
		order       string
		arrangement string
		offset      string
		limit       string
		q           string
	)

	params = map[string]interface{}{}

	query := `select 
				id, 
				name, 
				number_of_pages,
				created_at,
				updated_at
			from books`
	filter = " WHERE true"
	order = " ORDER BY created_at"
	arrangement = " DESC"
	offset = " OFFSET 0"
	limit = " LIMIT 10"

	if req.Page > 0 {
		req.Page = (req.Page - 1) * req.Limit
		params["offset"] = req.Page
		offset = " OFFSET @offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT @limit"
	}

	cQ := `SELECT count(1) FROM books` + filter

	err = b.db.QueryRow(ctx, cQ, pgx.NamedArgs(params)).Scan(
		&resp.Count,
	)

	if err != nil {
		return resp, err
	}

	q = query + filter + order + arrangement + offset + limit

	rows, err := b.db.Query(ctx, q, pgx.NamedArgs(params))
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		book := &book_service.Book{}
		result := &Book{}

		err = rows.Scan(
			&result.Id,
			&result.Name,
			&result.NumberOfPages,
			&result.CreatedAt,
			&result.UpdatedAt,
		)

		if err != nil {
			return resp, err
		}

		if result.CreatedAt.Valid {
			book.CreatedAt = result.CreatedAt.String
		}

		if result.UpdatedAt.Valid {
			book.UpdatedAt = result.UpdatedAt.String
		}

		book.Id = result.Id
		book.Name = result.Name
		book.NumberOfPages = result.NumberOfPages

		resp.Books = append(resp.Books, book)
	}

	return
}

func (b *bookRepo) Update(ctx context.Context, req *book_service.UpdateBookRequest) (rowsAffected int64, err error) {
	query := `update books SET
		name = @name,
		number_of_pages = @number_of_pages,
		updated_at = now()
	WHERE
		id = @id`

	params := map[string]interface{}{
		"id":              req.Book.Id,
		"name":            req.Book.Name,
		"number_of_pages": req.Book.NumberOfPages,
	}

	result, err := b.db.Exec(ctx, query, pgx.NamedArgs(params))
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (b *bookRepo) PatchUpdate(ctx context.Context, req *book_service.PatchUpdateRequest) (rowsAffected int64, err error) {
	columns := ""
	params := make(map[string]interface{})

	for _, v := range req.Params {
		params[v.Key] = v.Value
		columns += " " + v.Key + " = @" + v.Key + ","
	}

	query := `
		UPDATE 
			books 
		SET` + columns + `
	updated_at = now()
	WHERE 
		id = @id`

	params["id"] = req.Id

	result, err := b.db.Exec(ctx, query, pgx.NamedArgs(params))
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (b *bookRepo) Delete(ctx context.Context, req *book_service.BookPrimaryKey) (rowsAffected int64, err error) {
	query := `delete from books where id = $1`

	result, err := b.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
