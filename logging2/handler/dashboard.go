package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Routes
// @Summary Hello World
// @Description get string
// @Produce  plain
// @Success 200 {string} string "Hello, World!"
// @Router / [get]
func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Routes
// @Summary Get error response
// @Description Get error response
// @Produce  json
// @Success 500 {string} string "Internal Server Error"
// @Router /error [get]
func GetErrorResponse(c echo.Context) error {
	return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
}

// Routes
// @Summary Get panic response
// @Description Get panic response
// @Produce  json
// @Success 500 {string} string "Internal Server Error"
// @Router /panic [get]
func GetPanicResponse(c echo.Context) error {
	panic("panic")
}
