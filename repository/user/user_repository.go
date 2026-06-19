package user

import (
	"context"
	"errors"

	"restaurant-mie-api/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(
	ctx context.Context,
	user *domain.User,
) error {

	return r.DB.
		WithContext(ctx).
		Create(user).
		Error
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*domain.User, error) {

	var user domain.User

	err := r.DB.
		WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrUserNotFound
	}

	return &user, err
}

func (r *UserRepository) FindByID(
	ctx context.Context,
	id uint,
) (*domain.User, error) {

	var user domain.User

	err := r.DB.
		WithContext(ctx).
		First(&user, id).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrUserNotFound
	}

	return &user, err
}

func (r *UserRepository) ExistsByEmail(
	ctx context.Context,
	email string,
) (bool, error) {

	var count int64

	err := r.DB.
		WithContext(ctx).
		Model(&domain.User{}).
		Where("email = ?", email).
		Count(&count).
		Error

	return count > 0, err
}
