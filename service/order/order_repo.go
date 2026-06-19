package order

import (
	"context"

	"restaurant-mie-api/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		order *domain.Order,
	) error

	FindByID(
		ctx context.Context,
		id uint,
	) (*domain.Order, error)

	Update(
		ctx context.Context,
		order *domain.Order,
	) error

	DeleteItemsByOrderID(
		ctx context.Context,
		orderID uint,
	) error
}
