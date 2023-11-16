package controllers

import (
	"net/http"
	"q1/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ChargeController struct {
	chargeService services.ChargeService
}

func (controller ChargeController) GetAll(c echo.Context) error {
	result := controller.chargeService.GetAll()

	return c.JSON(http.StatusOK, result)
}

func NewChargeController(db *gorm.DB) ChargeController {
	return ChargeController{chargeService: services.NewChargeService(db)}
}
