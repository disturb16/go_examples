package repository

import "context"

type User struct {
	ID        int64
	FirstName string
	LastName  string
}

// GetUserByID returns a user by id
func (r *Repository) GetUserByID(ctx context.Context, id int64) (*User, error) {

	u := &User{}
	row := r.db.QueryRowContext(ctx, "SELECT ID, FirstName, LastName FROM users WHERE id = ?", id)

	err := row.Scan(u.ID, u.FirstName, u.LastName)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUsers returns all users
func (r *Repository) GetUsers(ctx context.Context) ([]*User, error) {

	users := []*User{}

	rows, err := r.db.QueryContext(ctx, "SELECT ID, FirstName, LastName FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

// SaveUser saves a user
func (r *Repository) SaveUser(ctx context.Context, firstName, lastName string) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO users (FirstName, LastName) VALUES (?, ?)",
		firstName,
		lastName,
	)
	if err != nil {
		return err
	}

	return nil
}
