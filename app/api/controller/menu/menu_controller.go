package menu

import (
	"net/http"
	"strconv"

	menuservice "restaurant-mie-api/service/menu"

	"github.com/labstack/echo/v4"
)

type MenuController struct {
	service *menuservice.Service
}

func NewMenuController(
	service *menuservice.Service,
) *MenuController {

	return &MenuController{
		service: service,
	}
}

func (m *MenuController) Create(
	c echo.Context,
) error {

	var req menuservice.CreateMenuRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": err.Error(),
			},
		)
	}

	if err := m.service.Create(
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

func (m *MenuController) GetAll(
	c echo.Context,
) error {

	page, _ := strconv.Atoi(
		c.QueryParam("page"),
	)

	limit, _ := strconv.Atoi(
		c.QueryParam("limit"),
	)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	search := c.QueryParam("search")

	menus, total, err := m.service.GetAll(
		c.Request().Context(),
		page,
		limit,
		search,
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
		map[string]interface{}{
			"data":  menus,
			"total": total,
			"page":  page,
			"limit": limit,
		},
	)
}
