package http

import (
	"crypto/cipher"
	"net/http"

	"github.com/freitzzz/sword-health-technical-challenge/tasks/internal/domain"
	"github.com/labstack/echo/v4"
)

// Tooled with json-go-struct mapper: https://mholt.github.io/json-to-go/

type TaskPage []struct {
	ID     uint   `json:"id"`
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

func ToTaskView(task domain.Task, block cipher.Block) TaskView {

	return TaskView{
		ID:                 task.ID,
		UserID:             task.UserID,
		Summary:            domain.Summary(task, block),
		CreatedTimestampMS: task.CreatedAt.Unix(),
		UpdatedTimestampMS: task.UpdatedAt.Unix(),
	}
}

func ToTaskPage(tasks []*domain.Task) TaskPage {

	tp := make(TaskPage, len(tasks))

	for i, t := range tasks {
		tp[i] = struct {
			ID     uint   "json:\"id\""
			UserID string "json:\"userId\""
		}{
			ID:     t.ID,
			UserID: t.UserID,
		}
	}

	return tp

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

func Created(c echo.Context, body interface{}) error {
	return c.JSON(http.StatusCreated, body)
}
