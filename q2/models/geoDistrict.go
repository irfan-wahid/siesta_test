package models

import "time"

type GeoDistrict struct {
	ID         int       `gorm:"column:id;primaryKey" json:"id"`
	CityID     int       `gorm:"column:city_id" json:"city_id"`
	Name       string    `gorm:"column:name" json:"name"`
	Lat        float64   `gorm:"column:lat" json:"lat"`
	Long       float64   `gorm:"column:long" json:"long"`
	CreatedOn  time.Time `gorm:"column:created_on" json:"created_on"`
	ModifiedOn time.Time `gorm:"column:modified_on" json:"modified_on"`
	Deleted    int       `gorm:"column:deleted" json:"deleted"`
}

func (GeoDistrict) TableName() string {
	return "geo_districts"
}
