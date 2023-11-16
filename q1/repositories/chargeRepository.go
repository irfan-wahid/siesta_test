package repositories

import (
	"q1/models"

	"gorm.io/gorm"
)

type dbCharge struct {
	Conn *gorm.DB
}

// GetChargeTax implements ChargeRepository
func (db *dbCharge) GetChargeTax(chargeID int) (models.Tax_type, error) {
	var taxType models.Tax_type

	result := db.Conn.Table("charge").
		Select("tax_type.percentage").
		Joins("inner join charge_type on charge.charge_type_id = charge_type.charge_type_id and charge_type.is_deleted = false").
		Joins("inner join charge_type_tax_list charge_type_tax_list on charge_type.charge_type_id = charge_type_tax_list.charge_type_id").
		Joins("left join tax_type tax_type on charge_type_tax_list.tax_type_id = tax_type.tax_type_id and tax_type.is_deleted = false").
		Where("charge.is_deleted = ?", false).
		Where("charge.Charge_id = ?", chargeID).
		Scan(&taxType).Error

	return taxType, result
}

// GetChargeType implements ChargeRepository

func (db *dbCharge) GetChargeType(chargeID int) (models.Charge_Type, error) {
	var chargeType models.Charge_Type

	result := db.Conn.Table("charge_type").
		Select("charge_type.Charge_type").
		Joins("inner join charge on charge.Charge_type_ID = charge_type.Charge_type_ID").
		Where("charge.Charge_id = ?", chargeID).
		Scan(&chargeType).Error

	return chargeType, result
}

// GetAll implements ChargeRepository
func (db *dbCharge) GetCharge() ([]models.Charge, error) {
	var data []models.Charge

	result := db.Conn.Joins("inner join charge_type on charge.charge_type_id = charge_type.charge_type_id and charge_type.is_deleted = false").
		Joins("inner join charge_type_tax_list charge_type_tax_list on charge_type.charge_type_id = charge_type_tax_list.charge_type_id").
		Joins("left join tax_type tax_type on charge_type_tax_list.tax_type_id = tax_type.tax_type_id and tax_type.is_deleted = false").
		Where("charge.is_deleted = ?", false).
		Find(&data).Error

	return data, result
}

type ChargeRepository interface {
	GetCharge() ([]models.Charge, error)
	GetChargeType(chargeID int) (models.Charge_Type, error)
	GetChargeTax(chargeID int) (models.Tax_type, error)
}

func NewChargeRepository(Conn *gorm.DB) ChargeRepository {
	return &dbCharge{Conn: Conn}
}
