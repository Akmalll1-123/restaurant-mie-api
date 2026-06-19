package menu

import (
	"context"

	"restaurant-mie-api/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		menu *domain.Menu,
	) error

	FindAll(
		ctx context.Context,
		page int,
		limit int,
		search string,
	) ([]domain.Menu, int64, error)

	FindByID(
		ctx context.Context,
		id uint,
	) (*domain.Menu, error)

	Update(
		ctx context.Context,
		menu *domain.Menu,
	) error

	Delete(
		ctx context.Context,
		id uint,
	) error
}
