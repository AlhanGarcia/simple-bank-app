package handlers

import (
	"net/http"
	database "simple-bank-app/db"
	"simple-bank-app/models"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&account)
	c.JSON(http.StatusOK, account)
}

func GetAccounts(c *gin.Context) {
	var accounts []models.Account
	database.DB.Find(&accounts)
	c.JSON(http.StatusOK, accounts)
}
