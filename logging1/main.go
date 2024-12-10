package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "go-logger/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Echo API
// @version 1
// @description This is a sample server Echo server.
// @host localhost:1234
// @BasePath /
// @schemes http
// @produces json
// @consumes json

// @Summary Get hello world
// @Description Get hello world
// @ID get-hello-world
// @Produce  json
// @Success 200 {string} string "Hello, World!"
// @Router / [get]
func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// @Summary Get error response
// @Description Get error response
// @ID get-error-response
// @Produce  json
// @Success 500 {string} string "Internal Server Error"
// @Router /error [get]
func getErrorResponse(c echo.Context) error {
	return errors.New("Error message")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		msg := "Internal Server Error"
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			msg = he.Message.(string)
			log.Println(he.Internal)
		}

		c.JSON(code, map[string]string{
			"error": msg,
		})
	}

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", helloWorld)

	e.GET("/error", getErrorResponse)

	e.GET("/panic", func(c echo.Context) error {
		panic("Panic!")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}
