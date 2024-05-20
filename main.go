package main

import (
	database "simple-bank-app/db"
	"simple-bank-app/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.InitDatabase()

	// Define routes
	r.POST("/accounts", handlers.CreateAccount)
	r.GET("/accounts", handlers.GetAccounts)
	r.POST("/transactions", handlers.CreateTransaction)

	r.Run()
}
