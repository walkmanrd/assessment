package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/walkmanrd/assessment/controllers"
)

// ExpenseController is a struct for expense controller
var expenseController controllers.ExpenseController

// ExpenseRouter is a function to set expense routes
func ExpenseRouter(e *echo.Group) {
	e.GET("", expenseController.Index)
	e.GET("/:id", expenseController.Show)
	e.POST("", expenseController.Store)
	e.PUT("/:id", expenseController.Update)
}
