package main

import (
	"go-logger/config"
	m "go-logger/middleware"
	"go-logger/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Setup Echo
	e := echo.New()

	// Setup logrus
	config.SetupLogrus()

	// Middleware
	e.Use(m.LoggerMiddleware)
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = m.ErrorHandlerMiddleware

	// Initialize routes
	routes.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}
