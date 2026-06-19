package menu

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
	req CreateMenuRequest,
) error {

	menu := &domain.Menu{
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price,
		Stock:    req.Stock,
		IsActive: true,
	}

	return s.repo.Create(
		ctx,
		menu,
	)
}

func (s *Service) GetAll(
	ctx context.Context,
	page int,
	limit int,
	search string,
) ([]domain.Menu, int64, error) {

	return s.repo.FindAll(
		ctx,
		page,
		limit,
		search,
	)
}

func (s *Service) GetByID(
	ctx context.Context,
	id uint,
) (*domain.Menu, error) {

	return s.repo.FindByID(
		ctx,
		id,
	)
}

func (s *Service) Update(
	ctx context.Context,
	id uint,
	req UpdateMenuRequest,
) error {

	menu, err := s.repo.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return err
	}

	menu.Name = req.Name
	menu.Category = req.Category
	menu.Price = req.Price
	menu.Stock = req.Stock
	menu.IsActive = req.IsActive

	return s.repo.Update(
		ctx,
		menu,
	)
}

func (s *Service) Delete(
	ctx context.Context,
	id uint,
) error {

	_, err := s.repo.FindByID(
		ctx,
		id,
	)

	if err != nil {
		return err
	}

	return s.repo.Delete(
		ctx,
		id,
	)
}
