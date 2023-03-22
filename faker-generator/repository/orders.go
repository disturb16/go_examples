package repository

import (
	"context"
	"faker-generator/entities"
)

func (r *Repository) InsertOrder(ctx context.Context, a entities.Order) error {

	query := `INSERT INTO orders (user_id, order_date, payment_method, payment_reference)
		VALUES (:user_id, :order_date, :payment_method, :payment_reference)`

	_, err := r.db.NamedExecContext(ctx, query, a)

	return err
}

func (r *Repository) OrderByUserID(ctx context.Context, userID int64) ([]entities.Order, error) {
	orders := []entities.Order{}

	query := `SELECT * FROM orders WHERE user_id = ?`

	err := r.db.SelectContext(ctx, &orders, query, userID)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
