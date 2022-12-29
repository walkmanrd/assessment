package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/walkmanrd/assessment/repositories"
	"github.com/walkmanrd/assessment/services"
	"github.com/walkmanrd/assessment/types"
)

type ExpenseController struct {
	requestExpense    types.ExpenseRequest
	expenseService    services.ExpenseService
	expenseRepository repositories.ExpenseRepository
}

// Show is a function to get an expense by id
func (c *ExpenseController) Show(e echo.Context) error {
	id := e.Param("id")

	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		return e.JSON(http.StatusBadRequest, types.Error{Message: "invalid parameter id"})
	}

	expense, status, err := c.expenseService.GetById(id)

	if err != nil {
		return e.JSON(status, types.Error{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, expense)
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

	expense, err := c.expenseService.Create(c.requestExpense)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, types.Error{Message: err.Error()})
	}

	return e.JSON(http.StatusCreated, expense)
}

// Update is a function to get an expense by id
func (c *ExpenseController) Update(e echo.Context) error {
	id := e.Param("id")

	if _, err := strconv.ParseInt(id, 10, 64); err != nil {
		return e.JSON(http.StatusBadRequest, types.Error{Message: "invalid parameter id"})
	}

	err := e.Bind(&c.requestExpense)

	if err != nil {
		return e.JSON(http.StatusBadRequest, types.Error{Message: err.Error()})
	}

	_, status, err := c.expenseService.GetById(id)

	if err != nil {
		return e.JSON(status, types.Error{Message: err.Error()})
	}

	newExpense, err := c.expenseService.UpdateById(id, c.requestExpense)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, types.Error{Message: err.Error()})
	}

	return e.JSON(http.StatusOK, newExpense)
}
