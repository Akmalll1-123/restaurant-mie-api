package user

import (
	"net/http"

	userservice "restaurant-mie-api/service/user"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service *userservice.Service
}

func NewUserController(
	service *userservice.Service,
) *UserController {

	return &UserController{
		service: service,
	}
}

func (u *UserController) Register(
	c echo.Context,
) error {

	var req userservice.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "invalid request",
			},
		)
	}

	if err := u.service.Register(
		c.Request().Context(),
		req,
	); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusCreated,
		map[string]string{
			"message": "success",
		},
	)
}

func (u *UserController) Login(
	c echo.Context,
) error {

	var req userservice.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "invalid request",
			},
		)
	}

	token, err := u.service.Login(
		c.Request().Context(),
		req,
	)

	if err != nil {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		userservice.LoginResponse{
			Token: token,
		},
	)
}
