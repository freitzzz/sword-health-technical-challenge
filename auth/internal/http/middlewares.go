package http

import (
	"fmt"

	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/domain"
	"github.com/freitzzz/sword-health-technical-challenge/auth/internal/logging"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	dbMiddlewareKey = "db"
	jbMiddlewareKey = "jb"
)

func RegisterMiddlewares(e *echo.Echo, db *gorm.DB, jb domain.JWTBundle) {

	e.Use(dbAccessMiddleware(db))
	e.Use(jwtBundleAccessMiddleware(jb))
	e.Use(loggingMiddleware())

}

func dbAccessMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(dbMiddlewareKey, db)
			next(c)
			return nil
		}
	}
}

func jwtBundleAccessMiddleware(jb domain.JWTBundle) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(jbMiddlewareKey, jb)
			next(c)
			return nil
		}
	}
}

func loggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			req := c.Request()

			logging.LogInfo(fmt.Sprintf("Host: %s | Method: %s | Path: %s", req.Host, req.Method, req.URL.RequestURI()))

			next(c)

			return nil
		}
	}
}
