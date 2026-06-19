package user

import (
	"context"

	"restaurant-mie-api/domain"
	bcryptutil "restaurant-mie-api/util/bcrypt"
	jwtutil "restaurant-mie-api/util/jwt"
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

func (s *Service) Register(
	ctx context.Context,
	req RegisterRequest,
) error {

	exists, err := s.repo.ExistsByEmail(
		ctx,
		req.Email,
	)

	if err != nil {
		return err
	}

	if exists {
		return domain.ErrEmailAlreadyExists
	}

	hash, err := bcryptutil.Hash(
		req.Password,
	)

	if err != nil {
		return err
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hash,
		Role:     req.Role,
	}

	return s.repo.Create(
		ctx,
		user,
	)
}

func (s *Service) Login(
	ctx context.Context,
	req LoginRequest,
) (string, error) {

	user, err := s.repo.FindByEmail(
		ctx,
		req.Email,
	)

	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	if err := bcryptutil.Compare(
		user.Password,
		req.Password,
	); err != nil {
		return "", domain.ErrInvalidCredentials
	}

	return jwtutil.GenerateToken(
		user.ID,
		user.Email,
		user.Role,
	)
}
