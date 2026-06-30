package router

import (
	menucontroller "restaurant-mie-api/app/api/controller/menu"
	"restaurant-mie-api/internal/middleware"
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

	menusGroup := api.Group("/menus", middleware.JWTMiddleware)
	{
		authenticatedMenusGroup := menusGroup.Group("", middleware.RoleMiddleware("USER", "KASIR"))
		{
			authenticatedMenusGroup.GET("", menuController.GetAll)
		}

		userMenusGroup := menusGroup.Group("", middleware.RoleMiddleware("KASIR"))
		{
			userMenusGroup.POST("", menuController.Create)
		}
	}
}
