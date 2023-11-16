package main

import (
	"project/config"
	"project/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()

	route := echo.New()

	repaymentController := controllers.NewRepaymentController(db)
	route.POST("/create", repaymentController.CreateRepayment)
	route.GET("/get_all", repaymentController.GetAll)
	route.GET("/pembayaran", repaymentController.GetPayment)

	route.Start(":8080")
}
