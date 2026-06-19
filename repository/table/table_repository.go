package table

import (
	"context"

	"restaurant-mie-api/domain"

	"gorm.io/gorm"
)

type TableRepository struct {
	DB *gorm.DB
}

func NewTableRepository(
	db *gorm.DB,
) *TableRepository {
	return &TableRepository{
		DB: db,
	}
}

func (r *TableRepository) Create(
	ctx context.Context,
	table *domain.Table,
) error {

	return r.DB.
		WithContext(ctx).
		Create(table).
		Error
}

func (r *TableRepository) FindAll(
	ctx context.Context,
) ([]domain.Table, error) {

	var tables []domain.Table

	err := r.DB.
		WithContext(ctx).
		Find(&tables).
		Error

	return tables, err
}

func (r *TableRepository) FindByID(
	ctx context.Context,
	id uint,
) (*domain.Table, error) {

	var table domain.Table

	err := r.DB.
		WithContext(ctx).
		First(&table, id).
		Error

	if err != nil {
		return nil, err
	}

	return &table, nil
}

func (r *TableRepository) Update(
	ctx context.Context,
	table *domain.Table,
) error {

	return r.DB.
		WithContext(ctx).
		Save(table).
		Error
}

func (r *TableRepository) Delete(
	ctx context.Context,
	id uint,
) error {

	return r.DB.
		WithContext(ctx).
		Delete(&domain.Table{}, id).
		Error
}
