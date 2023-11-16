package main

import (
	"q1/config"
	"q1/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	app := echo.New()

	chargeController := controllers.NewChargeController(db)

	app.GET("/get_all", chargeController.GetAll)

	app.Start(":8080")

}
