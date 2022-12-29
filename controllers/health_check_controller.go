package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ExpenseController is a struct for expense controller
type HealthCheckController struct{}

// GET /health-check
// Index is a function to get all expenses
func (c *HealthCheckController) Index(e echo.Context) error {
	return e.JSON(http.StatusOK, []byte(`{"status": "ok"}`))
}
