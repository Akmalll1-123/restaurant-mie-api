package domain

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrMenuNotFound    = errors.New("menu not found")
	ErrOrderNotFound   = errors.New("order not found")
	ErrPaymentNotFound = errors.New("payment not found")

	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidRole        = errors.New("invalid role")

	ErrInsufficientStock = errors.New("insufficient stock")
	ErrOrderNotPending   = errors.New("order is not pending")
)
