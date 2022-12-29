package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/walkmanrd/assessment/configs"
	"github.com/walkmanrd/assessment/controllers"
	"github.com/walkmanrd/assessment/types"

	_ "github.com/lib/pq"
)

// CustomValidator is a struct for custom validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate is a function to validate request
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// ExpenseController is a struct for expense controller
var expenseController controllers.ExpenseController

// init is a function that run before main
func init() {
	db := configs.ConnectDatabase()
	configs.AutoMigrate(db)
	defer db.Close()
}

// AuthHeader is a middleware to check authorization header
func AuthHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")

		if authorization == "November 10, 2009" {
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, types.Error{Message: "Unauthorized"})
	}
}

// main is a function that run after init
func main() {
	// Echo instance
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(AuthHeader)

	// Routes
	e.GET("/expenses/:id", expenseController.Show)
	e.POST("/expenses", expenseController.Store)
	e.PUT("/expenses/:id", expenseController.Update)

	// Start server
	port := os.Getenv("PORT")
	log.Println("Server started at " + os.Getenv("PORT"))
	log.Fatal(e.Start(port))
	log.Println("Bye bye!")
}
