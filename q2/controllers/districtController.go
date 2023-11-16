package controllers

import (
	"net/http"
	"q2/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DistrictController struct {
	districtService services.DistrictService
}

func (controller DistrictController) GetDistric(c echo.Context) error {
	result := controller.districtService.GetDistrict()

	return c.JSON(http.StatusOK, result)
}

func NewDistrictController(db *gorm.DB) DistrictController {
	return DistrictController{districtService: services.NewDistrictService(db)}
}
