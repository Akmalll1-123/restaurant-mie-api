package table

import (
	"net/http"

	tableservice "restaurant-mie-api/service/table"

	"github.com/labstack/echo/v4"
)

type TableController struct {
	service *tableservice.Service
}

func NewTableController(
	service *tableservice.Service,
) *TableController {

	return &TableController{
		service: service,
	}
}

func (t *TableController) Create(
	c echo.Context,
) error {

	var req tableservice.CreateTableRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	if err := t.service.Create(
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
			"message": "table created",
		},
	)
}

func (t *TableController) GetAll(
	c echo.Context,
) error {

	tables, err := t.service.GetAll(
		c.Request().Context(),
	)

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		tables,
	)
}
