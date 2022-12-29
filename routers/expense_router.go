package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/walkmanrd/assessment/controllers"
)

// ExpenseController is a struct for expense controller
var expenseController controllers.ExpenseController

// ExpenseRouter is a function to set expense routes
func ExpenseRouter(e *echo.Echo) {
	e.GET("/expenses", expenseController.Index)
	e.GET("/expenses/:id", expenseController.Show)
	e.POST("/expenses", expenseController.Store)
	e.PUT("/expenses/:id", expenseController.Update)
}
