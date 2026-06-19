package payment

import (
	"net/http"
	"strconv"

	paymentservice "restaurant-mie-api/service/payment"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	service *paymentservice.Service
}

func NewPaymentController(
	service *paymentservice.Service,
) *PaymentController {

	return &PaymentController{
		service: service,
	}
}

func (p *PaymentController) Create(
	c echo.Context,
) error {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "invalid order id",
			},
		)
	}

	err = p.service.Create(
		c.Request().Context(),
		uint(id),
		100000, // sementara hardcode
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
			"message": "payment created",
		},
	)
}

func (p *PaymentController) GetByOrderID(
	c echo.Context,
) error {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {

		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "invalid order id",
			},
		)
	}

	payment, err := p.service.GetByOrderID(
		c.Request().Context(),
		uint(id),
	)

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
		payment,
	)
}
