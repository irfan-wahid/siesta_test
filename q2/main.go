package main

import (
	"q2/config"
	"q2/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	app := echo.New()

	chargeController := controllers.NewDistrictController(db)

	app.GET("/get_all", chargeController.GetDistric)

	app.Start(":8080")

}
