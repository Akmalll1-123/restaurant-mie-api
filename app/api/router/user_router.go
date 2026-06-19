package router

import (
	usercontroller "restaurant-mie-api/app/api/controller/user"
	userrepo "restaurant-mie-api/repository/user"
	userservice "restaurant-mie-api/service/user"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func registerUserRoutes(
	api *echo.Group,
	db *gorm.DB,
) {

	userRepository := userrepo.NewUserRepository(db)

	userService := userservice.NewService(
		userRepository,
	)

	userController := usercontroller.NewUserController(
		userService,
	)

	auth := api.Group("/auth")

	auth.POST(
		"/register",
		userController.Register,
	)

	auth.POST(
		"/login",
		userController.Login,
	)
}
