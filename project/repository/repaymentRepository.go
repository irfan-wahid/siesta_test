package repository

import (
	"log"
	"project/models"

	"gorm.io/gorm"
)

type dbRepayment struct {
	Conn *gorm.DB
}

// GetDataByName implements RepaymentRepository
func (db *dbRepayment) GetDataByName(name string) ([]models.Repayment, error) {
	var data []models.Repayment

	result := db.Conn.Where("name", name).Find(&data).Error

	return data, result
}

func (db *dbRepayment) GetPembayaran(name string) (map[string]interface{}, error) {
	var result map[string]interface{}

	err := db.Conn.Model(&models.Repayment{}).
		Select("sum(fee) as fee, fee_stamp_duty, sum(interest) as interest, sum(principal_month) as principal_month, sum(demand_per_month) as demand, sum(amount) as amount").
		Where("name = ?", name).
		Group("name, fee_stamp_duty").
		Order("id").
		Limit(1).
		Scan(&result).Error

	return result, err
}

// UpdateDataByName implements RepaymentRepository
func (db *dbRepayment) UpdateDataByName(name string) error {
	return db.Conn.Model(&models.Repayment{}).Where("name", name).Update("fee_stamp_duty", 10000).Error
}

// GetAmountByName implements RepaymentRepository
func (db *dbRepayment) GetAmountByName(name string) (float64, error) {
	var data float64

	result := db.Conn.Model(&models.Repayment{}).Select("sum(amount) as amount").Where("name", name).Group("name").Scan(&data).Error

	return data, result
}

// CreateRepayment implements RepaymentRepository
func (db *dbRepayment) CreateRepayment(repayment models.Repayment) error {
	err := db.Conn.Create(&repayment).Error

	if err != nil {
		log.Printf("Error creating Barang: %v", err)
	}

	return err
}

// GetAll implements RepaymentRepository
func (db *dbRepayment) GetAll() ([]models.Repayment, error) {
	var data []models.Repayment

	result := db.Conn.Find(&data).Error

	return data, result
}

type RepaymentRepository interface {
	CreateRepayment(repayment models.Repayment) error
	GetAmountByName(name string) (float64, error)
	UpdateDataByName(name string) error
	GetDataByName(name string) ([]models.Repayment, error)
	GetAll() ([]models.Repayment, error)
	GetPembayaran(name string) (map[string]interface{}, error)
}

func NewRepaymentRepository(Conn *gorm.DB) RepaymentRepository {
	return &dbRepayment{Conn: Conn}
}
