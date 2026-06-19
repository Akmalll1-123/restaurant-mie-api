package router

import (
	ordercontroller "restaurant-mie-api/app/api/controller/order"
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

	api.POST("/orders", orderController.Create)

	api.GET("/orders/:id", orderController.GetByID)

	api.PATCH("/orders/:id/status", orderController.UpdateStatus)

	api.PUT("/orders/:id", orderController.Update)
}
