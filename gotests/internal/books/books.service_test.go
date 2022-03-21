package books

import (
	"context"
	"os"
	"testing"
)

var s BookService

func TestMain(m *testing.M) {

	repo := &MockDBInteractor{}
	s = NewService(repo)

	code := m.Run()
	os.Exit(code)
}

func TestListBooks(t *testing.T) {
	testCases := []struct {
		Name          string
		ExpectedError error
	}{
		{
			Name:          "List books",
			ExpectedError: nil,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			_, err := s.ListBooks(ctx)

			if err != tc.ExpectedError {
				t.Errorf("expected %v, got %v", tc.ExpectedError, err)
			}

		})
	}
}

func TestBook(t *testing.T) {
	testCases := []struct {
		Name          string
		BookID        int64
		ExpectedError error
	}{
		{
			Name:          "Book by ID",
			BookID:        1,
			ExpectedError: nil,
		},
		{
			Name:          "Invalid ID",
			BookID:        -1,
			ExpectedError: ErrInvaliBookdID,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			_, err := s.Book(ctx, tc.BookID)

			if err != tc.ExpectedError {
				t.Errorf("expected %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
