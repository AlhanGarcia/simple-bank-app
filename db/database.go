package database

import (
	"simple-bank-app/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connec to database")
	}

	// migrate the schema to keep it up to date
	DB.AutoMigrate(&models.Account{}, &models.Transaction{})
}
