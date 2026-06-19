package user

import (
	"context"

	"restaurant-mie-api/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		user *domain.User,
	) error

	FindByEmail(
		ctx context.Context,
		email string,
	) (*domain.User, error)

	FindByID(
		ctx context.Context,
		id uint,
	) (*domain.User, error)

	ExistsByEmail(
		ctx context.Context,
		email string,
	) (bool, error)
}
