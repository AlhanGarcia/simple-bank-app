package models

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"`
}
