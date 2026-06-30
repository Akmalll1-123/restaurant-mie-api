package order

import (
	"context"
	"errors"
	"fmt"
	"time"

	"restaurant-mie-api/domain"
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
	userID uint,
	req CreateOrderRequest,
) error {

	order := &domain.Order{
		OrderNo: fmt.Sprintf(
			"ORD-%d",
			time.Now().UnixNano(),
		),

		UserID: userID,

		TableNo: req.TableID,

		Status: domain.OrderPending,

		Notes: req.Notes,
	}

	var total float64

	for _, item := range req.Items {

		subtotal := float64(item.Qty) * 10000

		total += subtotal

		order.Items = append(
			order.Items,
			domain.OrderItem{
				MenuID:   item.MenuID,
				Qty:      item.Qty,
				Price:    10000,
				Subtotal: subtotal,
			},
		)
	}

	order.TotalPrice = total

	return s.repo.Create(
		ctx,
		order,
	)
}

func (s *Service) GetByID(
	ctx context.Context,
	id uint,
) (*domain.Order, error) {

	return s.repo.FindByID(
		ctx,
		id,
	)
}

func (s *Service) UpdateStatus(
	ctx context.Context,
	id uint,
	status string,
) error {

	order, err := s.repo.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return err
	}

	switch status {
	case domain.OrderPending,
		domain.OrderProcessing,
		domain.OrderCompleted,
		domain.OrderPaid,
		domain.OrderCancelled:

	default:
		return errors.New("invalid status")
	}

	order.Status = status

	return s.repo.Update(
		ctx,
		order,
	)
}

func (s *Service) Update(
	ctx context.Context,
	id uint,
	req UpdateOrderRequest,
) error {

	order, err := s.repo.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return err
	}

	if order.Status != domain.OrderPending {
		return errors.New(
			"order cannot be modified",
		)
	}

	err = s.repo.DeleteItemsByOrderID(
		ctx,
		id,
	)

	if err != nil {
		return err
	}

	order.Items = nil

	var total float64

	for _, item := range req.Items {

		subtotal := float64(item.Qty) * 10000

		total += subtotal

		order.Items = append(
			order.Items,
			domain.OrderItem{
				MenuID:   item.MenuID,
				Qty:      item.Qty,
				Price:    10000,
				Subtotal: subtotal,
			},
		)
	}

	order.TotalPrice = total
	order.Notes = req.Notes
	order.TableNo = req.TableID

	return s.repo.Update(
		ctx,
		order,
	)
}
