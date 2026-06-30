package router

import (
	ordercontroller "restaurant-mie-api/app/api/controller/order"
	"restaurant-mie-api/internal/middleware"
	orderrepo "restaurant-mie-api/repository/order"
	orderservice "restaurant-mie-api/service/order"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func registerOrderRoutes(
	api *echo.Group,
	db *gorm.DB,
) {

	orderRepository := orderrepo.NewOrderRepository(db)

	orderService := orderservice.NewService(
		orderRepository,
	)

	orderController := ordercontroller.NewOrderController(
		orderService,
	)

	kasirorderGroup := api.Group("/orders", middleware.JWTMiddleware, middleware.RoleMiddleware("KASIR"))
	{
		kasirorderGroup.PATCH("/:id/status", orderController.UpdateStatus)

		kasirorderGroup.PUT("/:id", orderController.Update)
	}

	userorderGroup := api.Group("/orders", middleware.JWTMiddleware, middleware.RoleMiddleware("USER"))
	{
		userorderGroup.POST("", orderController.Create)

		userorderGroup.GET("/:id", orderController.GetByID)
	}
}
