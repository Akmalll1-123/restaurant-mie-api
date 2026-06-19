package table

import (
	"context"

	"restaurant-mie-api/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		table *domain.Table,
	) error

	FindAll(
		ctx context.Context,
	) ([]domain.Table, error)

	FindByID(
		ctx context.Context,
		id uint,
	) (*domain.Table, error)

	Update(
		ctx context.Context,
		table *domain.Table,
	) error

	Delete(
		ctx context.Context,
		id uint,
	) error
}
