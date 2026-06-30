package router

import (
	tablecontroller "restaurant-mie-api/app/api/controller/table"
	"restaurant-mie-api/internal/middleware"
	tablerepo "restaurant-mie-api/repository/table"
	tableservice "restaurant-mie-api/service/table"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func registerTableRoutes(
	api *echo.Group,
	db *gorm.DB,
) {

	tableRepository := tablerepo.NewTableRepository(
		db,
	)

	tableService := tableservice.NewService(
		tableRepository,
	)

	tableController := tablecontroller.NewTableController(
		tableService,
	)

	tableGroup := api.Group("/tables")
	{
		tableGroup.GET("", tableController.GetAll)

		kasirTableGroup := tableGroup.Group("", middleware.RoleMiddleware("KASIR"))
		{
			kasirTableGroup.POST("", tableController.Create)
		}
	}
}
