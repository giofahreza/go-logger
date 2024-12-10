package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandlerMiddleware(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	he, ok := err.(*echo.HTTPError)
	if ok {
		code = he.Code
		message = fmt.Sprintf("%v", he.Message)
	}

	msg := message
	if he.Internal != nil {
		msg = "Internal : " + he.Internal.Error() + " | UserMessage : " + message
	}
	MakeLogEntry(c).Error(msg)

	c.JSON(code, map[string]interface{}{
		"error": message,
	})
}
