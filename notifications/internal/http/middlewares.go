package http

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/logging"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	dbMiddlewareKey = "db"
	ucMiddlewareKey = "uc"
)

func RegisterMiddlewares(e *echo.Echo, db *gorm.DB) {

	e.Use(dbAccessMiddleware(db))
	e.Use(resourceIdentifierValidationMiddleware())
	e.Use(translateHeadersInUserContextMiddleware())
	e.Use(onlyAllowManagerMiddleware())
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

func resourceIdentifierValidationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			params := c.ParamNames()
			for _, param := range params {
				if strings.Contains(param, notificationId) {
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

func translateHeadersInUserContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			headers := c.Request().Header

			uid := headers.Get(userIdHeader)

			if len(uid) == 0 {

				logging.LogWarning("Received request with empty user id")

				return NotAuthorized(c)

			}

			role, perr := strconv.Atoi(headers.Get(roleHeader))

			if perr != nil || (role < 0 || role > 1) {
				logging.LogWarning(fmt.Sprintf("Received request with invalid role =:> %d", role))

				return NotAuthorized(c)
			}

			uc := UserContext{ID: uid, Role: role}

			c.Set(ucMiddlewareKey, uc)
			next(c)

			return nil
		}
	}
}

func onlyAllowManagerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			uc, uok := c.Get(ucMiddlewareKey).(UserContext)

			if !uok {
				logging.LogError("User Context not available in middleware")

				return InternalServerError(c)
			}

			if IsTechnician(uc) {
				return NotAuthorized(c)
			}

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
