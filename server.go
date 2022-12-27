package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/walkmanrd/assessment/configs"
	"github.com/walkmanrd/assessment/controllers"

	_ "github.com/lib/pq"
)

var db *sql.DB

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

var expenseController controllers.ExpenseController

func init() {
	db = configs.ConnectDatabase()
	configs.AutoMigrate(db)
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/expenses", expenseController.Store)

	defer db.Close()
	port := os.Getenv("PORT")

	log.Println("Server started at " + port)
	log.Fatal(e.Start(port))
	log.Println("Bye bye!")
}
