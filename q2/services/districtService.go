package services

import (
	"q2/helpers"
	"q2/repositories"

	"gorm.io/gorm"
)

type districtService struct {
	districtRepo repositories.DistrictRepository
}

// GetDistrict implements ChargeService
func (service *districtService) GetDistrict() helpers.DistrictResponse {
	var response helpers.DistrictResponse

	data, err := service.districtRepo.GetDistrict()

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

type DistrictService interface {
	GetDistrict() helpers.DistrictResponse
}

func NewDistrictService(db *gorm.DB) DistrictService {
	return &districtService{districtRepo: repositories.NewDistrictRepository(db)}
}
