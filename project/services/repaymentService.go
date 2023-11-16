package services

import (
	"project/helpers"
	"project/models"
	"project/repository"

	"gorm.io/gorm"
)

type repaymentService struct {
	repaymentRepo repository.RepaymentRepository
}

// GetPembayaran implements RepaymentService
func (service *repaymentService) GetPembayaran(repayment models.Repayment) helpers.Response {
	var responses helpers.Response

	var response []map[string]interface{}

	result, err := service.repaymentRepo.GetPembayaran(repayment.Name)
	if err != nil {
		// Handle error
		return responses
	}

	principalMonth, _ := result["principal_month"].(float64)
	interest, _ := result["interest"].(float64)
	feeStampDuty, _ := result["fee_stamp_duty"].(float64)
	fee, _ := result["fee"].(float64)
	amount, _ := result["amount"].(float64)

	item := map[string]interface{}{
		"Fee":                result["fee"],
		"Fee_stamp_duty":     result["fee_stamp_duty"],
		"Interest":           result["interest"],
		"Pokok yang dibayar": result["principal_month"],
		"Tagihan":            principalMonth + interest + feeStampDuty + fee,
		"Outstanding":        amount - principalMonth,
	}

	response = append(response, item)

	if err != nil {
		responses.Status = 500
		responses.Messages = "Failed to get items"
	} else {
		responses.Status = 200
		responses.Messages = "Success to get items"
		responses.Data = response
	}

	return responses
}

// CreateRepayment implements RepaymentService
func (service *repaymentService) CreateRepayment(repayment models.Repayment) helpers.Response {
	var response helpers.Response

	amount, errs := service.repaymentRepo.GetAmountByName(repayment.Name)

	if errs != nil {
		response.Status = 500
		response.Messages = "Failed to get items"
	}

	if amount+repayment.Amount > 5000000 {
		repayment.Fee_stamp_duty = 10000
		service.repaymentRepo.UpdateDataByName(repayment.Name)
	}

	repayment.Principal = repayment.Amount
	repayment.Principal_month = repayment.Principal / repayment.Tenor
	repayment.Fee = repayment.Principal * 0.05
	repayment.Interest = repayment.Principal * 0.0199
	repayment.Total_interest = repayment.Interest * repayment.Tenor
	repayment.Total_Payment = repayment.Principal + repayment.Total_interest
	repayment.Demand_per_month = repayment.Principal_month + repayment.Interest

	err := service.repaymentRepo.CreateRepayment(repayment)

	if err != nil {
		response.Status = 500
		response.Messages = "Failed to create new item" + err.Error()
	} else {
		response.Status = 200
		response.Messages = "Success to create new item"
	}

	return response
}

// GetAll implements RepaymentService
func (service *repaymentService) GetAll() helpers.Response {
	var response helpers.Response

	data, err := service.repaymentRepo.GetAll()

	if err != nil {
		response.Status = 500
		response.Messages = "Failed to get items"
	} else {
		response.Status = 200
		response.Messages = "Success to get items"
		response.Data = data
	}

	return response
}

type RepaymentService interface {
	CreateRepayment(repayment models.Repayment) helpers.Response
	GetAll() helpers.Response
	GetPembayaran(repayment models.Repayment) helpers.Response
}

func NewRepaymentService(db *gorm.DB) RepaymentService {
	return &repaymentService{repaymentRepo: repository.NewRepaymentRepository(db)}
}
