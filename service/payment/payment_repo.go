package payment

import (
	"context"

	"restaurant-mie-api/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		payment *domain.Payment,
	) error

	FindByOrderID(
		ctx context.Context,
		orderID uint,
	) (*domain.Payment, error)

	FindByInvoiceID(
		ctx context.Context,
		invoiceID string,
	) (*domain.Payment, error)

	Update(
		ctx context.Context,
		payment *domain.Payment,
	) error
}
