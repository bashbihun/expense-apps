package routes

import (
	"expense-tracker/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterExpenseRoute(request *gin.Engine) {
	expense := request.Group("/expense")
	{
		expense.POST("/", controllers.CreateExpense)
		expense.GET("/", controllers.GetAllExpense)
		expense.GET("/:id", controllers.GetExpenseById)
		expense.PUT("/:id", controllers.UpdateExpense)
	}
}
