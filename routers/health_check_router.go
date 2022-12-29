package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/walkmanrd/assessment/controllers"
)

// healthCheckController is a struct for health check controller
var healthCheckController controllers.HealthCheckController

// HealthCheckRouter is a function to set health check router
func HealthCheckRouter(e *echo.Echo) {
	e.GET("/health-check", healthCheckController.Index)
}
