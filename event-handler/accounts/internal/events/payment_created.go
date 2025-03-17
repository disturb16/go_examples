package events

type PaymentCreated struct {
	AccountID string `json:"account_id"`
	PaymentID string `json:"payment_id"`
	TotalPaid int64  `json:"total_paid"`
	CreatedAt string `json:"created_at"`
}
