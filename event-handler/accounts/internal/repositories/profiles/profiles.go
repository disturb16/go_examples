package profiles

import (
	"context"
	"database/sql"
	"errors"
)

type repo struct {
	db *sql.DB
}

type ProfileStorer interface {
	Save(ctx context.Context, accountID, subscriptionID string) error
	UpdateStatus(ctx context.Context, accountID, status string) error
	ByAccountID(ctx context.Context, id string) (ProfileRecord, error)
}

type ProfileRecord struct {
	AccountID      string `db:"account_id"`
	SubscriptionID string `db:"subscription_id"`
	Status         string `db:"status"`
}

const (
	ProfileStatusActive  = "active"
	ProfileStatusPending = "pending"
)

func New(db *sql.DB) ProfileStorer {
	return &repo{
		db: db,
	}
}

func (r *repo) Save(ctx context.Context, accountID, subscriptionID string) error {
	query := `
	insert into profiles (account_id, subscription_id, status)
	values (?, ?, ?);
	`

	_, err := r.db.ExecContext(ctx, query, accountID, subscriptionID, ProfileStatusPending)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) UpdateStatus(ctx context.Context, accountID, status string) error {
	query := `
	update profiles set
		status = ?
	where account_id = ?;`

	_, err := r.db.ExecContext(ctx, query, status, accountID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) ByAccountID(ctx context.Context, id string) (ProfileRecord, error) {
	query := `
	select
		account_id,
		subscription_id,
		status
	from profiles
	where account_id = ?;
	`

	p := ProfileRecord{}
	row := r.db.QueryRowContext(ctx, query, id)

	err := row.Scan(&p.AccountID, &p.SubscriptionID, &p.Status)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return p, err
	}

	return p, nil
}
