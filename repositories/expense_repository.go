package repositories

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/walkmanrd/assessment/configs"
	"github.com/walkmanrd/assessment/models"
	"github.com/walkmanrd/assessment/types"
)

// ExpenseRepository is a repository for expense
type ExpenseRepository struct {
	db *sql.DB
}

// FindOne is a function to get an expenses by id
func (r *ExpenseRepository) FindOne(id string) (models.Expense, error) {
	r.db = configs.ConnectDatabase()
	defer r.db.Close()

	stmt, err := r.db.Prepare("SELECT * FROM expenses WHERE id = $1")
	if err != nil {
		return models.Expense{}, err
	}

	row := stmt.QueryRow(id)
	expense := models.Expense{}

	err = row.Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
	if err != nil {
		return models.Expense{}, err
	}

	return expense, nil
}

// Create is a function to create a new expense
func (r *ExpenseRepository) Create(expenseRequest types.ExpenseRequest) (models.Expense, error) {
	r.db = configs.ConnectDatabase()
	defer r.db.Close()

	sqlCommand := `INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4) RETURNING id`
	expense := models.Expense{
		Title:  expenseRequest.Title,
		Amount: expenseRequest.Amount,
		Note:   expenseRequest.Note,
		Tags:   expenseRequest.Tags,
	}

	tags := pq.Array(expense.Tags)
	row := r.db.QueryRow(sqlCommand, &expense.Title, &expense.Amount, &expense.Note, tags)

	err := row.Scan(&expense.ID)
	if err != nil {
		fmt.Println("can't scan id on ExpenseRepository", err)
		return models.Expense{}, err
	}

	return expense, nil
}
