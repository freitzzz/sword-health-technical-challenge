package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InternalServerError(c echo.Context) error {
	return c.NoContent(http.StatusInternalServerError)
}

func NotAuthorized(c echo.Context) error {
	return c.NoContent(http.StatusUnauthorized)
}

func NotFound(c echo.Context) error {
	return c.NoContent(http.StatusNotFound)
}

func NoContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func Ok(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
