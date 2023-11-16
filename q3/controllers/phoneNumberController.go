package controllers

import (
	"net/http"
	"q3/services"

	"github.com/labstack/echo/v4"

	vl "github.com/go-playground/validator/v10"
)

type PhoneNumberController struct {
	phoneNumberService services.PhoneNumberService
	validate           vl.Validate
}

func (controller PhoneNumberController) PhoneFormat(c echo.Context) error {
	type payload struct {
		Nomor string `json:"nomor" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	result := controller.phoneNumberService.NumberFormat(payloadValidator.Nomor)

	return c.JSON(http.StatusOK, result)
}

func NewPhoneNumberController() PhoneNumberController {
	return PhoneNumberController{phoneNumberService: services.NewPhoneNumberService(), validate: *vl.New()}
}
