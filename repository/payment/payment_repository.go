package payment

import (
	"context"

	"restaurant-mie-api/domain"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(
	db *gorm.DB,
) *PaymentRepository {

	return &PaymentRepository{
		DB: db,
	}
}

func (r *PaymentRepository) Create(
	ctx context.Context,
	payment *domain.Payment,
) error {

	return r.DB.
		WithContext(ctx).
		Create(payment).
		Error
}

func (r *PaymentRepository) FindByInvoiceID(
	ctx context.Context,
	invoiceID string,
) (*domain.Payment, error) {

	var payment domain.Payment

	err := r.DB.
		WithContext(ctx).
		Where("xendit_invoice_id = ?", invoiceID).
		First(&payment).
		Error

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *PaymentRepository) Update(
	ctx context.Context,
	payment *domain.Payment,
) error {

	return r.DB.
		WithContext(ctx).
		Save(payment).
		Error
}

func (r *PaymentRepository) FindByOrderID(
	ctx context.Context,
	orderID uint,
) (*domain.Payment, error) {

	var payment domain.Payment

	err := r.DB.
		WithContext(ctx).
		Where("order_id = ?", orderID).
		First(&payment).
		Error

	if err != nil {
		return nil, err
	}

	return &payment, nil
}
