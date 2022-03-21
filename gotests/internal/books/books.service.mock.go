package books

import (
	"context"
	"gotests/internal/books/models"
)

type ServiceMocked struct{}

func (sm *ServiceMocked) ListBooks(ctx context.Context) ([]*models.Book, error) {
	return nil, nil
}

func (sm *ServiceMocked) Book(ctx context.Context, id int64) (*models.Book, error) {
	return nil, nil
}

func (sm *ServiceMocked) SaveBook(ctx context.Context, book *models.Book) error {
	return nil
}
