package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"integrations_service/genproto/book_service"
	"testing"

	"github.com/bxcodec/faker/v4"

	"github.com/stretchr/testify/assert"
)

func createBook() (string, error) {
	bookRepo := NewBookRepo(db)

	resp, err := bookRepo.Create(context.Background(), &book_service.CreateBookRequest{
		Name:          "Test",
		NumberOfPages: 10,
	})

	if err != nil {
		return "", err
	}

	return resp.Id, nil
}

func TestCreate(t *testing.T) {
	bookRepo := NewBookRepo(db)

	resp, err := bookRepo.Create(context.Background(), &book_service.CreateBookRequest{
		Name:          faker.FirstName(),
		NumberOfPages: 12,
	})

	assert.NoError(t, err)

	fmt.Print("Create Book------->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}

func TestGet(t *testing.T) {
	bookRepo := NewBookRepo(db)

	resp, err := bookRepo.Get(context.Background(), &book_service.BookPrimaryKey{
		// need to call create first so that we can get the id
		Id: "e58ae4cb-a98a-4750-b0aa-11e867d76593",
	})

	assert.NoError(t, err)

	fmt.Print("Get Book---->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}

func TestGetList(t *testing.T) {
	bookRepo := NewBookRepo(db)

	resp, err := bookRepo.GetList(context.Background(), &book_service.GetBooksListRequest{
		Page:  1,
		Limit: 3,
	})

	assert.NoError(t, err)

	fmt.Print("Get List---->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}

func TestUpdate(t *testing.T) {
	bookRepo := NewBookRepo(db)

	resp, err := bookRepo.Update(context.Background(), &book_service.UpdateBookRequest{
		Book: &book_service.Book{
			Id:            "e58ae4cb-a98a-4750-b0aa-11e867d76593",
			Name:          faker.Name(),
			NumberOfPages: 10,
		},
	})

	assert.NoError(t, err)

	fmt.Print("Update book---->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}

func TestDelete(t *testing.T) {
	bookRepo := NewBookRepo(db)

	resp, err := bookRepo.Delete(context.Background(), &book_service.BookPrimaryKey{
		Id: "e58ae4cb-a98a-4750-b0aa-11e867d76593",
	})

	assert.NoError(t, err)

	fmt.Print("Delete book---->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))
}
