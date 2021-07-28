package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type InvalidParam struct {
	Message string
	Type    string
	Name    string
}

var (
	paginationIndexNotInteger = InvalidParam{
		Message: "Not an integer",
		Type:    "Query",
		Name:    paginationIndex,
	}

	summaryExceeds2500Characters = InvalidParam{
		Message: "Sumamry exceeds 2500 characters",
		Type:    "Body",
		Name:    summaryField,
	}
)

func InvalidParamBadRequest(c echo.Context, ip InvalidParam) error {
	return c.JSON(http.StatusBadRequest, ip)
}

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
