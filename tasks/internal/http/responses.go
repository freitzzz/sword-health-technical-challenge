package http

import (
	"net/http"

	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/domain"
	"github.com/labstack/echo/v4"
)

// Tooled with json-go-struct mapper: https://mholt.github.io/json-to-go/

type TaskPage []struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
}

type TaskView struct {
	ID                 uint   `json:"id"`
	UserID             string `json:"userId"`
	Summary            string `json:"summary"`
	CreatedTimestampMS int64  `json:"createdTimestampMS"`
	UpdatedTimestampMS int64  `json:"updatedTimestampMS"`
}

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

func ToTaskView(task domain.Task) TaskView {
	return TaskView{
		ID:                 task.ID,
		UserID:             task.UserID,
		Summary:            task.Summary,
		CreatedTimestampMS: task.CreatedAt.Unix(),
		UpdatedTimestampMS: task.UpdatedAt.Unix(),
	}
}

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

func Ok(c echo.Context, body interface{}) error {
	return c.JSON(http.StatusOK, body)
}
