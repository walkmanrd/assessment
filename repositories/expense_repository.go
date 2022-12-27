package repositories

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/walkmanrd/assessment/configs"
	"github.com/walkmanrd/assessment/models"
)

type ExpenseRepository struct{}

var db *sql.DB

func (r *ExpenseRepository) Create(expense models.Expense) (models.Expense, error) {
	db = configs.ConnectDatabase()
	sqlCommand := `INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4) RETURNING id`
	tags := pq.Array(expense.Tags)
	row := db.QueryRow(sqlCommand, &expense.Title, &expense.Amount, &expense.Note, tags)

	err := row.Scan(&expense.ID)

	if err != nil {
		fmt.Println("can't scan id on ExpenseRepository", err)
		return models.Expense{}, err
	}

	return expense, nil
}
