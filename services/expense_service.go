package services

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/walkmanrd/assessment/models"
	"github.com/walkmanrd/assessment/repositories"
	"github.com/walkmanrd/assessment/types"
)

type ExpenseServicer interface {
	// Gets is a function to get all expenses
	Gets() ([]models.Expense, error)

	// GetById is a function to get an expense by id
	GetById(id string) (models.Expense, int, error)

	// Create is a function to create a new expense
	Create(expenseRequest types.ExpenseRequest) (models.Expense, error)

	// UpdateById is a function to update an expense by id
	UpdateById(id string, expenseRequest types.ExpenseRequest) (models.Expense, error)
}

// ExpenseService is a struct for expense service
type ExpenseService struct {
	expenseRepository repositories.ExpenseRepository
}

// GetById is a service function to get an expense by id
func (c *ExpenseService) Gets() ([]models.Expense, error) {
	expenses, err := c.expenseRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return expenses, nil
}

// GetById is a service function to get an expense by id
func (c *ExpenseService) GetById(id string) (models.Expense, int, error) {
	expense, err := c.expenseRepository.FindOne(id)

	switch err {
	case sql.ErrNoRows:
		return models.Expense{}, http.StatusNotFound, errors.New("expense not found")
	case nil:
		return expense, 0, nil
	default:
		return models.Expense{}, http.StatusInternalServerError, err
	}
}

// Create is a service function to create a new expense
func (c *ExpenseService) Create(expenseRequest types.ExpenseRequest) (models.Expense, error) {
	expense, err := c.expenseRepository.Create(expenseRequest)

	if err != nil {
		return models.Expense{}, err
	}

	return expense, nil
}

// UpdateById is a service function to update an expense by id
func (c *ExpenseService) UpdateById(id string, expenseRequest types.ExpenseRequest) (models.Expense, error) {
	expense, err := c.expenseRepository.Update(id, expenseRequest)

	if err != nil {
		return models.Expense{}, err
	}

	return expense, nil
}
