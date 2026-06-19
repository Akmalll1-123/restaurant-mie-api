package menu

import (
	"context"
	"errors"

	"restaurant-mie-api/domain"

	"gorm.io/gorm"
)

type MenuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository(
	db *gorm.DB,
) *MenuRepository {

	return &MenuRepository{
		DB: db,
	}
}

func (r *MenuRepository) Create(
	ctx context.Context,
	menu *domain.Menu,
) error {

	return r.DB.
		WithContext(ctx).
		Create(menu).
		Error
}

func (r *MenuRepository) FindByID(
	ctx context.Context,
	id uint,
) (*domain.Menu, error) {

	var menu domain.Menu

	err := r.DB.
		WithContext(ctx).
		First(&menu, id).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrMenuNotFound
	}

	return &menu, err
}

func (r *MenuRepository) FindAll(
	ctx context.Context,
	page int,
	limit int,
	search string,
) ([]domain.Menu, int64, error) {

	var menus []domain.Menu
	var total int64

	query := r.DB.
		WithContext(ctx).
		Model(&domain.Menu{})

	if search != "" {
		query = query.Where(
			"name ILIKE ?",
			"%"+search+"%",
		)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	err := query.
		Order("id DESC").
		Limit(limit).
		Offset(offset).
		Find(&menus).
		Error

	return menus, total, err
}

func (r *MenuRepository) Update(
	ctx context.Context,
	menu *domain.Menu,
) error {

	return r.DB.
		WithContext(ctx).
		Save(menu).
		Error
}

func (r *MenuRepository) Delete(
	ctx context.Context,
	id uint,
) error {

	return r.DB.
		WithContext(ctx).
		Delete(&domain.Menu{}, id).
		Error
}
