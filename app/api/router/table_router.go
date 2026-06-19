package router

import (
	tablecontroller "restaurant-mie-api/app/api/controller/table"
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

	api.GET("/tables", tableController.GetAll)

	api.POST("/tables", tableController.Create)
}
