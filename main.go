package main

import (
	"expense-tracker/config"

	"github.com/gin-gonic/gin"
)

func main() {
	request := gin.Default()
	config.ConnectDB()
	request.Run(":8080")
}
