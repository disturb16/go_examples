package entities

type User struct {
	ID        int64  `db:"id"`
	FirstName string `db:"first_name" faker:"first_name"`
	LastName  string `db:"last_name" faker:"last_name"`
	Email     string `db:"email" faker:"email"`
	Phone     string `db:"phone" faker:"phone_number"`
}
