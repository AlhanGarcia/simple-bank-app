package models

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	AccountID uint    `json:"account_id"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"` // "credit" or "debit"
}
