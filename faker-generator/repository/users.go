package repository

import (
	"context"
	"faker-generator/entities"
	"log"
)

func (r *Repository) InsertUser(ctx context.Context, u entities.User) (int64, error) {

	query := `INSERT INTO users (first_name, last_name, email, phone)
		VALUES (:first_name, :last_name, :email, :phone)`

	result, err := r.db.NamedExecContext(ctx, query, u)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *Repository) AllUsers(ctx context.Context) ([]entities.User, error) {
	users := []entities.User{}

	query := `SELECT * FROM users`

	err := r.db.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, err
	}

	log.Println("Total users:", len(users))

	return users, nil
}

func (r *Repository) InsertUserWithOrders(ctx context.Context, u entities.User, orders []entities.Order) error {
	var err error

	u.ID, err = r.InsertUser(ctx, u)
	if err != nil {
		log.Panic(err)
	}

	for i := range orders {
		orders[i].UserID = u.ID
	}

	for _, order := range orders {
		err = r.InsertOrder(ctx, order)
		if err != nil {
			log.Panic(err)
		}
	}

	return nil
}
