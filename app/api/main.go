package main

import (
	"os"

	"restaurant-mie-api/app/api/router"
	"restaurant-mie-api/util/db"

	"github.com/labstack/echo/v4"
)

func main() {

	database := db.Connect()

	e := echo.New()

	router.Register(
		e,
		database,
	)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(
		e.Start(":" + port),
	)
}
