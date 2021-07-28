package http

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/logging"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	dbMiddlewareKey = "db"
	ucMiddlewareKey = "uc"
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

func TranslateHeadersInUserContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			headers := c.Request().Header

			uid := headers.Get(userIdHeader)

			role, perr := strconv.Atoi(headers.Get(userIdHeader))

			if perr != nil || (role < 0 || role > 1) {
				logging.LogWarning(fmt.Sprintf("Received request with invalid role =:> %d", role))
			}

			uc := UserContext{ID: uid, Role: role}

			c.Set(ucMiddlewareKey, uc)
			next(c)

			return nil
		}
	}
}
