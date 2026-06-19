package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(
	e *echo.Echo,
	db *gorm.DB,
) {
	api := e.Group("/api/v1")

	registerUserRoutes(api, db)
	registerMenuRoutes(api, db)
	registerOrderRoutes(api, db)
	registerPaymentRoutes(api, db)
	registerTableRoutes(api, db)
}
