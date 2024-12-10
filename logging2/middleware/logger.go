package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func MakeLogEntry(c echo.Context) *logrus.Entry {
	if c == nil {
		return logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	req := c.Request()
	return logrus.WithFields(logrus.Fields{
		"source": req.RemoteAddr,
		"method": req.Method,
		"uri":    req.RequestURI,
		"ip":     c.RealIP(),
	})
}

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		MakeLogEntry(c).Debug("request received")

		return next(c)
	}
}
