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

// FindAll is a function to get all expenses
func (r *ExpenseRepository) FindAll() ([]models.Expense, error) {
	r.db = configs.ConnectDatabase()

	stmt, err := r.db.Prepare("SELECT * FROM expenses ORDER BY id ASC")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	expenses := []models.Expense{}

	for rows.Next() {
		expense := models.Expense{}
		err := rows.Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
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

	sqlCommand := `
	INSERT INTO expenses (id, title, amount, note, tags) values (DEFAULT, $1, $2, $3, $4)
	RETURNING id, title, amount, note, tags`

	expense := models.Expense{
		Title:  expenseRequest.Title,
		Amount: expenseRequest.Amount,
		Note:   expenseRequest.Note,
		Tags:   expenseRequest.Tags,
	}

	tags := pq.Array(expense.Tags)
	row := r.db.QueryRow(sqlCommand, &expense.Title, &expense.Amount, &expense.Note, tags)
	err := row.Scan(&expense.ID, &expense.Title, &expense.Amount, &expense.Note, pq.Array(&expense.Tags))

	defer r.db.Close()

	if err != nil {
		fmt.Println("can't scan id on ExpenseRepository", err)
		return models.Expense{}, err
	}

	return expense, nil
}

// Update is a function to update an expense by id
func (r *ExpenseRepository) Update(id string, expenseRequest types.ExpenseRequest) (models.Expense, error) {
	r.db = configs.ConnectDatabase()
	defer r.db.Close()

	sqlCommand := `
	UPDATE expenses SET title = $1, amount = $2, note = $3, tags = $4
	WHERE id = $5
	RETURNING id`

	expense := models.Expense{
		Title:  expenseRequest.Title,
		Amount: expenseRequest.Amount,
		Note:   expenseRequest.Note,
		Tags:   expenseRequest.Tags,
	}

	tags := pq.Array(expenseRequest.Tags)
	row := r.db.QueryRow(sqlCommand, &expense.Title, &expense.Amount, &expense.Note, tags, id)
	err := row.Scan(&expense.ID)

	if err != nil {
		fmt.Println("can't scan id on ExpenseRepository", err)
		return models.Expense{}, err
	}

	return expense, nil
}
