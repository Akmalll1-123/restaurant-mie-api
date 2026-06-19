package table

import (
	"context"

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
	req CreateTableRequest,
) error {

	table := &domain.Table{
		TableNo:  req.TableNo,
		Capacity: req.Capacity,
		IsActive: true,
	}

	return s.repo.Create(
		ctx,
		table,
	)
}

func (s *Service) GetAll(
	ctx context.Context,
) ([]domain.Table, error) {

	return s.repo.FindAll(
		ctx,
	)
}

func (s *Service) GetByID(
	ctx context.Context,
	id uint,
) (*domain.Table, error) {

	return s.repo.FindByID(
		ctx,
		id,
	)
}

func (s *Service) Update(
	ctx context.Context,
	id uint,
	req UpdateTableRequest,
) error {

	table, err := s.repo.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return err
	}

	table.TableNo = req.TableNo
	table.Capacity = req.Capacity
	table.IsActive = req.IsActive

	return s.repo.Update(
		ctx,
		table,
	)
}

func (s *Service) Delete(
	ctx context.Context,
	id uint,
) error {

	return s.repo.Delete(
		ctx,
		id,
	)
}
