package main

import (
	"expense-tracker/config"
	"expense-tracker/models"
	"expense-tracker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	request := gin.Default()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Expense{})

	routes.RegisterExpenseRoute(request)
	request.Run(":8080")
}
