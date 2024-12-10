package routes

import (
	_ "go-logger/docs"
	"go-logger/handler"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", handler.HelloWorld)

	e.GET("/error", handler.GetErrorResponse)

	e.GET("/panic", handler.GetPanicResponse)

	e.POST("/login", handler.Login)
}
