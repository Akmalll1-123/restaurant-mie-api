package router

import (
	paymentcontroller "restaurant-mie-api/app/api/controller/payment"
	paymentrepo "restaurant-mie-api/repository/payment"
	paymentservice "restaurant-mie-api/service/payment"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func registerPaymentRoutes(
	api *echo.Group,
	db *gorm.DB,
) {

	paymentRepository := paymentrepo.NewPaymentRepository(
		db,
	)

	paymentService := paymentservice.NewService(
		paymentRepository,
	)

	paymentController := paymentcontroller.NewPaymentController(
		paymentService,
	)

	api.POST(
		"/orders/:id/payment",
		paymentController.Create,
	)

	api.GET(
		"/orders/:id/payment",
		paymentController.GetByOrderID,
	)
}
