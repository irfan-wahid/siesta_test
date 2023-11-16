package controllers

import (
	"fmt"
	"net/http"
	"project/models"
	"project/services"
	"time"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RepaymentController struct {
	repaymentService services.RepaymentService
	validate         vl.Validate
}

func (controller RepaymentController) CreateRepayment(c echo.Context) error {
	type payload struct {
		Name      string    `json:"name" validate:"required"`
		Loan_date time.Time `json:"loan_date" validate:"required"`
		Amount    float64   `json:"amount" validate:"required"`
		Tenor     float64   `json:"tenor" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return err
	}

	result := controller.repaymentService.CreateRepayment(models.Repayment{Name: payloadValidator.Name, Loan_date: payloadValidator.Loan_date, Amount: payloadValidator.Amount, Tenor: payloadValidator.Tenor})

	if result.Status != 200 {
		fmt.Printf("Create error: %s", result.Messages)
	}
	return c.JSON(http.StatusOK, result)

}

func (controller RepaymentController) GetAll(c echo.Context) error {
	result := controller.repaymentService.GetAll()

	return c.JSON(http.StatusOK, result)
}

func (controller RepaymentController) GetPayment(c echo.Context) error {
	type payload struct {
		Name string `json:"name" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return err
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return err
	}

	result := controller.repaymentService.GetPembayaran(models.Repayment{Name: payloadValidator.Name})

	if result.Status != 200 {
		fmt.Printf("Create error: %s", result.Messages)
	}
	return c.JSON(http.StatusOK, result)
}

func NewRepaymentController(db *gorm.DB) RepaymentController {
	controller := RepaymentController{
		repaymentService: services.NewRepaymentService(db),
		validate:         *vl.New(),
	}

	return controller
}
