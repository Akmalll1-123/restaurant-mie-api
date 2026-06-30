package order

import (
	"net/http"
	"strconv"

	custommiddleware "restaurant-mie-api/internal/middleware"
	orderservice "restaurant-mie-api/service/order"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	service *orderservice.Service
}

func NewOrderController(
	service *orderservice.Service,
) *OrderController {

	return &OrderController{
		service: service,
	}
}

func (o *OrderController) Create(
	c echo.Context,
) error {

	var req orderservice.CreateOrderRequest

	if err := c.Bind(&req); err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	claims, _ := c.Get("user").(*custommiddleware.JWTClaims)

	err := o.service.Create(
		c.Request().Context(),
		claims.ID,
		req,
	)

	if err != nil {

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
			"message": "order created",
		},
	)
}

func (o *OrderController) GetByID(
	c echo.Context,
) error {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "invalid id",
			},
		)
	}

	order, err := o.service.GetByID(
		c.Request().Context(),
		uint(id),
	)

	claims, _ := c.Get("user").(*custommiddleware.JWTClaims)

	if order.UserID != claims.ID {
		return c.JSON(http.StatusForbidden, map[string]string{
			"message": "You do not have permission to view this order",
		})
	}

	if err != nil {

		return c.JSON(
			http.StatusNotFound,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		order,
	)
}

func (o *OrderController) UpdateStatus(
	c echo.Context,
) error {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "invalid id",
			},
		)
	}

	var req orderservice.UpdateOrderStatusRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	err = o.service.UpdateStatus(
		c.Request().Context(),
		uint(id),
		req.Status,
	)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"message": "status updated",
		},
	)
}

func (o *OrderController) Update(
	c echo.Context,
) error {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "invalid id",
			},
		)
	}

	var req orderservice.UpdateOrderRequest

	if err := c.Bind(
		&req,
	); err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	err = o.service.Update(
		c.Request().Context(),
		uint(id),
		req,
	)

	if err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"message": "order updated",
		},
	)
}
