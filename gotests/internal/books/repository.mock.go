package books

import (
	"context"
	"gotests/internal/books/models"
)

type RepositoryMocked struct{}

func (rm *RepositoryMocked) AllBooks(ctx context.Context) ([]*models.Book, error) {
	return []*models.Book{}, nil
}

func (rm *RepositoryMocked) BookByID(ctx context.Context, id int64) (*models.Book, error) {
	return &models.Book{}, nil
}

func (rm *RepositoryMocked) SaveBook(ctx context.Context, book *models.Book) error {
	return nil
}
