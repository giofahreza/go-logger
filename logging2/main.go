package main

import (
	"go-logger/config"
	m "go-logger/middleware"
	"go-logger/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @host localhost:1234
// @BasePath /

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
