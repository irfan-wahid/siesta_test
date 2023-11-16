package repositories

import (
	"q2/models"

	"gorm.io/gorm"
)

type dbDistrict struct {
	Conn *gorm.DB
}

// GetDistric implements ChargeRepository
func (db *dbDistrict) GetDistrict() (models.GeoDistrict, error) {
	var district models.GeoDistrict

	result := db.Conn.Select("* ,6371 * 2 * ASIN(SQRT(POWER(SIN((3.234587754461181 - ABS(geo_districts.lat)) * PI() / 180 / 2), 2) +COS(3.234587754461181 * PI() / 180) * COS(ABS(geo_districts.lat) * PI() / 180) *POWER(SIN((97.32793709754745 - geo_districts.long) * PI() / 180 / 2), 2))) AS distance").
		Order("distance").
		Limit(1).
		Find(&district).Error

	return district, result
}

type DistrictRepository interface {
	GetDistrict() (models.GeoDistrict, error)
}

func NewDistrictRepository(Conn *gorm.DB) DistrictRepository {
	return &dbDistrict{Conn: Conn}
}
