package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/walkmanrd/assessment/repositories"
	"github.com/walkmanrd/assessment/types"
)

type ExpenseController struct{}

var expenseRepository repositories.ExpenseRepository
var requestExpense types.ExpenseRequest

func (c *ExpenseController) Store(e echo.Context) error {
	err := e.Bind(&requestExpense)

	if err != nil {
		return e.JSON(http.StatusBadRequest, types.Error{Message: err.Error()})
	}

	if err = e.Validate(requestExpense); err != nil {
		return err
	}

	expense, err := expenseRepository.Create(requestExpense)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, types.Error{Message: err.Error()})
	}

	return e.JSON(http.StatusCreated, expense)
}
