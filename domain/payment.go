package domain

import (
	"time"
)

const (
	PaymentPending = "PENDING"
	PaymentSuccess = "SUCCESS"
	PaymentFailed  = "FAILED"
	PaymentExpired = "EXPIRED"
)

type Payment struct {
	ID uint `gorm:"primaryKey"`

	OrderID uint

	XenditInvoiceID string

	PaymentMethod string

	Amount float64

	Status string

	InvoiceURL string

	PaidAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
