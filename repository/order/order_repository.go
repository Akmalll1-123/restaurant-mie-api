package order

import (
	"context"

	"restaurant-mie-api/domain"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(
	db *gorm.DB,
) *OrderRepository {

	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) Create(
	ctx context.Context,
	order *domain.Order,
) error {

	return r.DB.
		WithContext(ctx).
		Create(order).
		Error
}

func (r *OrderRepository) FindByID(
	ctx context.Context,
	id uint,
) (*domain.Order, error) {

	var order domain.Order

	err := r.DB.
		WithContext(ctx).
		Preload("Items").
		First(&order, id).
		Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *OrderRepository) Update(
	ctx context.Context,
	order *domain.Order,
) error {

	return r.DB.
		WithContext(ctx).
		Session(&gorm.Session{
			FullSaveAssociations: true,
		}).
		Updates(order).
		Error
}

func (r *OrderRepository) DeleteItemsByOrderID(
	ctx context.Context,
	orderID uint,
) error {

	return r.DB.
		WithContext(ctx).
		Where(
			"order_id = ?",
			orderID,
		).
		Delete(
			&domain.OrderItem{},
		).
		Error
}
