package domain

import (
	"time"

	"gorm.io/gorm"
)

const (
	OrderPending    = "PENDING"
	OrderProcessing = "PROCESSING"
	OrderCompleted  = "COMPLETED"
	OrderPaid       = "PAID"
	OrderCancelled  = "CANCELLED"
)

type Order struct {
	ID uint `gorm:"primaryKey"`

	OrderID string `gorm:"unique"`

	UserID uint

	CashierID *uint

	TableID uint

	Table Table `gorm:"foreignKey:TableID"`

	Status string

	TotalPrice float64

	Notes string

	Items []OrderItem `gorm:"foreignKey:OrderID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
