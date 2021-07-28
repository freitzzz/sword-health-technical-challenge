package http

import (
	"net/http"

	"github.com/freitzzz/sword-health-technical-challenge/notifications/internal/domain"
	"github.com/labstack/echo/v4"
)

// Tooled with json-go-struct mapper: https://mholt.github.io/json-to-go/

type NotificationPage []struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
}

func ToNotificationPage(notifications []*domain.Notification) NotificationPage {

	np := make(NotificationPage, len(notifications))

	for i, n := range notifications {
		np[i] = struct {
			ID      uint   "json:\"id\""
			Message string "json:\"message\""
		}{
			ID:      n.ID,
			Message: n.Message,
		}
	}

	return np

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

func Ok(c echo.Context, body interface{}) error {
	return c.JSON(http.StatusOK, body)
}
