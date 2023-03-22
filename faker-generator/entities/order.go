package entities

import "time"

type Order struct {
	ID            int64     `db:"id"`
	UserID        int64     `db:"user_id"`
	Date          time.Time `db:"order_date"`
	PaymentMethod string    `db:"payment_method" faker:"oneof: cc, paypal, check, cash"`
	PaymentRef    string    `db:"payment_reference" faker:"uuid_digit"`
}
