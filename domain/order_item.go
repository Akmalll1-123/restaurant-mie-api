package domain

type OrderItem struct {
	ID uint `gorm:"primaryKey"`

	OrderID uint

	MenuID uint

	Qty int

	Price float64

	Subtotal float64
}
