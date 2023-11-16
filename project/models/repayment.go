package models

import "time"

type Repayment struct {
	ID               int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name             string    `json:"name" gorm:"column:name"`
	Loan_date        time.Time `json:"loan_date" gorm:"column:loan_date;type:date"`
	Amount           float64   `json:"amount" gorm:"column:amount"`
	Tenor            float64   `json:"tenor" gorm:"column:tenor"`
	Principal        float64   `json:"principal" gorm:"column:principal"`
	Principal_month  float64   `json:"principal_month" gorm:"column:principal_month"`
	Fee              float64   `json:"fee" gorm:"column:fee"`
	Interest         float64   `json:"interest" gorm:"column:interest"`
	Total_interest   float64   `json:"total_interest" gorm:"column:total_interest"`
	Fee_stamp_duty   float64   `json:"fee_stamp_duty" gorm:"column:fee_stamp_duty"`
	Total_Payment    float64   `json:"total_payment" gorm:"column:total_payment"`
	Demand_per_month float64   `json:"demand_per_month" gorm:"column:demand_per_month"`
}

func (Repayment) TableName() string {
	return "repayment"
}
