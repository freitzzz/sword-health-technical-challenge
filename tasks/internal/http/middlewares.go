package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	dbMiddlewareKey = "db"
)

func RegisterMiddlewares(e *echo.Echo, db *gorm.DB) {

	e.Use(DbAccessMiddleware(db))
	e.Use(ResourceIdentifierValidationMiddleware())

}

func DbAccessMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(dbMiddlewareKey, db)
			next(c)
			return nil
		}
	}
}

func ResourceIdentifierValidationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			params := c.ParamNames()
			for _, param := range params {
				if strings.Contains(param, taskId) {
					id, err := strconv.Atoi(c.Param(param))
					if err != nil || id <= 0 {
						return c.NoContent(http.StatusNotFound)
					}
				}
			}
			next(c)
			return nil
		}
	}
}
