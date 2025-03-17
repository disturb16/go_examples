package paymentstore

import (
	"context"
	"database/sql"
)

type PaymentStorer interface {
	Save(ctx context.Context, accountID, paymentID string, total int64) error
}

type store struct {
	db *sql.DB
}

func New(db *sql.DB) PaymentStorer {
	return &store{
		db: db,
	}
}

func (s *store) Save(ctx context.Context, accountID, paymentID string, total int64) error {
	query := `
	insert into payments (account_id, payment_id, total)
	values(?,?,?);
	`

	_, err := s.db.ExecContext(ctx, query, accountID, paymentID, total)
	if err != nil {
		return err
	}

	return nil
}
