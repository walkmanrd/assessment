package controllers

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/walkmanrd/assessment/repositories"
	"github.com/walkmanrd/assessment/types"
)

type ExpenseController struct {
	requestExpense    types.ExpenseRequest
	expenseRepository repositories.ExpenseRepository
}

// Show is a function to get an expense by id
func (c *ExpenseController) Show(e echo.Context) error {
	id := e.Param("id")
	expense, err := c.expenseRepository.FindOne(id)

	switch err {
	case sql.ErrNoRows:
		return e.JSON(http.StatusNotFound, types.Error{Message: "expense not found"})
	case nil:
		return e.JSON(http.StatusOK, expense)
	default:
		return e.JSON(http.StatusInternalServerError, types.Error{Message: err.Error()})
	}
}

// Store is a function to create a new expense
func (c *ExpenseController) Store(e echo.Context) error {
	err := e.Bind(&c.requestExpense)

	if err != nil {
		return e.JSON(http.StatusBadRequest, types.Error{Message: err.Error()})
	}

	if err = e.Validate(c.requestExpense); err != nil {
		return err
	}

	expense, err := c.expenseRepository.Create(c.requestExpense)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, types.Error{Message: err.Error()})
	}

	return e.JSON(http.StatusCreated, expense)
}
