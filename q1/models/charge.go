package models

import "time"

type Charge struct {
	Charge_ID      int       `json:"charge_id" gorm:"column:charge_id;primaryKey;autoIncrement"`
	Selling_date   time.Time `json:"selling_date" gorm:"column:selling_date;type:date"`
	Amount         float64   `json:"amount" gorm:"column:amount"`
	Is_deleted     int       `json:"is_deleted" gorm:"column:is_deleted"`
	Charge_type_ID int       `json:"charge_type_id" gorm:"foreignKey:charge_type_id"`
}

type Charge_Type struct {
	Charge_type_ID int    `json:"charge_type_id" gorm:"column:charge_type_id;primaryKey;autoIncrement"`
	Charge_type    string `json:"charge_type" gorm:"column:charge_type"`
	Is_deleted     int    `json:"is_deleted" gorm:"column:is_deleted"`
}

type Charge_type_task_list struct {
	Task_type_ID   int `json:"task_type_id" gorm:"column:task_type_id"`
	Charge_type_ID int `json:"charge_type_id" gorm:"column:charge_type_id"`
}

type Tax_type struct {
	Tax_type_ID int     `json:"tax_type_id" gorm:"column:task_type_id;primaryKey;autoIncrement"`
	Tax_type    string  `json:"tax_type" gorm:"column:task_type"`
	Percentage  float64 `json:"percentage" gorm:"column:percentage"`
	Is_deleted  int     `json:"is_deleted" gorm:"column:is_deleted"`
}

func (Charge) TableName() string {
	return "charge"
}

func (Charge_Type) TableName() string {
	return "charge_type"
}

func (Charge_type_task_list) TableName() string {
	return "charge_type_task_list"
}

func (Tax_type) TableName() string {
	return "tax_type"
}
