package main

import (
	"q3/controllers"

	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()

	phoneController := controllers.NewPhoneNumberController()

	app.GET("/phone_number", phoneController.PhoneFormat)

	app.Start(":8080")

}
