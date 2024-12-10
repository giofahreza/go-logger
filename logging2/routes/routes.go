package routes

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	// Initialize routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/error", func(c echo.Context) error {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request").WithInternal(errors.New("[Internal] Bad Request"))
	})

	e.GET("/panic", func(c echo.Context) error {
		panic("panic")
	})
}
