package payment

import (
	"context"
	"errors"

	"restaurant-mie-api/domain"
	xenditpkg "restaurant-mie-api/pkg/xendit"
)

type Service struct {
	repo Repository
}

func NewService(
	repo Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(
	ctx context.Context,
	orderID uint,
	totalPrice float64,
) error {

	// cek payment sudah pernah dibuat
	existingPayment, err := s.repo.FindByOrderID(
		ctx,
		orderID,
	)

	if err == nil && existingPayment != nil {
		return errors.New("payment already exists")
	}

	invoice, err := xenditpkg.CreateInvoice(
		orderID,
		totalPrice,
	)

	if err != nil {
		return err
	}

	payment := &domain.Payment{
		OrderID: orderID,

		XenditInvoiceID: invoice.ID,

		PaymentMethod: "XENDIT",

		InvoiceURL: invoice.InvoiceURL,

		Amount: totalPrice,

		Status: domain.PaymentPending,
	}

	return s.repo.Create(
		ctx,
		payment,
	)
}

func (s *Service) GetByOrderID(
	ctx context.Context,
	orderID uint,
) (*domain.Payment, error) {

	return s.repo.FindByOrderID(
		ctx,
		orderID,
	)
}
