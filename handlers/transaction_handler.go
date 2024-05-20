package handlers

import (
	"net/http"
	database "simple-bank-app/db"
	"simple-bank-app/models"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account models.Account
	// Find account by ID
	if err := database.DB.First(&account, transaction.AccountID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	//Update the account balance based  on the transaction type
	if transaction.Type == "debit" && account.Balance < transaction.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient funds"})
		return
	}

	if transaction.Type == "debit" {
		account.Balance -= transaction.Amount
	} else {
		account.Balance += transaction.Amount
	}

	database.DB.Save(&account)
	database.DB.Create(&transaction)
	c.JSON(http.StatusOK, transaction)

}
