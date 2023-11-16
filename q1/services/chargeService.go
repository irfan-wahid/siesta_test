package services

import (
	"q1/helpers"
	"q1/repositories"

	"gorm.io/gorm"
)

type chargeService struct {
	chargeRepo repositories.ChargeRepository
}

// GetAll implements ChargeService
func (service *chargeService) GetAll() helpers.ChargeResponse {
	var responses helpers.ChargeResponse

	data, err := service.chargeRepo.GetCharge()

	var response []map[string]interface{}

	// Loop melalui data dan membentuk respons JSON
	for _, charge := range data {
		chargeType, err := service.chargeRepo.GetChargeType(charge.Charge_ID)
		if err != nil {
			continue
		}

		taxType, err := service.chargeRepo.GetChargeTax(charge.Charge_ID)

		total := charge.Amount * (taxType.Percentage / 100)

		// Membuat respons JSON untuk setiap data
		item := map[string]interface{}{
			"Charge_ID":    charge.Charge_ID,
			"Selling_date": charge.Selling_date,
			"Amount":       charge.Amount,
			"Charge_type":  chargeType.Charge_type,
			"Tax_Total":    total,
		}

		// Menambahkan respons JSON ke slice
		response = append(response, item)
	}

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

type ChargeService interface {
	GetAll() helpers.ChargeResponse
}

func NewChargeService(db *gorm.DB) ChargeService {
	return &chargeService{chargeRepo: repositories.NewChargeRepository(db)}
}
