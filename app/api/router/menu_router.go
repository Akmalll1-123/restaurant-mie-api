package router

import (
	menucontroller "restaurant-mie-api/app/api/controller/menu"
	menurepo "restaurant-mie-api/repository/menu"
	menuservice "restaurant-mie-api/service/menu"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func registerMenuRoutes(
	api *echo.Group,
	db *gorm.DB,
) {

	menuRepository := menurepo.NewMenuRepository(
		db,
	)

	menuService := menuservice.NewService(
		menuRepository,
	)

	menuController := menucontroller.NewMenuController(
		menuService,
	)

	api.GET(
		"/menus",
		menuController.GetAll,
	)

	api.POST(
		"/menus",
		menuController.Create,
	)
}
