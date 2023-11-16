package services

import (
	"q3/helpers"
	"strings"
)

type phoneNumberService struct {
}

// NumberFormat implements PhoneNumberService
func (*phoneNumberService) NumberFormat(number string) helpers.PhoneNumberResponse {
	var response helpers.PhoneNumberResponse

	var result string

	if strings.HasPrefix(number, "08") {
		result = "62" + number[1:]
	} else if strings.HasPrefix(number, "+62") {
		result = "" + number[1:]
	}

	if strings.HasPrefix(number, "01") {
		result = "60" + number[1:]
	} else if strings.HasPrefix(number, "+60") {
		result = "" + number[1:]
	}

	response.Data = result

	return response
}

type PhoneNumberService interface {
	NumberFormat(number string) helpers.PhoneNumberResponse
}

func NewPhoneNumberService() PhoneNumberService {
	return &phoneNumberService{}
}
