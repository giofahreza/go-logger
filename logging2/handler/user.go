package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Login
// @Description Login
// @Accept  json
// @Produce  json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} map[string]string
// @Router /login [post]
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username or Password is empty")
	}

	if username != "admin" || password != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login success",
	})
}
